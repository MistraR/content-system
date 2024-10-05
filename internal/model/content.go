package model

import "time"

type ContentDetail struct {
	ID             int64         `gorm:"column:id;primary_key"`  //ID
	Title          string        `gorm:"column:title"`           //内容标题
	Description    string        `gorm:"column:description"`     //描述
	Author         string        `gorm:"column:author"`          //作者
	VideoURL       string        `gorm:"column:video_url"`       //视频url
	Thumbnail      string        `gorm:"column:thumbnail"`       //封面图url
	Category       string        `gorm:"column:category"`        //分类
	Duration       time.Duration `gorm:"column:duration"`        //时长
	Resolution     string        `gorm:"column:resolution"`      //分辨率
	FileSize       int64         `gorm:"column:file_size"`       //文件大小
	Format         string        `gorm:"column:format"`          //格式
	Quality        int           `gorm:"column:quality"`         //视频质量
	ApprovalStatus int           `gorm:"column:approval_status"` //审核状态
	UpdatedAt      time.Time     `gorm:"column:updated_at"`      //更新时间
	CreatedAt      time.Time     `gorm:"column:created_at"`      //创建时间
}

func (a ContentDetail) TableName() string {
	table := "cms_account.content_detail"
	return table
}
