package method

import (
	"LingDei-Middleware/model"
	"LingDei-Middleware/utils"
	"errors"
)

// AddVideo 添加Video
func AddVideo(video model.Video) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(video); err != nil {
		return err
	}

	// 检查 category_uuid 是否存在
	if flag, _ := CheckCategoryExist(video.Category_UUID); !flag {
		return errors.New("category_uuid 不存在")
	}

	if err := db.Create(&video).Error; err != nil {
		return err
	}

	return nil
}

// GetVideoList 获取Video列表
func GetVideoList(category_uuid, user_uuid string, page, page_size int) ([]model.Video, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var videos []model.Video

	if category_uuid != "" {
		if err := db.Where("category_uuid = ?", category_uuid).Offset((page - 1) * page_size).Limit(page_size).Find(&videos).Error; err != nil {
			return nil, err
		}
	} else if user_uuid != "" {
		if err := db.Where("author_uuid = ?", user_uuid).Offset((page - 1) * page_size).Limit(page_size).Find(&videos).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Offset((page - 1) * page_size).Limit(page_size).Find(&videos).Error; err != nil {
			return nil, err
		}
	}

	return videos, nil
}

// GetVideo 获取Video
func GetVideo(uuid string) (model.Video, error) {
	db, err := getDB()
	if err != nil {
		return model.Video{}, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var video model.Video

	if err := db.Where("uuid = ?", uuid).First(&video).Error; err != nil {
		return model.Video{}, err
	}

	return video, nil
}

// DeleteVideo 删除Video
func DeleteVideo(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var video model.Video

	if err := db.Where("uuid = ?", uuid).First(&video).Error; err != nil {
		return err
	}

	if err := utils.DeleteFile("videos/" + video.UUID + ".mp4"); err != nil {
		return err
	}

	if err := db.Delete(&video).Error; err != nil {
		return err
	}

	return nil
}

// 检查Video_uuid是否存在
func CheckVideoExist(Video_uuid string) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var video model.Video

	if err := db.Where("uuid = ?", Video_uuid).First(&video).Error; err != nil {
		return false, err
	}

	return true, nil
}

// UpdateVideo 更新Video
func UpdateVideo(video model.Video) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 检查Video_uuid是否存在
	if flag, err := CheckVideoExist(video.UUID); !flag || err != nil {
		return err
	}

	if err := db.Model(&video).Where("uuid = ?", video.UUID).Updates(&video).Error; err != nil {
		return err
	}

	return nil
}

// GetVideoListByUserUUID 获取用户相关的视频列表
func GetVideoListByUserUUID(user_uuid string) ([]model.Video, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var videos []model.Video

	if err := db.Where("author_uuid = ?", user_uuid).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

// AddVideoViewsCount 添加视频播放量
func AddVideoViewsCount(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var video model.Video

	if err := db.Where("uuid = ?", uuid).First(&video).Error; err != nil {
		return err
	}

	if err := db.Model(&video).Where("uuid = ?", uuid).Update("views", video.Views+1).Error; err != nil {
		return err
	}

	return nil
}

// GetFollowVideos 获取我关注的用户的视频
func GetFollowVideos(user_uuid string, page, page_size int) ([]model.Video, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var followList []model.Follow
	var videoList []model.Video

	if err := db.Where("user_uuid = ?", user_uuid).Find(&followList).Error; err != nil {
		return nil, err
	}

	for _, follow := range followList {
		var video model.Video

		if err := db.Where("author_uuid = ?", follow.Follow_UUID).Offset((page - 1) * page_size).Limit(page_size).Find(&video).Error; err != nil {
			return nil, err
		}

		videoList = append(videoList, video)
	}

	return videoList, nil
}

// SearchVideo 搜索视频
func SearchVideo(keyword string, page, page_size int) ([]model.Video, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var videos []model.Video

	// 分页查询
	if err := db.Where("name LIKE ?", "%"+keyword+"%").Offset((page - 1) * page_size).Limit(page_size).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}
