package model

// Like 点赞
type Like struct {
	UUID       string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Video_UUID string `json:"video_uuid" validate:"required"`
	User_UUID  string `json:"user_uuid" validate:"required"`
	Like       bool   `json:"is_like" validate:"required"`
}

// LikeResp 点赞响应
type LikeResp struct {
	Code int  `json:"code"`
	Like Like `json:"like"`
}

// 点赞数量响应
type LikeCountResp struct {
	Code  int `json:"code"`
	Count int `json:"count"`
}

// LikeListResp 点赞列表响应
type LikeListResp struct {
	Code     int    `json:"code"`
	LikeList []Like `json:"like_list"`
}
