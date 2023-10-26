package method

import (
	"LingDei_Middleware/config"
	"LingDei_Middleware/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	return db, err
}

func init() {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 迁移 schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Discipline{})
	db.AutoMigrate(&model.Lesson{})
	db.AutoMigrate(&model.Subscribe{})
	db.AutoMigrate(&model.WhiteList{})
	db.AutoMigrate(&model.ScheduleItem{})
	db.AutoMigrate(&model.Video{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.GitLib{})
}
