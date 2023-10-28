package model

type Category struct {
	UUID string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Name string `json:"name" validate:"required" gorm:"unique"`
}

type CategoryResp struct {
	Code     int      `json:"code"`
	Category Category `json:"category"`
}

type CategoryListResp struct {
	Code         int        `json:"code"`
	CategoryList []Category `json:"category_list"`
}
