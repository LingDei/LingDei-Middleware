package method

import (
	"LingDei-Middleware/model"
)

// AddComment 添加Comment
func AddComment(comment model.Comment) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(comment); err != nil {
		return err
	}

	if err := db.Create(&comment).Error; err != nil {
		return err
	}

	return nil
}

// GetCommentList 获取Comment列表
func GetCommentList(video_uuid string, page, page_size int) ([]model.Comment, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var comments []model.Comment

	if err := db.Where("video_uuid = ?", video_uuid).Offset((page - 1) * page_size).Limit(page_size).Find(&comments).Error; err != nil {
		return nil, err
	}

	// 补充 User Profile 信息
	for i := 0; i < len(comments); i++ {
		profile, err := GetProfile(comments[i].User_UUID)
		if err != nil {
			return nil, err
		}

		comments[i].Profile = profile
	}

	return comments, nil
}

// GetCommentCount 获取Comment数量
func GetCommentCount(video_uuid string) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Model(&model.Comment{}).Where("video_uuid = ?", video_uuid).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

// GetComment 获取Comment
func GetComment(uuid string) (model.Comment, error) {
	db, err := getDB()
	if err != nil {
		return model.Comment{}, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var comment model.Comment

	if err := db.Where("uuid = ?", uuid).First(&comment).Error; err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

// DeleteComment 删除Comment
func DeleteComment(uuid, user_uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("uuid = ? AND user_uuid = ?", uuid, user_uuid).Delete(&model.Comment{}).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCommentByUUID 删除Comment
func DeleteCommentByUUID(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("uuid = ?", uuid).Delete(&model.Comment{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateComment 更新Comment
func UpdateComment(comment model.Comment) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(comment); err != nil {
		return err
	}

	if err := db.Save(&comment).Error; err != nil {
		return err
	}

	return nil
}
