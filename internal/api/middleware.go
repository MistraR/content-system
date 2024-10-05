package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SessionKey = "session_id"

type SessionAuth struct {
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionId := ctx.GetHeader(SessionKey)
	// TODO 鉴权
	if sessionId == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "session is null")
	}
	fmt.Println("begin sessionId = ", sessionId)
	ctx.Next()
	fmt.Println("end sessionId = ", sessionId)
}
