package method

import (
	"LingDei-Middleware/model"
)

// AddBarrage 添加Barrage
func AddBarrage(barrage model.Barrage) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(barrage); err != nil {
		return err
	}

	if err := db.Create(&barrage).Error; err != nil {
		return err
	}

	return nil
}

// GetBarrageList 获取Barrage列表
func GetBarrageList(video_uuid string, page, page_size int) ([]model.Barrage, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var barrages []model.Barrage

	if err := db.Where("video_uuid = ?", video_uuid).Offset((page - 1) * page_size).Limit(page_size).Find(&barrages).Error; err != nil {
		return nil, err
	}

	return barrages, nil
}

// GetBarrageCount 获取Barrage数量
func GetBarrageCount(video_uuid string) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Model(&model.Barrage{}).Where("video_uuid = ?", video_uuid).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

// DeleteBarrage 删除Barrage
func DeleteBarrage(uuid string, user_uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var barrage model.Barrage

	if err := db.Where("uuid = ? AND user_uuid = ?", uuid, user_uuid).First(&barrage).Error; err != nil {
		return err
	}

	if err := db.Delete(&barrage).Error; err != nil {
		return err
	}

	return nil
}

// DeleteBarrageByUUID 删除Barrage
func DeleteBarrageByUUID(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("uuid = ?", uuid).Delete(&model.Barrage{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateBarrage 更新Barrage
func UpdateBarrage(barrage model.Barrage) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(barrage); err != nil {
		return err
	}

	if err := db.Save(&barrage).Error; err != nil {
		return err
	}

	return nil
}

// GetBarrage 获取Barrage
func GetBarrage(uuid string) (model.Barrage, error) {
	db, err := getDB()
	if err != nil {
		return model.Barrage{}, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var barrage model.Barrage

	if err := db.Where("uuid = ?", uuid).First(&barrage).Error; err != nil {
		return model.Barrage{}, err
	}

	return barrage, nil
}
