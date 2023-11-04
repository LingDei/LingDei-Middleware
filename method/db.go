package method

import (
	"LingDei-Middleware/config"
	"LingDei-Middleware/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
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
	db.AutoMigrate(&model.Profile{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Video{})
	db.AutoMigrate(&model.Like{})
	db.AutoMigrate(&model.Collect{})
	db.AutoMigrate(&model.Follow{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Barrage{})

	// 执行原生 SQL 以创建外键
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.videos ADD CONSTRAINT fk_users_video FOREIGN KEY (author_uuid) REFERENCES ` + config.DB_NAME + `.users (id) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.likes ADD CONSTRAINT fk_users_like FOREIGN KEY (user_uuid) REFERENCES ` + config.DB_NAME + `.users (id) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.likes ADD CONSTRAINT fk_videos_like FOREIGN KEY (video_uuid) REFERENCES ` + config.DB_NAME + `.videos (uuid) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.collects ADD CONSTRAINT fk_users_collect FOREIGN KEY (user_uuid) REFERENCES ` + config.DB_NAME + `.users (id) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.collects ADD CONSTRAINT fk_videos_collect FOREIGN KEY (video_uuid) REFERENCES ` + config.DB_NAME + `.videos (uuid) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.follows ADD CONSTRAINT fk_users_follow FOREIGN KEY (user_uuid) REFERENCES ` + config.DB_NAME + `.users (id) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.follows ADD CONSTRAINT fk_follows_follow FOREIGN KEY (follow_uuid) REFERENCES ` + config.DB_NAME + `.users (id) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.comments ADD CONSTRAINT fk_users_comment FOREIGN KEY (user_uuid) REFERENCES ` + config.DB_NAME + `.users (id) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.comments ADD CONSTRAINT fk_videos_comment FOREIGN KEY (video_uuid) REFERENCES ` + config.DB_NAME + `.videos (uuid) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.barrages ADD CONSTRAINT fk_users_barrage FOREIGN KEY (user_uuid) REFERENCES ` + config.DB_NAME + `.users (id) ON DELETE CASCADE ON UPDATE CASCADE;`)
	db.Exec(`ALTER TABLE ` + config.DB_NAME + `.barrages ADD CONSTRAINT fk_videos_barrage FOREIGN KEY (video_uuid) REFERENCES ` + config.DB_NAME + `.videos (uuid) ON DELETE CASCADE ON UPDATE CASCADE;`)
}
