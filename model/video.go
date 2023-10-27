package model

type Video struct {
	UUID          string `json:"id" validate:"required" gorm:"primaryKey"`
	Name          string `json:"name" validate:"required"`
	Category_UUID string `json:"category_uuid" validate:"required"`
	URL           string `json:"url" validate:"required"`
	Thumbnail_URL string `json:"thumbnail_url" validate:"required"`
	Views         int    `json:"views" gorm:"default:0"`
	Author_UUID   string `json:"author_id" validate:"required"`
}

type VideoResp struct {
	Code   int    `json:"code"`
	Video  Video  `json:"video"`
	Base64 string `json:"base64"`
}

type VideoListResp struct {
	Code       int     `json:"code"`
	Video_List []Video `json:"video_list"`
}
