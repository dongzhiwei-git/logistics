package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type Store struct {
}

func (s *Store) GetStoreInfo() (store *models.StoreInfo, err error) {
	storeInfo := new(models.StoreInfo)
	err = dao.Orm.Find(storeInfo).Error
	//info, err := json.Marshal(storeInfo)
	if err != nil {
		log.Println("[services.GetStoreInfo], err")
	}

	return storeInfo, err
}

// 批量存数据库
func (s *Store) CreateStoreInfo(allStore []models.StoreInfo) (err error) {
	for i := 0; i < len(allStore) - 1; i++ {
		err = dao.Orm.Create(allStore[i+1]).Error
	}


	return err
}
