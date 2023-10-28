package model

// Collect 收藏
type Collect struct {
	UUID       string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Video_UUID string `json:"video_uuid" validate:"required"`
	User_UUID  string `json:"user_uuid" validate:"required"`
}

// CollectResp 收藏响应
type CollectResp struct {
	Code    int     `json:"code"`
	Collect Collect `json:"collect"`
}

// CollectListResp 收藏列表响应
type CollectListResp struct {
	Code        int       `json:"code"`
	CollectList []Collect `json:"collect_list"`
}
