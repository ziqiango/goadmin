package core

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)
var MyDB *gorm.DB

func MysqlDriveInit()  {
	db, err := gorm.Open("mysql", viper.GetString("mysql.url"))
	if err != nil {
		panic(err)
	}

	db.DB().SetConnMaxLifetime(8*time.Second)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(50)
	MyDB = db
}
