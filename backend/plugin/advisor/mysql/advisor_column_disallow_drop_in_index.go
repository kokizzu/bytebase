package mysql

// Framework code is generated by the generator.

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/catalog"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnDisallowDropInIndexAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLColumnDisallowDropInIndex, &ColumnDisallowDropInIndexAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLColumnDisallowDropInIndex, &ColumnDisallowDropInIndexAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLColumnDisallowDropInIndex, &ColumnDisallowDropInIndexAdvisor{})
}

// ColumnDisallowDropInIndexAdvisor is the advisor checking for disallow DROP COLUMN in index.
type ColumnDisallowDropInIndexAdvisor struct {
}

// Check checks for disallow Drop COLUMN in index statement.
func (*ColumnDisallowDropInIndexAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parser result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}

	checker := &columnDisallowDropInIndexChecker{
		level:   level,
		title:   string(ctx.Rule.Type),
		tables:  make(tableState),
		catalog: ctx.Catalog,
	}

	for _, stmtNode := range stmtList {
		checker.baseLine = stmtNode.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmtNode.Tree)
	}

	return checker.adviceList, nil
}

type columnDisallowDropInIndexChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine   int
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	tables     tableState // the variable mean whether the column in index.
	catalog    *catalog.Finder
}

func (checker *columnDisallowDropInIndexChecker) EnterCreateTable(ctx *mysql.CreateTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableName() == nil {
		return
	}
	if ctx.TableElementList() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableName(ctx.TableName())
	for _, tableElement := range ctx.TableElementList().AllTableElement() {
		if tableElement == nil || tableElement.TableConstraintDef() == nil {
			continue
		}
		if tableElement.TableConstraintDef().GetType_() == nil {
			continue
		}
		switch tableElement.TableConstraintDef().GetType_().GetTokenType() {
		case mysql.MySQLParserINDEX_SYMBOL, mysql.MySQLParserKEY_SYMBOL:
			// do nothing.
		default:
			continue
		}
		if tableElement.TableConstraintDef().KeyListVariants() == nil {
			continue
		}

		columnList := mysqlparser.NormalizeKeyListVariants(tableElement.TableConstraintDef().KeyListVariants())
		for _, column := range columnList {
			if checker.tables[tableName] == nil {
				checker.tables[tableName] = make(columnSet)
			}
			checker.tables[tableName][column] = true
		}
	}
}

func (checker *columnDisallowDropInIndexChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.AlterTableActions() == nil {
		return
	}
	if ctx.AlterTableActions().AlterCommandList() == nil {
		return
	}
	if ctx.AlterTableActions().AlterCommandList().AlterList() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableRef(ctx.TableRef())
	for _, item := range ctx.AlterTableActions().AlterCommandList().AlterList().AllAlterListItem() {
		if item == nil || item.DROP_SYMBOL() == nil || item.ColumnInternalRef() == nil {
			continue
		}

		index := checker.catalog.Origin.Index(&catalog.TableIndexFind{
			// In MySQL, the SchemaName is "".
			SchemaName: "",
			TableName:  tableName,
		})

		if index != nil {
			if checker.tables[tableName] == nil {
				checker.tables[tableName] = make(columnSet)
			}
			for _, indexColumn := range *index {
				for _, column := range indexColumn.ExpressionList() {
					checker.tables[tableName][column] = true
				}
			}
		}

		columnName := mysqlparser.NormalizeMySQLColumnInternalRef(item.ColumnInternalRef())
		if !checker.canDrop(tableName, columnName) {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.DropIndexColumn.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("`%s`.`%s` cannot drop index column", tableName, columnName),
				StartPosition: &storepb.Position{
					Line: int32(checker.baseLine + item.GetStart().GetLine()),
				},
			})
		}
	}
}

func (checker *columnDisallowDropInIndexChecker) canDrop(table string, column string) bool {
	if _, ok := checker.tables[table][column]; ok {
		return false
	}
	return true
}
