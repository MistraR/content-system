package services

import (
	"content-system/internal/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContentDeleteRequest struct {
	ID int64 `json:"id" binding:"required"` //内容标题
}

type ContentDeleteResponse struct {
	Message string `json:"message"`
}

func (c *CmsApp) ContentDelete(ctx *gin.Context) {
	var req ContentDeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	contentDao := dao.NewContentDao(c.db)
	exist, err := contentDao.IsExist(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "内容不存在"})
		return
	}
	err = contentDao.Delete(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentDeleteResponse{
			Message: "ok",
		},
	})
}
