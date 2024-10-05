package services

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	connectDB(app)
	connectRDB(app)
	return app
}

func connectDB(app *CmsApp) {
	dsn := "root:YQSaGPzeRRCvDx2mOHVMOw@tcp(101.132.113.82:3306)/cms_account?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	app.db = mysqlDB
}

func connectRDB(app *CmsApp) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.132.113.82:6378",
		Password: "pP6vY4sD", // no password set
		DB:       0,          // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	app.rdb = rdb
}
