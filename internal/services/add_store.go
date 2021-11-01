package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

func AddStore(newStore models.StoreInfo)  {
	if err:=dao.Orm.Create(newStore);err!=nil{
		return
	}
}
