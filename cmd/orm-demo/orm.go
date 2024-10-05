package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID        int64     `gorm:"column:id;primary_key"`
	UserId    string    `gorm:"column:user_id"`
	Password  string    `gorm:"column:password"`
	Nickname  string    `gorm:"column:nickname"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// 这样定义表名也可以
func (a Account) TableName() string {
	//分库分表时可以动态计算表名
	table := fmt.Sprintf("account_%d", a.ID%10)
	fmt.Println(table)
	return "account"
}

func main() {
	db := connectDB()
	var accounts []Account
	if err := db.Table("account").Find(&accounts).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(accounts)
	var account Account
	if err := db.Where("id=?", 1).First(&account).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(account)
}

func connectDB() *gorm.DB {
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
	//if env == "test" {
	//	mysqlDB = mysqlDB.Debug()
	//}
	return mysqlDB.Debug()
}
