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

// 总仓到分仓info
func GetStoreInfo(ctx *gin.Context) {
	//Parameter parsing
	info := models.StoreInfo{}
	err := ctx.ShouldBind(&info)
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
		"date":   storeInfo,
	})

	return
}

func GetCenterInfo(ctx *gin.Context) {
	//Parameter parsing
	info := models.CenterInfo{}
	err := ctx.ShouldBindJSON(&info)
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
		"date":   centerInfo,
	})

	return
}

// 得到入库表info
func GetInputInfo(ctx *gin.Context) {
	//Parameter parsing
	info := models.Input{}
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		fmt.Printf("[api.GetInputInfo], Parameter parsing error")
	}

	input := new(services.Input)
	inputInfo, err := input.GetInputInfo()
	if err != nil {
		fmt.Printf("[api.GetInputInfo], err: %v", err)

		return
	}

	fmt.Println(inputInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date":   inputInfo,
	})

	return
}

// 得到出库表info
func GetOutputInfo(ctx *gin.Context) {
	//Parameter parsing
	info := models.Output{}
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		fmt.Printf("[api.GetOutputInfo], Parameter parsing error")
	}

	output := new(services.Output)
	outputInfo, err := output.GetOutputInfo()
	if err != nil {
		fmt.Printf("[api.GetOutputInfo], err: %v", err)

		return
	}

	fmt.Println(outputInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date":   outputInfo,
	})

	return
}

//	得到备件等级info
func GetProductLevel(ctx *gin.Context) {
	level := new(services.ProductLevel)
	levelInfo, err := level.GetProductLevel()
	if err != nil {
		fmt.Printf("[api.GetInputInfo], err: %v", err)

		return
	}

	log.Println(levelInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date":   levelInfo,
	})
}

//	更新备件等级信息
func UpdateProductLevel(ctx *gin.Context) {
	//Parameter parsing
	info := models.ProductLevel{}
	err := ctx.ShouldBindJSON(&info)
	fmt.Println("00000000", info, ctx)
	if err != nil {
		fmt.Printf("[api.GetOutputInfo], Parameter parsing error: %v", err)
		return
	}
	fmt.Println(info)
	if info.Id <= 0 {
		ctx.JSON(http.StatusOK, "id不能小于1")
		ctx.JSON(http.StatusOK, "id不能小于1")
		return
	}
	if info.ForwardDate <= 0 {
		ctx.JSON(http.StatusOK, "ForwardDate不能小于1")
		return
	}
	if info.TotalAmount <= 0 {
		ctx.JSON(http.StatusOK, "TotalAmount不能小于1")
		return
	}
	if err != nil {
		fmt.Printf("[api.UpdateProductLevel], Parameter parsing error")
	}

	level := new(services.ProductLevel)
	productLevel, err := level.UpdateProductLevel(info.Id, info.TotalAmount, info.General, info.ForwardDate)
	level.UpdateLevel(productLevel, info.Id)
	if err != nil {
		fmt.Printf("[api.UpdateProductLevel], err: %v", err)

		return
	}

	log.Println(productLevel)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"level":  productLevel,
	})

	return
}

//等到库存控制信息
func GetCtrInfo(ctx *gin.Context) {
	//Parameter parsing
	info := models.CtrStock{}
	err := ctx.ShouldBindJSON(&info)
	fmt.Println("00000000", info, ctx)
	if err != nil {
		fmt.Printf("[api.GetOutputInfo], Parameter parsing error: %v", err)
		return
	}

	ctr := new(services.CtrStock)
	ctrStock, err := ctr.GetCtrStockInfo(info.ProductNum)
	if err != nil {
		fmt.Printf("[api.UpdateProductLevel], err: %v", err)

		return
	}
	//err := ctx.ShouldBindJSON(&info)
	fmt.Println("00000000", info, ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"level":  ctrStock,
	})
}
