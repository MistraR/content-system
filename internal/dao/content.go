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

func (a *ContentDao) Update(content model.ContentDetail) error {
	//为空的字段不会更新
	if err := a.db.Where("id=?", content.ID).Updates(&content).Error; err != nil {
		fmt.Printf("ContentDao update error = [%v]", err)
		return err
	}
	return nil
}

func (a *ContentDao) IsExist(contentId int64) (bool, error) {
	var content = model.ContentDetail{}
	err := a.db.Where("id=?", contentId).First(&content).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Printf("ContentDao isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (a *ContentDao) Delete(contentId int64) error {
	if err := a.db.Where("id=?", contentId).Delete(&model.ContentDetail{}).Error; err != nil {
		fmt.Printf("ContentDao Delete error = [%v]", err)
		return err
	}
	return nil
}

type QueryParam struct {
	ID       int64
	Author   string
	Title    string
	Page     int
	PageSize int
}

func (a *ContentDao) Query(param *QueryParam) ([]*model.ContentDetail, int64, error) {
	query := a.db.Model(&model.ContentDetail{})
	if param.ID != 0 {
		query = query.Where("id=?", param.ID)
	}
	if param.Author != "" {
		query = query.Where("author=?", param.Author)
	}
	if param.Title != "" {
		query = query.Where("title=?", param.Title)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var page, pageSize = 1, 10
	if param.Page > 0 {
		page = param.Page
	}
	if param.PageSize > 0 {
		pageSize = param.PageSize
	}
	offset := (page - 1) * pageSize
	var data []*model.ContentDetail
	if err := query.Offset(offset).Limit(pageSize).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, total, nil
}
