package mysql

// Framework code is generated by the generator.

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bytebase/bytebase/plugin/advisor"
)

func TestColumnCommentConvention(t *testing.T) {
	tests := []advisor.TestCase{
		{
			Statement: `CREATE TABLE t(a int COMMENT 'some comments')`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		{
			Statement: `CREATE TABLE t(
				a int COMMENT 'some comments',
				b int,
				c int)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.NoColumnComment,
					Title:   "column.comment",
					Content: "Column `t`.`b` requires comments",
					Line:    3,
				},
				{
					Status:  advisor.Warn,
					Code:    advisor.NoColumnComment,
					Title:   "column.comment",
					Content: "Column `t`.`c` requires comments",
					Line:    4,
				},
			},
		},
		{
			Statement: `ALTER TABLE t ADD COLUMN b int`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.NoColumnComment,
					Title:   "column.comment",
					Content: "Column `t`.`b` requires comments",
					Line:    1,
				},
			},
		},
		{
			Statement: `ALTER TABLE t CHANGE COLUMN a b int`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.NoColumnComment,
					Title:   "column.comment",
					Content: "Column `t`.`b` requires comments",
					Line:    1,
				},
			},
		},
		{
			Statement: `ALTER TABLE t MODIFY COLUMN b int`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.NoColumnComment,
					Title:   "column.comment",
					Content: "Column `t`.`b` requires comments",
					Line:    1,
				},
			},
		},
		{
			Statement: `ALTER TABLE t MODIFY COLUMN b int COMMENT 'abcdefghiakljhakljdsfalugelkhnabsdguelkadf'`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.ColumnCommentTooLong,
					Title:   "column.comment",
					Content: "The length of column `t`.`b` comment should be within 20 characters",
					Line:    1,
				},
			},
		},
	}

	payload, err := json.Marshal(advisor.CommentConventionRulePayload{
		Required:  true,
		MaxLength: 20,
	})
	require.NoError(t, err)
	advisor.RunSQLReviewRuleTests(t, tests, &ColumnCommentConventionAdvisor{}, &advisor.SQLReviewRule{
		Type:    advisor.SchemaRuleColumnCommentConvention,
		Level:   advisor.SchemaRuleLevelWarning,
		Payload: string(payload),
	}, advisor.MockMySQLDatabase)
}
