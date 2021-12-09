package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/models"
	"inherited/internal/services"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Init() {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

}

//func main() {
//	Init()
//	//storeSql()
//	//center()
//	//input()
//	//output()
//	//vehicle()
//	productLevel()
//
//}

func storeSql() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		fmt.Println("er")
		return
	}
	data := make([]models.StoreInfo, 144)

	rows, err := f.GetRows("总仓到配送中心")
	for i, row := range rows {
		fmt.Println("i", i)
		data[i].Id = i + 1
		data[i].Number = row[0]
		data[i].Area = row[1]
		data[i].ProductName = row[2]
		NeedNum, _ := strconv.Atoi(row[3])
		data[i].NeedSum = NeedNum
		data[i].ProductType = row[4]
	}

	println()
	store := new(services.Store)
	err = store.CreateStoreInfo(data)
}

// 总仓到分仓
func center() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		fmt.Println("er")
		return
	}
	data := make([]models.CenterInfo, 4)
	rows, err := f.GetRows("配送中心到用户")
	for i, row := range rows {
		fmt.Println("i", i)
		data[i].Id = i + 1
		data[i].Number = row[0]
		data[i].UserName = row[1]
		data[i].ProductName = row[2]
		NeedNum, _ := strconv.Atoi(row[3])
		data[i].NeedSum = NeedNum
		data[i].ProductType = row[4]
	}
	println()

	center := new(services.Center)
	err = center.CreateCenterInfo(data)
	if err != nil {
		fmt.Println(err, 123)
	}
	fmt.Println(data)
}

// 入库表
func input() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		fmt.Println("er")
		return
	}
	data := make([]models.Input, 4)
	rows, err := f.GetRows("采购页面")
	for i, row := range rows {
		fmt.Println("i", i)
		data[i].Id = i + 1
		data[i].Area = row[0]
		data[i].ProductName = row[1]
		NeedNum, _ := strconv.Atoi(row[2])
		data[i].OrderSum = NeedNum
		LeftSum, _ := strconv.Atoi(row[3])
		data[i].LeftSum = LeftSum
		data[i].PurchaseLevel = row[4]
	}
	println()

	center := new(services.Input)
	err = center.CreateInputInfo(data)
	if err != nil {
		fmt.Println(err, 123)
	}
	fmt.Println(data)
}

// 出库表
func output() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		fmt.Println("er")
		return
	}
	data := make([]models.Output, 4)
	rows, err := f.GetRows("物流端")
	for i, row := range rows {
		fmt.Println("i", i)
		data[i].Id = i
		data[i].ProductName = row[0]
		data[i].OrderNum = row[1]
		orderSum, _ := strconv.Atoi(row[2])
		data[i].OrderSum = orderSum
		data[i].OrderCustomer = row[3]
		data[i].OrderCustomerID = row[4]
		data[i].ToArea = row[5]
		data[i].ArriveTime = row[6]
		data[i].TradeStatus = row[7]
	}
	println()

	output := new(services.Output)
	err = output.CreateOutputInfo(data)
	if err != nil {
		fmt.Println(err, 123)
	}
	fmt.Println(data)
}

// 车辆表
func vehicle() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		fmt.Println("er")
		return
	}
	data := make([]models.Vehicle, 4)
	rows, err := f.GetRows("车辆情况")
	for i, row := range rows {
		fmt.Println("i", i)
		data[i].Id = i
		data[i].VehicleNum = row[0]
		data[i].ViSeverCondition = row[1]

		data[i].ViPosition = row[2]
		data[i].ViLoadWeight = row[3]
		data[i].ViTransitValue = row[4]
	}
	println()

	vehicle := new(services.Vehicle)
	err = vehicle.CreateVehicleInfo(data)
	if err != nil {
		fmt.Println(err, 123)
	}
	fmt.Println(data)
}

func productLevel() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		fmt.Println("er")
		return
	}
	data := make([]models.ProductLevel, 166)
	rows, err := f.GetRows("备件等级")
	for i, row := range rows {
		fmt.Println("i", i)
		data[i].Id = i
		data[i].ProductName = row[0]
		data[i].TotalAmount, _ = strconv.ParseInt(row[1], 10, 64)
		data[i].General, _ = strconv.ParseInt(row[2], 10, 64)
		data[i].ForwardDate, _ = strconv.ParseInt(row[3], 10, 64)
		data[i].ProLevel = row[4]
	}
	println()

	level := new(services.ProductLevel)
	err = level.CreateLevelInfo(data)
	if err != nil {
		fmt.Println(err, 123)
	}
	fmt.Println(data)
}
