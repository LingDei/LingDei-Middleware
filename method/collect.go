package method

import (
	"LingDei-Middleware/model"
	"errors"
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

	// 检查video_uuid和user_uuid是否存在
	if flag, _ := CheckCollectExist(collect.Video_UUID, collect.User_UUID); flag {
		return errors.New("已经收藏过了")
	}

	if err := db.Create(&collect).Error; err != nil {
		return err
	}

	return nil
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

	// 给Collect添加Video信息
	video, err := GetVideo(collect.Video_UUID)
	if err != nil {
		return model.Collect{}, err
	}
	collect.Video = video

	return collect, nil
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

	// 给Collect列表添加Video信息
	for i := 0; i < len(collects); i++ {
		video, err := GetVideo(collects[i].Video_UUID)
		if err != nil {
			return nil, err
		}
		collects[i].Video = video
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

// DeleteCollect 删除Collect
func DeleteCollect(video_uuid, user_uuid string) error {
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
func GetCollectListByUserUUID(user_uuid string, page, page_size int) ([]model.Collect, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var collects []model.Collect

	if err := db.Where("user_uuid = ?", user_uuid).Offset((page - 1) * page_size).Limit(page_size).Find(&collects).Error; err != nil {
		return nil, err
	}

	// 给Collect列表添加Video信息
	for i := 0; i < len(collects); i++ {
		video, err := GetVideo(collects[i].Video_UUID)
		if err != nil {
			return nil, err
		}
		collects[i].Video = video
	}

	return collects, nil
}

// GetCollectCountByUserUUID 获取用户相关的Collect数量
func GetCollectCountByUserUUID(user_uuid string) (int64, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Model(&model.Collect{}).Where("user_uuid = ?", user_uuid).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
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
func CheckCollectExist(video_uuid string, user_uuid string) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var collect model.Collect

	if err := db.Where("video_uuid = ? AND user_uuid = ?", video_uuid, user_uuid).First(&collect).Error; err != nil {
		return false, err
	}

	return true, nil
}
