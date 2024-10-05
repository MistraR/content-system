package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name" binding:"required"`
}

type HelloResponse struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Hello(ctx *gin.Context) {
	var req HelloRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("Hello begin")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &HelloResponse{
			Message: fmt.Sprintf("hello %s", req.Name),
		},
	})
	fmt.Println("Hello end")
}
