package services

import (
	"content-system/internal/dao"
	"content-system/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
}

type RegisterResponse struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 密码加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
	// 校验用户是否存在
	accountDao := dao.NewAccountDao(c.db)
	exist, err := accountDao.IsExist(req.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "账号已存在",
		})
		return
	}
	// 用户信息持久化
	if err := accountDao.Create(model.Account{
		UserId:    req.UserId,
		Password:  hashedPassword,
		Nickname:  req.NickName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("RegisterRequest = %+v, hashedPassword=[%s] \n", req, hashedPassword)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &RegisterResponse{
			Message: fmt.Sprintf("注册成功"),
		},
	})
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("密码加密错误=%v", err)
		return "", err
	}
	return string(hashedPassword), err
}
