package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type CtrStock struct {
}

func (c *CtrStock) GetCtrStockInfo(pro string) (ctr *models.CtrStock, err error) {
	info := new(models.CtrStock)
	err = dao.Orm.Where("product_num = ?", pro).Find(info).Error
	//info, err := json.Marshal(storeInfo)
	if err != nil {
		log.Println("[services.GetStoreInfo], err")
	}

	return info, err
}
