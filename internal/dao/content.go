package dao

import (
	"content-system/internal/model"
	"fmt"
	"gorm.io/gorm"
)

type ContentDao struct {
	db *gorm.DB
}

func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{db: db}
}

func (a *ContentDao) Create(content model.ContentDetail) error {
	if err := a.db.Create(&content).Error; err != nil {
		fmt.Printf("ContentDao Create = [%v]", err)
		return err
	}
	return nil
}
