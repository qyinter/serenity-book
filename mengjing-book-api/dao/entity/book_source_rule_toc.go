package entity

const TableNameBookSourceRuleToc = "book_source_rule_toc"

// BookSourceRuleToc rule_toc
type BookSourceRuleToc struct {
	RuleTocID    int64  `gorm:"column:rule_toc_id;primaryKey;comment:规则目录id" json:"ruleTocID"` // 规则目录id
	BookSourceID int64  `gorm:"column:book_source_id;comment:root_id" json:"bookSourceID"`     // root_id
	ChapterList  string `gorm:"column:chapter_list;comment:章节列表" json:"chapterList"`           // 章节列表
	ChapterName  string `gorm:"column:chapter_name;comment:章节名称" json:"chapterName"`           // 章节名称
	ChapterURL   string `gorm:"column:chapter_url;comment:章节地址" json:"chapterURL"`             // 章节地址
}

// TableName BookSourceRuleToc's table name
func (*BookSourceRuleToc) TableName() string {
	return TableNameBookSourceRuleToc
}
