package entity

const TableNameBookSourceRuleSearch = "book_source_rule_search"

// BookSourceRuleSearch rule_search
type BookSourceRuleSearch struct {
	RuleSearchID int64  `gorm:"column:rule_search_id;primaryKey;comment:rule_search_id" json:"ruleSearchID"` // rule_search_id
	BookSourceID int64  `gorm:"column:book_source_id" json:"bookSourceID"`
	Author       string `gorm:"column:author;comment:tag.span.0@text" json:"author"`         // tag.span.0@text
	BookList     string `gorm:"column:book_list;comment:class.category-div" json:"bookList"` // class.category-div
	BookURL      string `gorm:"column:book_url;comment:tag.a.0@href" json:"bookURL"`         // tag.a.0@href
	CheckKeyWord string `gorm:"column:check_key_word;comment:剑来" json:"checkKeyWord"`        // 剑来
	CoverURL     string `gorm:"column:cover_url;comment:tag.img@src" json:"coverURL"`        // tag.img@src
	Intro        string `gorm:"column:intro;comment:class.intro@textNodes" json:"intro"`     // class.intro@textNodes
	Name         string `gorm:"column:name;comment:tag.h3@text" json:"name"`                 // tag.h3@text
}

// TableName BookSourceRuleSearch's table name
func (*BookSourceRuleSearch) TableName() string {
	return TableNameBookSourceRuleSearch
}
