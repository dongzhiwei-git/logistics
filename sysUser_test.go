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

func TestGetSysInfo(t *testing.T) {
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

func TestGetProductLevel(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	level := new(services.ProductLevel)
	levelInfo, err := level.GetProductLevel()
	if err != nil {
		fmt.Printf("[api.GetStoreInfo], err: %v", err)

		return
	}

	fmt.Println(levelInfo)
}

func TestUpdateProductLevel(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	level := new(services.ProductLevel)
	levelInfo, err := level.UpdateProductLevel(1, 2151380, 1, 72)
	if err != nil {
		fmt.Printf("[api.GetStoreInfo], err: %v", err)

		return
	}

	fmt.Println(levelInfo)
}
