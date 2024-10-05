package services

import (
	"content-system/internal/dao"
	"content-system/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ContentCreateRequest struct {
	Title          string        `json:"title" binding:"required"`       //内容标题
	Description    string        `json:"description" binding:"required"` //描述
	Author         string        `json:"author" binding:"required"`      //作者
	VideoURL       string        `json:"video_url"`                      //视频url
	Thumbnail      string        `json:"thumbnail"`                      //封面图url
	Category       string        `json:"category"`                       //分类
	Duration       time.Duration `json:"duration"`                       //时长
	Resolution     string        `json:"resolution"`                     //分辨率
	FileSize       int64         `json:"file_size"`                      //文件大小
	Format         string        `json:"format"`                         //格式
	Quality        int           `json:"quality"`                        //视频质量
	ApprovalStatus int           `json:"approval_status"`                //审核状态
}

type ContentCreateResponse struct {
	Message string `json:"message"`
}

func (c *CmsApp) ContentCreate(ctx *gin.Context) {
	var req ContentCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	contentDao := dao.NewContentDao(c.db)
	err := contentDao.Create(model.ContentDetail{
		Title:          req.Title,
		Description:    req.Description,
		Author:         req.Author,
		VideoURL:       req.VideoURL,
		Thumbnail:      req.Thumbnail,
		Category:       req.Category,
		Duration:       req.Duration,
		Resolution:     req.Resolution,
		FileSize:       req.FileSize,
		Format:         req.Format,
		Quality:        req.Quality,
		ApprovalStatus: req.ApprovalStatus,
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentCreateResponse{
			Message: "ok",
		},
	})
}
