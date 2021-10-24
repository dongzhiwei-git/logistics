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

	sysuser := new(services.SysUser)
	sysuserInfo, err := sysuser.GetSysUserInfo()
	if err != nil {
		fmt.Printf("[api.GetStoreInfo], err: %v", err)

		return
	}

	fmt.Println(sysuserInfo)
	//store := new(services.Store)
	//storeInfo, err := store.GetStoreInfo()
	//if err != nil {
	//	fmt.Printf("[api.GetStoreInfo], err: %v", err)
	//
	//	return
	//}
	//
	//fmt.Println(storeInfo)
}
