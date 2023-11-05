package model

type Video struct {
	UUID          string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Name          string `json:"name" validate:"required"`
	Author_UUID   string `json:"author_uuid" validate:"required" gorm:"type:varchar(191); not null"`
	Category_UUID string `json:"category_uuid" validate:"required" gorm:"type:varchar(191); not null"`
	URL           string `json:"url" validate:"required"`
	Thumbnail_URL string `json:"thumbnail_url" validate:"required"`
	Views         int    `json:"views" gorm:"default:0"`
	Timestamp     int64  `json:"timestamp" validate:"required" gorm:"type:bigint"`
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

type UploadTokenResp struct {
	Code         int    `json:"code"`
	Video_UUID   string `json:"video_uuid"`
	Upload_Token string `json:"upload_token"`
}
