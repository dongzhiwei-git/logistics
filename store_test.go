package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/models"
	"inherited/internal/services"
	"log"
	"testing"
)

func TestQuery(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	var store models.StoreInfo
	var err error
	store,err=services.QueryStoreByUserName("客户2")
	fmt.Println(err,store)
}
func TestAdd(t *testing.T)  {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	var store =models.StoreInfo{
		Id:          1,
		Number:      "测试",
		UserName:    "测试",
		ProductName: "测试",
		NeedSum:     1,
		ProductType: "测试",
	}
	services.AddStore(store)
}
func TestDelete(t *testing.T)  {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	services.RemoveStoreNumber("test_numbers")
}
func TestUpdate(t *testing.T)  {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	services.UpdateStorePNeedSum(300,30)
}
//[{2 ZL2021102701 客户1 S管阀（260） 30 A2}
//{3 ZL2021102702 客户2 长城L-CKD150 齿轮油 18L\16KG 50 C1}
//{4 ZL2021102703 客户3 眼镜板Φ230 230 B1}]

