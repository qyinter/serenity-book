package entity

import (
	"time"
)

const TableNameFavorite = "favorites"

// Favorite mapped from table <favorites>
type Favorite struct {
	FavoriteID int64     `gorm:"column:favorite_id;primaryKey;" json:"favoriteID"`
	UserID     int64     `gorm:"column:user_id" json:"userID"`
	NovelID    int64     `gorm:"column:novel_id" json:"novelID"`
	AddedDate  time.Time `gorm:"column:added_date" json:"addedDate"`
}

// TableName Favorite's table name
func (*Favorite) TableName() string {
	return TableNameFavorite
}
