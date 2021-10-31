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
func main() {
	Init()
	storeSql()

}

func storeSql() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		fmt.Println("er")
		return
	}
	data := make([]models.StoreInfo, 4)

	rows, err := f.GetRows("总仓到配送中心")
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
	rows, err := f.GetRows("中心到用户")
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

	center := new(services.Center)
	err = center.CreateCenterInfo(data)
	if err != nil {
		fmt.Println(err, 123)
	}
	fmt.Println(data)
}
