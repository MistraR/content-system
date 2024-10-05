package api

import (
	"content-system/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"net/http"
)

const SessionKey = "session_id"

type SessionAuth struct {
	rdb *redis.Client
}

func NewSessionAuth() *SessionAuth {
	s := &SessionAuth{}
	connectRDB(s)
	return s
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionId := ctx.GetHeader(SessionKey)

	if sessionId == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "session is null")
	}
	// session 鉴权
	authKey := utils.GetAuthKey(sessionId)
	loginTime, err := s.rdb.Get(ctx, authKey).Result()
	if err != nil && err != redis.Nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "session auth error")
	}
	if loginTime == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "session auth failure")
	}
	fmt.Println("begin sessionId = ", sessionId)
	ctx.Next()
	fmt.Println("end sessionId = ", sessionId)
}

func connectRDB(app *SessionAuth) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.132.113.82:6378",
		Password: "pP6vY4sD", // no password set
		DB:       0,          // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	app.rdb = rdb
}
