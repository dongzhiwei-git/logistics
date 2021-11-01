package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type Input struct {
}

// 批量存数据库
func (in *Input) CreateInputInfo(allInput []models.Input) (err error) {
	for i := 0; i < len(allInput)-1; i++ {
		err = dao.Orm.Create(allInput[i+1]).Error
	}

	return err
}

// 查入库info
func (s *Input) GetInputInfo() (store *[]models.Input, err error) {
	inputInfo := new([]models.Input)
	err = dao.Orm.Find(inputInfo).Error
	//info, err := json.Marshal(storeInfo)
	if err != nil {
		log.Println("[services.GetStoreInfo], err")
	}

	return inputInfo, err
}