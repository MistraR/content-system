package services

import (
	"content-system/internal/dao"
	"content-system/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContentQueryRequest struct {
	ID       int64  `json:"id"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type ContentQueryResponse struct {
	Message  string                 `json:"message"`
	Contents []*model.ContentDetail `json:"contents"`
	Total    int64                  `json:"total"`
}

func (c *CmsApp) ContentQuery(ctx *gin.Context) {
	var req ContentQueryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	contentDao := dao.NewContentDao(c.db)
	contentList, total, err := contentDao.Query(&dao.QueryParam{
		ID:       req.ID,
		Author:   req.Author,
		Title:    req.Title,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//可以转换为特定的VO
	//contents := make([]model.ContentDetail, 0, len(contentList))
	//for _, content := range contentList {
	//	contents = append(contents, model.ContentDetail{
	//		ID:             content.ID,
	//		Title:          content.Title,
	//		Description:    content.Description,
	//		Author:         content.Author,
	//		VideoURL:       content.VideoURL,
	//		Thumbnail:      content.Thumbnail,
	//		Category:       content.Category,
	//		Duration:       content.Duration,
	//		Resolution:     content.Resolution,
	//		FileSize:       content.FileSize,
	//		Format:         content.Format,
	//		Quality:        content.Quality,
	//		ApprovalStatus: content.ApprovalStatus,
	//		UpdatedAt:      content.UpdatedAt,
	//		CreatedAt:      content.CreatedAt,
	//	})
	//}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentQueryResponse{
			Message:  "ok",
			Contents: contentList,
			Total:    total,
		},
	})
}
