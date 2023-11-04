package model

// 弹幕
type Barrage struct {
	UUID       string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Video_UUID string `json:"video_uuid" validate:"required" gorm:"type:varchar(191)"`
	User_UUID  string `json:"user_uuid" validate:"required" gorm:"type:varchar(191)"`
	Content    string `json:"content" validate:"required"`
	Second     int64  `json:"second" validate:"required" gorm:"type:bigint"`
	Timestamp  int64  `json:"timestamp" validate:"required" gorm:"type:timestamp"`
}
