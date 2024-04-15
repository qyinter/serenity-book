package entity

import (
	"time"
)

const TableNameComment = "comments"

// Comment mapped from table <comments>
type Comment struct {
	CommentID   int64     `gorm:"column:comment_id;primaryKey;" json:"commentID"`
	ChapterID   int64     `gorm:"column:chapter_id" json:"chapterID"`
	UserID      int64     `gorm:"column:user_id" json:"userID"`
	Content     string    `gorm:"column:content" json:"content"`
	PublishDate time.Time `gorm:"column:publish_date" json:"publishDate"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
