package common

//数据库初始化
import (
	"Register-Login-Project/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "zxr:zxr123@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	DB = db
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Upload{})
	return db
}
func GetDB() *gorm.DB {
	return DB
}
