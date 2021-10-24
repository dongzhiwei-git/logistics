package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/services"
	"log"
	"testing"
)

func TestGetStoreInfo(t *testing.T){
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}


	store := new(services.Store)
	storeInfo, err := store.GetStoreInfo()
	if err != nil {
		fmt.Printf("[api.GetStoreInfo], err: %v", err)

		return
	}

	fmt.Println(storeInfo)
}
