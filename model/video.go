package model

type Video struct {
	UUID          string `json:"id" validate:"required" gorm:"primaryKey"`
	Name          string `json:"name" validate:"required"`
	Category      string `json:"category" validate:"required"`
	Thumbnail_URL string `json:"thumbnail_url" validate:"required"`
	Views         int    `json:"views" validate:"required"`
	Author_ID     string `json:"author_id" validate:"required"`
}

type ContractResp struct {
	Code   int    `json:"code"`
	URL    string `json:"url"`
	Video  Video  `json:"video"`
	Base64 string `json:"base64"`
}

type ContractListResp struct {
	Code   int     `json:"code"`
	Videos []Video `json:"videos"`
}
