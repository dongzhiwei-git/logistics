package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type Output struct {
}

// 批量存数据库
func (out *Output) CreateOutputInfo(allOnput []models.Output) (err error) {
	for i := 0; i < len(allOnput)-1; i++ {
		err = dao.Orm.Create(allOnput[i+1]).Error
	}

	return err
}

// 查出库info
func (out *Output) GetOutputInfo() (store *[]models.Output, err error) {
	outputInfo := new([]models.Output)
	err = dao.Orm.Find(outputInfo).Error
	//info, err := json.Marshal(storeInfo)
	if err != nil {
		log.Println("[services.GetOutputInfo], err")
	}

	return outputInfo, err
}
