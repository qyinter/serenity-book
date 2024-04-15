package model

type BookSourceInfoModel struct {
	BookSourceName string `json:"bookSourceName"`
	BookSourceURL  string `json:"bookSourceURL"`
	PageSize       int    `json:"pageSize"`
	PageNum        int    `json:"pageNum"`
}
