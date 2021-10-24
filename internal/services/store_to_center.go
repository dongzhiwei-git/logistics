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
	if err != nil{
		log.Println("[services.GetStoreInfo], err")
	}

	return storeInfo, err
}
