package service

import (
	"AssetBot/model"
	"AssetBot/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var (
	err error
	db  *gorm.DB
)

func createTable(db *gorm.DB) {
	if err := db.AutoMigrate(&model.ChainTrade{}, &model.FoodTrade{}); err != nil {
		log.Printf("建表时出现异常", err)
	}

	log.Println("数据库建表成功...")
}

func GetDB() *gorm.DB {
	if db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s",
			utils.Config.Datasource.UserName,
			utils.Config.Datasource.Password,
			utils.Config.Datasource.Host,
			utils.Config.Datasource.Database,
			utils.Config.Datasource.Charset, utils.Config.Datasource.Loc)
		fmt.Println("dsn", dsn)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: gormlogger.Default.LogMode(gormlogger.Info),
		})
		if err != nil {
			panic("failed to connect database")
		}
		log.Println("数据库初始化成功...")
		createTable(db)
	}
	return db
}
