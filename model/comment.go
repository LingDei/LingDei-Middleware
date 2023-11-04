package model

// 评论
type Comment struct {
	UUID       string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Video_UUID string `json:"video_uuid" validate:"required" gorm:"type:varchar(191)"`
	User_UUID  string `json:"user_uuid" validate:"required" gorm:"type:varchar(191)"`
	Content    string `json:"content" validate:"required"`
	Timestamp  int64  `json:"timestamp" validate:"required" gorm:"type:bigint"`
}

// CommentListResp 评论列表响应
type CommentListResp struct {
	Code        int       `json:"code"`
	CommentList []Comment `json:"comment_list"`
}

// CommentResp 评论响应
type CommentResp struct {
	Code    int     `json:"code"`
	Comment Comment `json:"comment"`
}

// CommentCountResp 评论数量响应
type CommentCountResp struct {
	Code  int `json:"code"`
	Count int `json:"count"`
}
