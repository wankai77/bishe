package db

import (
	"backend/config/config"
	model "backend/models"
	"fmt"

	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func SetupDBLink() error {
	var err error
	var dbConfig = config.Config.Db
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)

	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	if Db.Error != nil {
		panic(Db.Error)
	}

	Db.AutoMigrate(&model.UserInfo{})
	Db.AutoMigrate(&model.ReleaseInfo{})
	Db.AutoMigrate(&model.MessageInfo{})
	Db.AutoMigrate(&model.ApplicatUser{})
	fmt.Sprintln("loading")
	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
	return nil
}
