package entity

import (
	"encoding/json"
	"github.com/ulule/deepcopier"
	"strconv"
	"time"
)

const TableNameBookSourceInfo = "book_source_info"

// BookSourceInfo root
type BookSourceInfo struct {
	BookSourceID      int64  `gorm:"column:book_source_id;primaryKey;comment:爬虫任务id" json:"bookSourceID"`            // 爬虫任务id
	BookSourceComment string `gorm:"column:book_source_comment;comment:书源评论" json:"bookSourceComment"`               // 书源评论
	BookSourceGroup   string `gorm:"column:book_source_group;comment:书源组" json:"bookSourceGroup"`                    // 书源组
	BookSourceName    string `gorm:"column:book_source_name;comment:书源名称" json:"bookSourceName"`                     // 书源名称
	BookSourceType    int64  `gorm:"column:book_source_type;comment:书籍来源类型" json:"bookSourceType"`                   // 书籍来源类型
	BookSourceURL     string `gorm:"column:book_source_url;comment:书源网址" json:"bookSourceURL"`                       // 书源网址
	BookURLPattern    string `gorm:"column:book_url_pattern;comment:图书网址模式" json:"bookURLPattern"`                   // 图书网址模式
	CustomOrder       int64  `gorm:"column:custom_order;comment:custom_order" json:"customOrder"`                    // custom_order
	Enabled           string `gorm:"column:enabled;default:1;comment:已启用 0,不启用 1,启用" json:"enabled"`                 // 已启用 0,不启用 1,启用
	EnabledCookieJar  string `gorm:"column:enabled_cookie_jar;comment:启用cookie_jar" json:"enabledCookieJar"`         // 启用cookie_jar
	EnabledExplore    string `gorm:"column:enabled_explore;default:1;comment:启用探索 0,不启用 1,启用" json:"enabledExplore"` // 启用探索 0,不启用 1,启用
	Header            string `gorm:"column:header;comment:ua" json:"header"`                                         // ua
	LastUpdateTime    int64  `gorm:"column:last_update_time;comment:最后更新时间" json:"lastUpdateTime"`                   // 最后更新时间
	RespondTime       int64  `gorm:"column:respond_time;comment:响应时间" json:"respondTime"`                            // 响应时间
	SearchURL         string `gorm:"column:search_url;comment:搜索网址" json:"searchURL"`                                // 搜索网址
	Weight            int64  `gorm:"column:weight;comment:权重" json:"weight"`                                         // 权重
}

// TableName BookSourceInfo's table name
func (*BookSourceInfo) TableName() string {
	return TableNameBookSourceInfo
}

// MarshalJSON 实现序列化接口
func (info BookSourceInfo) MarshalJSON() ([]byte, error) {
	roleReturn := new(struct {
		BookSourceID      string `json:"bookSourceID"`
		BookSourceComment string `json:"bookSourceComment"`
		BookSourceGroup   string `json:"bookSourceGroup"`
		BookSourceName    string `json:"bookSourceName"`
		BookSourceType    int64  `json:"bookSourceType"`
		BookSourceURL     string `json:"bookSourceURL"`
		BookURLPattern    string `json:"bookURLPattern"`
		CustomOrder       int64  `json:"customOrder"`
		Enabled           string `json:"enabled"`
		EnabledCookieJar  string `json:"enabledCookieJar"`
		EnabledExplore    string `json:"enabledExplore"`
		Header            string `json:"header"`
		LastUpdateTime    string `json:"lastUpdateTime"`
		RespondTime       int64  `json:"respondTime"`
		SearchURL         string `json:"searchURL"`
		Weight            int64  `json:"weight"`
	})
	err := deepcopier.Copy(info).To(roleReturn)
	if err != nil {
		return nil, err
	}
	roleReturn.BookSourceID = strconv.FormatInt(info.BookSourceID, 10)
	roleReturn.LastUpdateTime = time.Unix(info.LastUpdateTime, 0).Format(time.DateTime)
	return json.Marshal(roleReturn)
}
