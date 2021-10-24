package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/services"
	"log"
	"testing"
)

func TestGetCenterInfo(t *testing.T){
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	center := new(services.Center)
	centerInfo, err := center.GetCenterInfo()
	if err != nil {
		fmt.Printf("[api.GetStoreInfo], err: %v", err)

		return
	}

	fmt.Println(centerInfo)
}
