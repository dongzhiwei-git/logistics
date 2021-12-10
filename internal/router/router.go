package router

import (
	"fmt"
	"inherited/internal/api"
	"inherited/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	admin *gin.RouterGroup
	store *gin.RouterGroup
)

func InitRouter() {
	var r *gin.Engine
	r = gin.Default()
	r.Use(pkg.Cors())

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
		admin.GET("/get-center-info", api.GetCenterInfo)
		admin.GET("/get-input-info", api.GetInputInfo)
		admin.GET("/get-output-info", api.GetOutputInfo)
		// 获取备件等级表
		admin.GET("/get-prolevel-info", api.GetProductLevel)
		admin.GET("/update-level", api.UpdateProductLevel)
	}

	// setup listen
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("run.sh failed: %v\n", err)
		return

	}

}
