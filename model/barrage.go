package model

// 弹幕
type Barrage struct {
	UUID       string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Video_UUID string `json:"video_uuid" validate:"required" gorm:"type:varchar(191)"`
	User_UUID  string `json:"user_uuid" validate:"required" gorm:"type:varchar(191)"`
	Content    string `json:"content" validate:"required"`
	Second     int64  `json:"second" validate:"required" gorm:"type:bigint"`
	Timestamp  int64  `json:"timestamp" validate:"required" gorm:"type:bigint"`
}

// BarrageListResp 弹幕列表响应
type BarrageListResp struct {
	Code        int       `json:"code"`
	BarrageList []Barrage `json:"barrage_list"`
	Total       int64     `json:"total,omitempty"`
}

// BarrageResp 弹幕响应
type BarrageResp struct {
	Code    int     `json:"code"`
	Barrage Barrage `json:"barrage"`
}

// BarrageCountResp 弹幕数量响应
type BarrageCountResp struct {
	Code  int   `json:"code"`
	Count int64 `json:"count"`
}
