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

func main() {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		fmt.Println("er")
		return
	}
	//data := make([]models.StoreInfo, 4)
	// 获取工作表中指定单元格的值
	//cell, err := f.GetCellValue("txt001", "B2")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//println(cell)

	//获取 Sheet1 上所有单元格
	//rows, err := f.GetRows("中心到用户")
	//for i, row := range rows {
	//	fmt.Println("i", i)
	//	data[i].Id = i + 1
	//	data[i].Number = row[0]
	//	data[i].UserName = row[1]
	//	data[i].ProductName = row[2]
	//	NeedNum, _ := strconv.Atoi(row[3])
	//	data[i].NeedSum = NeedNum
	//	data[i].ProductType = row[4]
	//}
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
	//store := new(services.Store)
	//err = store.CreateStoreInfo(data)
	center := new(services.Center)
	err = center.CreateCenterInfo(data)
	if err != nil {
		fmt.Println(err, 123)
	}
	fmt.Println(data)
	//计时器
	//t1 := time.Now()
	//rows, _ := f.GetRows("txt001")
	//for err, _ := range rows {
	//	//fmt.Println(err)
	//	if err == 0 {
	//		fmt.Printf("2")
	//	}
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//
	//}
	//elapsed := time.Since(t1)
	//fmt.Println("App elapsed: ", elapsed)

	//并发读
	//t1 := time.Now()
	//wg := sync.WaitGroup{}
	//wg.Add(16010)
	//count := 0
	//rows, _ := f.GetRows("txt001")
	//for _, row := range rows {
	//	go func(i []string) {
	//		if i == nil {
	//			fmt.Printf("2")
	//		}
	//		wg.Done()
	//	}(row)
	//
	//}
	//wg.Wait()
	//elapsed := time.Since(t1)
	//fmt.Println("App elapsed: ", elapsed, "count=", count)

	//ch := make(chan int, 10)

	//ch <- 1
	//<-ch
	//fmt.Println(1)

}
