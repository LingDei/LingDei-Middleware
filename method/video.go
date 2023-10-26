package method

import (
	"LingDei_Middleware/model"
	"LingDei_Middleware/utils"
	"fmt"
	"io"
	"mime/multipart"
)

// AddVideo 添加Video
func AddVideo(video model.Video, file io.Reader) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	// 校验数据
	if err := validate.Struct(video); err != nil {
		return err
	}

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
func GetVideoList() ([]model.Video, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	var videos []model.Video

	if err := db.Find(&videos).Error; err != nil {
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
func CheckVideoExist(Video_uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	var video model.Video

	if err := db.Where("uuid = ?", Video_uuid).First(&video).Error; err != nil {
		return err
	}

	return nil
}

// UpdateVideo 更新Video
func UpdateVideo(video model.Video, file_mul *multipart.FileHeader) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	if file_mul != nil {
		file, err := file_mul.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		// 上传文件
		if file != nil {
			if err := utils.UploadFileByPieces("videos/"+video.UUID+".mp4", file); err != nil {
				return err
			}
		}
	}

	// 检查Video_uuid是否存在
	if err := CheckVideoExist(video.UUID); err != nil {
		return err
	}

	fmt.Println(video)
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

	var videos []model.Video

	if err := db.Where("author_uuid = ?", user_uuid).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}
