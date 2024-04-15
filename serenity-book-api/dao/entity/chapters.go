package entity

import (
	"time"
)

const TableNameChapter = "chapters"

// Chapter mapped from table <chapters>
type Chapter struct {
	ChapterID     int64     `gorm:"column:chapter_id;primaryKey;" json:"chapterID"`
	NovelID       int64     `gorm:"column:novel_id" json:"novelID"`
	Title         string    `gorm:"column:title" json:"title"`
	Content       string    `gorm:"column:content" json:"content"`
	ChapterNumber int32     `gorm:"column:chapter_number" json:"chapterNumber"`
	PublishDate   time.Time `gorm:"column:publish_date" json:"publishDate"`
}

// TableName Chapter's table name
func (*Chapter) TableName() string {
	return TableNameChapter
}
