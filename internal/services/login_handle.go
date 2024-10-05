package services

import (
	"content-system/internal/dao"
	"content-system/internal/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	SessionId string `json:"session_id"`
	UserId    string `json:"user_id" binding:"required"`
	NickName  string `json:"nick_name" binding:"required"`
}

func (c *CmsApp) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var (
		password = req.Password
	)
	accountDao := dao.NewAccountDao(c.db)
	account, err := accountDao.FirstByUserId(req.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "账号不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "密码错误"})
		return
	}
	sessionID, err := c.generateSessionId(ctx, account.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &LoginResponse{
			SessionId: sessionID,
			UserId:    account.UserId,
			NickName:  account.Nickname,
		},
	})
	return
}

func (c *CmsApp) generateSessionId(ctx context.Context, userId string) (string, error) {
	sessionId := uuid.New().String()
	sessionKey := utils.GetSessionKey(userId)
	err := c.rdb.Set(ctx, sessionKey, sessionId, time.Second*30).Err()
	if err != nil {
		fmt.Printf("rdb set error = %v \n", err)
		return "", err
	}
	authKey := utils.GetAuthKey(sessionId)
	err = c.rdb.Set(ctx, authKey, time.Now().Unix(), time.Second*30).Err()
	if err != nil {
		fmt.Printf("rdb set error = %v \n", err)
		return "", err
	}
	fmt.Println(sessionKey)
	fmt.Println(authKey)
	return sessionId, nil
}
