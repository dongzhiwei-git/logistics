package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/services"
	"log"
	"testing"
)

func Add_Store(t testing.T)  {
	if err:=internal.Init();err!=nil{
		log.Println("Init failed")
		return
	}
	store:=new(services.Store)
	storeInfo, err := store.GetStoreInfo()
	if err!=nil{
		fmt.Println("[api.GetStoreInfo], err%v",err)
		return
	}

	fmt.Println(storeInfo)
}
