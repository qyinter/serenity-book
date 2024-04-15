package entity

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	UserID       int64     `gorm:"column:user_id;primaryKey;" json:"userID"`
	Username     string    `gorm:"column:username" json:"username"`
	PasswordHash string    `gorm:"column:password_hash" json:"passwordHash"`
	Email        string    `gorm:"column:email" json:"email"`
	SignupDate   time.Time `gorm:"column:signup_date" json:"signupDate"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
