package api

import (
	"fmt"
	"log"
	"net/http"

	"inherited/internal/models"
	"inherited/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateAdminUser(ctx *gin.Context) {
	//Parameter parsing
	adminUser := models.SysUser{}
	err := ctx.BindJSON(&adminUser)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")

		return
	}

	userName := adminUser.UserName
	password := adminUser.Password
	if userName == "" || password == "" {
		log.Printf("userName or password is null")

		return
	}

	sysUser := new(services.SysUser)
	err = sysUser.CreateSysUser(adminUser.UserName, adminUser.Password)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"date": "success",
	})

	return
}

// GetStoreInfo 总仓到分仓info
func GetStoreInfo(ctx *gin.Context){
	//Parameter parsing
	info := models.StoreInfo{}
	err := ctx.BindJSON(&info)
	if err != nil {
		fmt.Printf("[api.GetStoreInfo], Parameter parsing error")
	}

	store := new(services.Store)
	storeInfo, err := store.GetStoreInfo()
	if err != nil {
		fmt.Printf("[api.GetStoreInfo], err: %v", err)

		return
	}

	fmt.Println(storeInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date" : storeInfo,
	})

	return
}

func GetCenterInfo(ctx *gin.Context){
	//Parameter parsing
	info := models.CenterInfo{}
	err := ctx.BindJSON(&info)
	if err != nil {
		fmt.Printf("[api.GetCenterInfo], Parameter parsing error")
	}

	center := new(services.Center)
	centerInfo, err := center.GetCenterInfo()
	if err != nil {
		fmt.Printf("[api.GetCenterInfo], err: %v", err)

		return
	}

	fmt.Println(centerInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date" : centerInfo,
	})

	return
}
