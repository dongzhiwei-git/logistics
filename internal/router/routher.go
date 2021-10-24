package router

import (
	"fmt"
	"inherited/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

var(
	admin *gin.RouterGroup
	store *gin.RouterGroup

)

func InitRouter() {
	var r *gin.Engine
	r = gin.Default()

	// to solve the cross domain
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	// 管理员相关
	admin = r.Group("/admin")
	{
		admin.POST("/reg", api.CreateAdminUser)
		// 仓库配送到分仓库
		admin.GET("/get-store-info", api.GetStoreInfo)
	}





	// setup listen
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return

	}

}
