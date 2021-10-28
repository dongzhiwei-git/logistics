package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/services"
	"log"
	"testing"
)

func TestCreateSysUser(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	SysUser := new(services.SysUser)
	SysUser.CreateSysUser("rot", "2355")
	fmt.Println("er")

}

func TestGetSysInfo(t *testing.T){
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	sysUser := new(services.SysUser)
	sysUserInfo, err := sysUser.GetSysUserInfo()
	if err != nil {
		fmt.Printf("[api.GetStoreInfo], err: %v", err)

		return
	}

	fmt.Println(sysUserInfo)
}
