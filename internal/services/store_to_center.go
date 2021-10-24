package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Store struct {

}

func (s *Store) GetStoreInfo() (store *models.StoreInfo, err error) {
	storeInfo := new(models.StoreInfo)
	err = dao.Orm.Find(storeInfo).Error

	return storeInfo, err
}
