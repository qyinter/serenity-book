package entity

import (
	"time"
)

const TableNameRating = "ratings"

// Rating mapped from table <ratings>
type Rating struct {
	RatingID  int64     `gorm:"column:rating_id;primaryKey;" json:"ratingID"`
	UserID    int64     `gorm:"column:user_id" json:"userID"`
	NovelID   int64     `gorm:"column:novel_id" json:"novelID"`
	Score     int32     `gorm:"column:score" json:"score"`
	RatedDate time.Time `gorm:"column:rated_date" json:"ratedDate"`
}

// TableName Rating's table name
func (*Rating) TableName() string {
	return TableNameRating
}
