// Package mssql is the advisor for MSSQL database.
package mssql

import (
	"testing"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

func TestMSSQLRules(t *testing.T) {
	mssqlRules := []advisor.SQLReviewRuleType{
		advisor.SchemaRuleStatementNoSelectAll,
		advisor.SchemaRuleTableNaming,
		advisor.SchemaRuleTableNameNoKeyword,
		advisor.SchemaRuleIdentifierNoKeyword,
		advisor.SchemaRuleStatementRequireWhere,
		advisor.SchemaRuleColumnMaximumVarcharLength,
		advisor.SchemaRuleTableDropNamingConvention,
		advisor.SchemaRuleTableRequirePK,
		advisor.SchemaRuleColumnNotNull,
		advisor.SchemaRuleTableNoFK,
		advisor.SchemaRuleTableDisallowDDL,
		advisor.SchemaRuleTableDisallowDML,
		advisor.SchemaRuleSchemaBackwardCompatibility,
		advisor.SchemaRuleRequiredColumn,
		advisor.SchemaRuleColumnTypeDisallowList,
		advisor.SchemaRuleFunctionDisallowCreate,
		advisor.SchemaRuleProcedureDisallowCreate,
	}

	for _, rule := range mssqlRules {
		advisor.RunSQLReviewRuleTest(t, rule, storepb.Engine_MSSQL, nil, false /* record */)
	}
}
