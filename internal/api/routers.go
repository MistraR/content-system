package api

import (
	"content-system/internal/services"
	"github.com/gin-gonic/gin"
)

const (
	rootPath   = "/api/"
	noAuthPath = "/out/api/"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := services.NewCmsApp()
	session := NewSessionAuth()
	//该group下所有接口都需要经过session鉴权
	root := r.Group(rootPath).Use(session.Auth)
	{
		// /api/cms/hello
		root.GET("/cms/hello", cmsApp.Hello)
		root.POST("/cms/content/create", cmsApp.ContentCreate)
	}

	noAuth := r.Group(noAuthPath)
	{
		// /out/api/cms/register
		noAuth.POST("/cms/register", cmsApp.Register)
		noAuth.POST("/cms/login", cmsApp.Login)
	}
}
