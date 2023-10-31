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

	// 检查target_uuid和user_uuid是否存在
	if flag, _ := CheckLikeExist(follow.Target_UUID, follow.User_UUID); flag {
		return errors.New("已经关注过该用户了")
	}

	if err := db.Create(&follow).Error; err != nil {
		return err
	}

	return nil
}

// GetFollowListByUserUUID 获取我的关注列表
func GetFollowListByUserUUID(user_uuid string) ([]model.Follow, error) {
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
		target, err := GetUser(follows[i].Target_UUID)
		if err != nil {
			return nil, err
		}
		follows[i].User = target
	}

	return follows, nil
}

// GetFanListByTargetUUID 获取我的粉丝列表
func GetFanListByTargetUUID(user_uuid string) ([]model.Follow, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var follows []model.Follow

	if err := db.Where("target_uuid = ?", user_uuid).Find(&follows).Error; err != nil {
		return nil, err
	}

	// 给Follow列表添加Follower信息
	for i := 0; i < len(follows); i++ {
		follower, err := GetUser(follows[i].Target_UUID)
		if err != nil {
			return nil, err
		}
		follows[i].User = follower
	}

	return follows, nil
}

// CheckFollowExist 检查Follow是否存在
func CheckFollowExist(target_uuid string, user_uuid string) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var follow model.Follow

	if err := db.Where("target_uuid = ? AND user_uuid = ?", target_uuid, user_uuid).First(&follow).Error; err != nil {
		return false, err
	}

	return true, nil
}

// DeleteFollow 删除Follow
func DeleteFollow(target_uuid string, user_uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("target_uuid = ? AND user_uuid = ?", target_uuid, user_uuid).Delete(&model.Follow{}).Error; err != nil {
		return err
	}

	return nil
}

// GetFollowCount 获取Follow数量
func GetFollowCount(user_uuid string) (int, error) {
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

	return int(count), nil
}

// GetFanCount 获取Fan数量
func GetFanCount(target_uuid string) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Model(&model.Follow{}).Where("target_uuid = ?", target_uuid).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

// GetFollowStatus 获取Follow状态
func GetFollowStatus(target_uuid string, user_uuid string) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var follow model.Follow

	if err := db.Where("target_uuid = ? AND user_uuid = ?", target_uuid, user_uuid).First(&follow).Error; err != nil {
		return false, err
	}

	return true, nil
}
