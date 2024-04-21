package model

import "gorm.io/gorm"

type System struct {
	gorm.Model
}

func (m *System) TableName() string {
	return "system"
}
