package entity

import (
	"time"
)

const TableNameNovel = "novels"

// Novel mapped from table <novels>
type Novel struct {
	NovelID     int64     `gorm:"column:novel_id;primaryKey;" json:"novelID"`
	AuthorID    int64     `gorm:"column:author_id" json:"authorID"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Genre       string    `gorm:"column:genre" json:"genre"`
	Status      string    `gorm:"column:status" json:"status"`
	PublishDate time.Time `gorm:"column:publish_date" json:"publishDate"`
}

// TableName Novel's table name
func (*Novel) TableName() string {
	return TableNameNovel
}
