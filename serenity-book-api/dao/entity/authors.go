package entity

const TableNameAuthor = "authors"

// Author mapped from table <authors>
type Author struct {
	AuthorID  int64  `gorm:"column:author_id;primaryKey;" json:"authorID"`
	Name      string `gorm:"column:name" json:"name"`
	Biography string `gorm:"column:biography" json:"biography"`
	Email     string `gorm:"column:email" json:"email"`
}

// TableName Author's table name
func (*Author) TableName() string {
	return TableNameAuthor
}
