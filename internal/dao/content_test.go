package dao

import (
	"content-system/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

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
	return mysqlDB
}

// 单元测试
func TestContentDao_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		content model.ContentDetail
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{db: connectDB()},
			args: args{
				content: model.ContentDetail{
					Title: "标题",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ContentDao{
				db: tt.fields.db,
			}
			if err := a.Create(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
