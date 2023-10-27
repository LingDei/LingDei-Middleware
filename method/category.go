package method

import (
	"LingDei-Middleware/model"
)

// AddCategory 添加Category
func AddCategory(category model.Category) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	// 校验数据
	if err := validate.Struct(category); err != nil {
		return err
	}

	if err := db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}
