package model

import "time"

// SysUser represents a user in the system.
type SysUser struct {
	ID         int64     `gorm:"primaryKey;comment:'ユーザID'"`                           // 用户ID
	Email      string    `gorm:"size:50;unique;comment:'邮箱'"`                          // 邮箱
	Username   string    `gorm:"size:20;not null;unique;comment:'用户名'"`                // 用户名
	Password   string    `gorm:"size:100;not null;comment:'密码'"`                       // 密码
	Avatar     string    `gorm:"size:255;not null;default:'default.jpg';comment:'头像'"` // 头像
	Name       string    `gorm:"size:255;comment:'昵称'"`                                // 昵称
	LoginDate  time.Time `gorm:"comment:'最后登录时间'"`                                     // 最后登录时间
	LoginIP    string    `gorm:"size:50;comment:'登录ip'"`                               // 登录IP
	Status     int       `gorm:"not null;default:0;comment:'状态（0正常 1異常）'"`             // 状态（0正常 1异常）
	CreateTime time.Time `gorm:"not null;comment:'作成時間'"`                              // 创建时间
	UpdateTime time.Time `gorm:"not null;comment:'更新時間'"`                              // 更新时间
	IsDelete   int       `gorm:"not null;default:0"`                                   // 是否删除标志
}

// TableName sets the name of the table in the database that this struct is linked to.
func (SysUser) TableName() string {
	return "sys_user"
}
