package method

import (
	"LingDei-Middleware/model"
	"errors"
)

// AddFollow 添加Follow
func AddFollow(follow model.Follow) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(follow); err != nil {
		return err
	}

	// 禁止关注自己
	if follow.Follow_UUID == follow.User_UUID {
		return errors.New("不能关注自己")
	}

	// 检查follow_uuid和user_uuid是否存在
	if flag, _ := CheckFollowExist(follow.Follow_UUID, follow.User_UUID); flag {
		return errors.New("已经关注过该用户了")
	}

	if err := db.Create(&follow).Error; err != nil {
		return err
	}

	return nil
}

// GetFollowListByUserUUID 获取我的关注列表
func GetFollowListByUserUUID(user_uuid string, page, page_size int) ([]model.Follow, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var follows []model.Follow

	if err := db.Where("user_uuid = ?", user_uuid).Find(&follows).Error; err != nil {
		return nil, err
	}

	// 给Follow列表添加Target信息
	for i := 0; i < len(follows); i++ {
		target, err := GetUser(follows[i].Follow_UUID)
		if err != nil {
			return nil, err
		}
		follows[i].User = target
	}

	return follows, nil
}

// GetFanListByFollowUUID 获取我的粉丝列表
func GetFanListByFollowUUID(user_uuid string, page, page_size int) ([]model.Follow, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var follows []model.Follow

	if err := db.Where("follow_uuid = ?", user_uuid).Offset((page - 1) * page_size).Limit(page_size).Find(&follows).Error; err != nil {
		return nil, err
	}

	// 给Follow列表添加Follower信息
	for i := 0; i < len(follows); i++ {
		follower, err := GetUser(follows[i].Follow_UUID)
		if err != nil {
			return nil, err
		}
		follows[i].User = follower
	}

	return follows, nil
}

// CheckFollowExist 检查Follow是否存在
func CheckFollowExist(follow_uuid string, user_uuid string) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var follow model.Follow

	if err := db.Where("follow_uuid = ? AND user_uuid = ?", follow_uuid, user_uuid).First(&follow).Error; err != nil {
		return false, err
	}

	return true, nil
}

// DeleteFollow 删除Follow
func DeleteFollow(follow_uuid string, user_uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("follow_uuid = ? AND user_uuid = ?", follow_uuid, user_uuid).Delete(&model.Follow{}).Error; err != nil {
		return err
	}

	return nil
}

// GetFollowCount 获取某个用户已关注用户的个数
func GetFollowCount(user_uuid string) (int64, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Model(&model.Follow{}).Where("user_uuid = ?", user_uuid).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetFanCount 获取Fan数量
func GetFanCount(follow_uuid string) (int64, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Model(&model.Follow{}).Where("follow_uuid = ?", follow_uuid).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetFollowStatus 获取Follow状态
func GetFollowStatus(follow_uuid string, user_uuid string) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var follow model.Follow

	if err := db.Where("follow_uuid = ? AND user_uuid = ?", follow_uuid, user_uuid).First(&follow).Error; err != nil {
		return false, err
	}

	return true, nil
}
