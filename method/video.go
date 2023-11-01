package method

import (
	"LingDei-Middleware/model"
	"LingDei-Middleware/utils"
	"errors"
	"mime/multipart"
)

// AddVideo 添加Video
func AddVideo(video model.Video, file_header *multipart.FileHeader) error {
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

	// file_header to file
	file, err := file_header.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// 上传视频
	if err := utils.UploadFileByPieces("videos/"+video.UUID+".mp4", file); err != nil {
		return err
	}

	if err := db.Create(&video).Error; err != nil {
		return err
	}

	return nil
}

// GetVideoList 获取Video列表
func GetVideoList(category_uuid string) ([]model.Video, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var videos []model.Video

	if category_uuid != "" {
		if err := db.Find(&videos).Error; err != nil {
			return nil, err
		}
		return videos, nil
	}

	if err := db.Where("category_uuid = ?", category_uuid).Find(&videos).Error; err != nil {
		return nil, err
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
