package pg

// Framework code is generated by the generator.

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*StatementPriorBackupCheckAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLStatementPriorBackupCheck, &StatementPriorBackupCheckAdvisor{})
}

// StatementPriorBackupCheckAdvisor is the advisor checking for disallow mix DDL and DML.
type StatementPriorBackupCheckAdvisor struct {
}

// Check checks for disallow mix DDL and DML.
func (*StatementPriorBackupCheckAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	if ctx.PreUpdateBackupDetail == nil || ctx.ChangeType != storepb.PlanCheckRunConfig_DML {
		return nil, nil
	}
	var adviceList []*storepb.Advice
	stmtList, ok := ctx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	title := string(ctx.Rule.Type)

	for _, stmt := range stmtList {
		var isDDL bool
		if _, ok := stmt.(ast.DDLNode); ok {
			isDDL = true
		}
		if isDDL {
			adviceList = append(adviceList, &storepb.Advice{
				Status:  level,
				Title:   title,
				Content: fmt.Sprintf("Data change can only run DML, \"%s\" is not DML", stmt.Text()),
				Code:    advisor.StatementPriorBackupCheck.Int32(),
				StartPosition: &storepb.Position{
					Line: int32(stmt.LastLine()),
				},
			})
		}
	}

	name := extractDatabaseName(ctx.PreUpdateBackupDetail.Database)
	if !ctx.Catalog.Origin.HasSchema(name) {
		adviceList = append(adviceList, &storepb.Advice{
			Status:  level,
			Title:   title,
			Content: fmt.Sprintf("Need schema %q to do prior backup but it does not exist", name),
			Code:    advisor.SchemaNotExists.Int32(),
			StartPosition: &storepb.Position{
				Line: 0,
			},
		})
	}

	return adviceList, nil
}

func extractDatabaseName(databaseUID string) string {
	segments := strings.Split(databaseUID, "/")
	return segments[len(segments)-1]
}
