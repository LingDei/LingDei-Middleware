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

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(category); err != nil {
		return err
	}

	if err := db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

// GetCategoryList 获取Category列表
func GetCategoryList() ([]model.Category, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var categories []model.Category

	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryByUUID 根据UUID获取Category
func GetCategoryByUUID(uuid string) (model.Category, error) {
	db, err := getDB()
	if err != nil {
		return model.Category{}, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var category model.Category

	if err := db.Where("uuid = ?", uuid).First(&category).Error; err != nil {
		return model.Category{}, err
	}

	return category, nil
}

// DeleteCategory 删除Category
func DeleteCategory(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("uuid = ?", uuid).Delete(&model.Category{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateCategory 更新Category
func UpdateCategory(category model.Category) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(category); err != nil {
		return err
	}

	if err := db.Save(&category).Error; err != nil {
		return err
	}

	return nil
}
