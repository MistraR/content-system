package main

import (
	"content-system/internal/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.CmsRouters(r)
	err := r.Run()
	if err != nil {
		fmt.Print("r run error = ", err)
		return
	}
}
