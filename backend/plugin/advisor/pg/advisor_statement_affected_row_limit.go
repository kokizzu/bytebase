package pg

// Framework code is generated by the generator.

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*StatementAffectedRowLimitAdvisor)(nil)
	_ ast.Visitor     = (*statementAffectedRowLimitChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLStatementAffectedRowLimit, &StatementAffectedRowLimitAdvisor{})
}

// StatementAffectedRowLimitAdvisor is the advisor checking for UPDATE/DELETE affected row limit.
type StatementAffectedRowLimitAdvisor struct {
}

// Check checks for UPDATE/DELETE affected row limit.
func (*StatementAffectedRowLimitAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &statementAffectedRowLimitChecker{
		level:  level,
		title:  string(ctx.Rule.Type),
		maxRow: payload.Number,
		ctx:    ctx.Context,
		driver: ctx.Driver,
	}

	if payload.Number > 0 && checker.driver != nil {
		for _, stmt := range stmtList {
			checker.text = advisor.NormalizeStatement(stmt.Text())
			ast.Walk(checker, stmt)
		}
	}

	return checker.adviceList, nil
}

type statementAffectedRowLimitChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	text       string
	title      string
	maxRow     int
	driver     *sql.DB
	ctx        context.Context
}

// Visit implements ast.Visitor interface.
func (checker *statementAffectedRowLimitChecker) Visit(in ast.Node) ast.Visitor {
	code := advisor.Ok
	rows := int64(0)
	switch node := in.(type) {
	case *ast.UpdateStmt, *ast.DeleteStmt:
		res, err := advisor.Query(checker.ctx, checker.driver, storepb.Engine_POSTGRES, fmt.Sprintf("EXPLAIN %s", node.Text()))
		if err != nil {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.InsertTooManyRows.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("\"%s\" dry runs failed: %s", checker.text, err.Error()),
				StartPosition: &storepb.Position{
					Line: int32(node.LastLine()),
				},
			})
		} else {
			rowCount, err := getAffectedRows(res)
			if err != nil {
				checker.adviceList = append(checker.adviceList, &storepb.Advice{
					Status:  checker.level,
					Code:    advisor.Internal.Int32(),
					Title:   checker.title,
					Content: fmt.Sprintf("failed to get row count for \"%s\": %s", checker.text, err.Error()),
					StartPosition: &storepb.Position{
						Line: int32(node.LastLine()),
					},
				})
			} else if rowCount > int64(checker.maxRow) {
				code = advisor.StatementAffectedRowExceedsLimit
				rows = rowCount
			}
		}
	}

	if code != advisor.Ok {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:  checker.level,
			Code:    code.Int32(),
			Title:   checker.title,
			Content: fmt.Sprintf("The statement \"%s\" affected %d rows (estimated). The count exceeds %d.", checker.text, rows, checker.maxRow),
			StartPosition: &storepb.Position{
				Line: int32(in.LastLine()),
			},
		})
	}
	return checker
}
