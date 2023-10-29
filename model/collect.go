package model

// Collect 收藏
type Collect struct {
	UUID       string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Video_UUID string `json:"video_uuid" validate:"required" gorm:"type:varchar(191)"`
	User_UUID  string `json:"user_uuid" validate:"required" gorm:"type:varchar(191)"`
	Video      Video  `json:"video,omitempty" validate:"-" gorm:"-"`
}

// CollectResp 收藏响应
type CollectResp struct {
	Code    int     `json:"code"`
	Collect Collect `json:"collect"`
}

// 收藏状态响应
type CollectStatusResp struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
}

// CollectListResp 收藏列表响应
type CollectListResp struct {
	Code        int       `json:"code"`
	CollectList []Collect `json:"collect_list"`
}
