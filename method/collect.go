package method

import (
	"LingDei-Middleware/model"
)

// AddCollect 添加Collect
func AddCollect(collect model.Collect) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(collect); err != nil {
		return err
	}

	if err := db.Create(&collect).Error; err != nil {
		return err
	}

	return nil
}

// GetCollectList 获取Collect列表
func GetCollectList() ([]model.Collect, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var collects []model.Collect

	if err := db.Find(&collects).Error; err != nil {
		return nil, err
	}

	return collects, nil
}

// GetCollectCount 获取Collect数量
func GetCollectCount(video_uuid string) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Model(&model.Collect{}).Where("video_uuid = ?", video_uuid).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

// GetCollect 获取Collect
func GetCollect(uuid string) (model.Collect, error) {
	db, err := getDB()
	if err != nil {
		return model.Collect{}, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var collect model.Collect

	if err := db.Where("uuid = ?", uuid).First(&collect).Error; err != nil {
		return model.Collect{}, err
	}

	return collect, nil
}

// DeleteCollect 删除Collect
func DeleteCollect(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var collect model.Collect

	if err := db.Where("uuid = ?", uuid).First(&collect).Error; err != nil {
		return err
	}

	if err := db.Delete(&collect).Error; err != nil {
		return err
	}

	return nil
}

// UpdateCollect 更新Collect
func UpdateCollect(collect model.Collect) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Save(&collect).Error; err != nil {
		return err
	}

	return nil
}

// GetCollectListByUserUUID 获取用户相关的Collect列表
func GetCollectListByUserUUID(user_uuid string) ([]model.Collect, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var collects []model.Collect

	if err := db.Where("user_uuid = ?", user_uuid).Find(&collects).Error; err != nil {
		return nil, err
	}

	return collects, nil
}

// GetCollectListByVideoUUID 获取视频相关的Collect列表
func GetCollectListByVideoUUID(video_uuid string) ([]model.Collect, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var collects []model.Collect

	if err := db.Where("video_uuid = ?", video_uuid).Find(&collects).Error; err != nil {
		return nil, err
	}

	return collects, nil
}

// CheckCollectExist 检查Collect是否存在
func CheckCollectExist(video_uuid string, user_uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var collect model.Collect

	if err := db.Where("video_uuid = ? AND user_uuid = ?", video_uuid, user_uuid).First(&collect).Error; err != nil {
		return err
	}

	return nil
}
