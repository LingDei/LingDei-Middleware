package model

// Follow
type Follow struct {
	UUID        string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Target_UUID string `json:"target_uuid" validate:"required" gorm:"type:varchar(191)"`
	User_UUID   string `json:"user_uuid" validate:"required" gorm:"type:varchar(191)"`
	User        User   `json:"user" gorm:"-"`
}

// FollowResp
type FollowResp struct {
	Code   int    `json:"code"`
	Follow Follow `json:"follow"`
}

type FollowCountResp struct {
	Code  int `json:"code"`
	Count int `json:"count"`
}

type FollowStatusResp struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
}

type FollowListResp struct {
	Code       int    `json:"code"`
	FollowList []Like `json:"like_list"`
}
