package method

import (
	"LingDei-Middleware/config"
	"LingDei-Middleware/model"
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
	db.AutoMigrate(&model.Video{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Like{})
	db.AutoMigrate(&model.Collect{})
}
