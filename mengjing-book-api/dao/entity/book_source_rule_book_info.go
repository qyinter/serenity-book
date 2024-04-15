package entity

const TableNameBookSourceRuleBookInfo = "book_source_rule_book_info"

// BookSourceRuleBookInfo rule_book_info
type BookSourceRuleBookInfo struct {
	RuleBookInfoID int64  `gorm:"column:rule_book_info_id;primaryKey;comment:rule_book_info_id" json:"ruleBookInfoID"` // rule_book_info_id
	BookSourceID   int64  `gorm:"column:book_source_id;comment:root_id" json:"bookSourceID"`                           // root_id
	Author         string `gorm:"column:author;comment:作者" json:"author"`                                              // 作者
	CoverURL       string `gorm:"column:cover_url;comment:封面网址" json:"coverURL"`                                       // 封面网址
	Intro          string `gorm:"column:intro;comment:介绍" json:"intro"`                                                // 介绍
	Kind           string `gorm:"column:kind;comment:种类" json:"kind"`                                                  // 种类
	LastChapter    string `gorm:"column:last_chapter;comment:最后一章" json:"lastChapter"`                                 // 最后一章
	Name           string `gorm:"column:name;comment:姓名" json:"name"`                                                  // 姓名
}

// TableName BookSourceRuleBookInfo's table name
func (*BookSourceRuleBookInfo) TableName() string {
	return TableNameBookSourceRuleBookInfo
}
