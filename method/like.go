package method

import (
	"LingDei-Middleware/model"
	"errors"
)

// AddLike 添加Like
func AddLike(like model.Like) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(like); err != nil {
		return err
	}

	// 检查video_uuid和user_uuid是否存在
	if flag, _ := CheckLikeExist(like.Video_UUID, like.User_UUID); flag {
		return errors.New("已经点赞过了")
	}

	if err := db.Create(&like).Error; err != nil {
		return err
	}

	return nil
}

// GetLikeList 获取Like列表
func GetLikeList() ([]model.Like, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var likes []model.Like

	if err := db.Find(&likes).Error; err != nil {
		return nil, err
	}

	// 给Like列表添加Video信息
	for i := 0; i < len(likes); i++ {
		video, err := GetVideo(likes[i].Video_UUID)
		if err != nil {
			return nil, err
		}
		likes[i].Video = video
	}

	return likes, nil
}

// GetLikeCount 获取Like数量
func GetLikeCount(video_uuid string) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Model(&model.Like{}).Where("video_uuid = ?", video_uuid).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

// GetLike 获取Like
func GetLike(video_uuid string, user_uuid string) (model.Like, error) {
	db, err := getDB()
	if err != nil {
		return model.Like{}, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var like model.Like

	if err := db.Where("video_uuid = ? AND user_uuid = ?", video_uuid, user_uuid).First(&like).Error; err != nil {
		return model.Like{}, err
	}

	// 给Like添加Video信息
	video, err := GetVideo(like.Video_UUID)
	if err != nil {
		return model.Like{}, err
	}
	like.Video = video

	return like, nil
}

// DeleteLike 删除Like
func DeleteLike(video_uuid string, user_uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("video_uuid = ? AND user_uuid = ?", video_uuid, user_uuid).Delete(&model.Like{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateLike 更新Like
func UpdateLike(like model.Like) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 检查Video_uuid是否存在
	if flag, err := CheckVideoExist(like.Video_UUID); !flag || err != nil {
		return errors.New("视频不存在")
	}

	if err := db.Model(&like).Where("video_uuid = ? AND user_uuid = ?", like.Video_UUID, like.User_UUID).Updates(&like).Error; err != nil {
		return err
	}

	return nil
}

// GetLikeListByUserUUID 获取用户相关的Like列表
func GetLikeListByUserUUID(user_uuid string) ([]model.Like, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var likes []model.Like

	if err := db.Where("user_uuid = ?", user_uuid).Find(&likes).Error; err != nil {
		return nil, err
	}

	// 给Like列表添加Video信息
	for i := 0; i < len(likes); i++ {
		video, err := GetVideo(likes[i].Video_UUID)
		if err != nil {
			return nil, err
		}
		likes[i].Video = video
	}

	return likes, nil
}

// GetLikeListByVideoUUID 获取视频相关的Like列表
func GetLikeListByVideoUUID(video_uuid string) ([]model.Like, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var likes []model.Like

	if err := db.Where("video_uuid = ?", video_uuid).Find(&likes).Error; err != nil {
		return nil, err
	}

	return likes, nil
}

// CheckLikeExist 检查Like是否存在
func CheckLikeExist(video_uuid string, user_uuid string) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var like model.Like

	if err := db.Where("video_uuid = ? AND user_uuid = ?", video_uuid, user_uuid).First(&like).Error; err != nil {
		return false, err
	}

	return true, nil
}
