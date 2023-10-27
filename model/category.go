package model

type Category struct {
	UUID string `json:"id" validate:"required" gorm:"primaryKey"`
	Name string `json:"name" validate:"required"`
}

type CategoryResp struct {
	Code     int      `json:"code"`
	Category Category `json:"category"`
}

type CategoryListResp struct {
	Code         int        `json:"code"`
	CategoryList []Category `json:"category_list"`
}
