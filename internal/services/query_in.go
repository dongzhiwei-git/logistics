package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

func QueryStoreById( storeId int)  (storeInfo models.StoreInfo,err error){
	var store models.StoreInfo
	err = dao.Orm.Find(&store, "id=?", storeId).Error
	return store,err
}
func QueryStoreByNumber( storeNumber string)  (storeInfo models.StoreInfo,err error){
	var store models.StoreInfo
	err = dao.Orm.Find(&store, "number=?", storeNumber).Error
	return store,err
}
func QueryStoreByUserName( storeName string)  (storeInfo models.StoreInfo,err error){
	var store models.StoreInfo
	err = dao.Orm.Find(&store, "user_name=?", storeName).Error
	return store,err
}
func QueryStoreByProductName( storeId string)  (storeInfo models.StoreInfo,err error){
	var store models.StoreInfo
	err = dao.Orm.Find(&store, "product_name=?", storeId).Error
	return store,err
}
func QueryStoreByNeedSum( storeNeedSum int)  (storeInfo models.StoreInfo,err error){
	var store models.StoreInfo
	err = dao.Orm.Find(&store, "need_sum=?", storeNeedSum).Error
	return store,err
}
func QueryStoreByProductType( productType string)  (storeInfo models.StoreInfo,err error){
	var store models.StoreInfo
	err = dao.Orm.Find(&store, "product_type=?", productType).Error
	return store,err
}
