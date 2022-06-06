package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:wang200315@tcp(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB

//type StoredUser struct {
//	Token          string
//	UserKey_Stored int64
//	User           entity.User `gorm:"foreignKey:UserKey_Stored"`
//}

func Init() {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db = database
	//自动迁移，保证数据库是最新的
	if err := db.AutoMigrate(&DbUser{}); err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&dbVideo{}); err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&UserFavorite{}); err != nil {
		log.Fatal(err)
	}
}

//测试用, 后续应该更改方案
