package entity

const TableNameBookSourceRuleContent = "book_source_rule_content"

// BookSourceRuleContent rule_content
type BookSourceRuleContent struct {
	RuleContentID  int64  `gorm:"column:rule_content_id;primaryKey;comment:rule_content_id" json:"ruleContentID"` // rule_content_id
	BookSourceID   int64  `gorm:"column:book_source_id;comment:info Id" json:"bookSourceID"`                      // info Id
	Content        string `gorm:"column:content;comment:内容" json:"content"`                                       // 内容
	NextContentURL string `gorm:"column:next_content_url;comment:下一个内容URL" json:"nextContentURL"`                 // 下一个内容URL
}

// TableName BookSourceRuleContent's table name
func (*BookSourceRuleContent) TableName() string {
	return TableNameBookSourceRuleContent
}
