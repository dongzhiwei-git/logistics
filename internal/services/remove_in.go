package services

import "inherited/internal/dao"

func RemoveStoreId(storeId int)  {
	store,err:=QueryStoreById(storeId)
	if err!=nil{
		return
	}
	err=dao.Orm.Delete(&store).Error
	if err!=nil {
		return
	}
}
func RemoveStoreNumber(storeNumbers string)  {
	store,err:=QueryStoreByNumber(storeNumbers)
	if err!=nil{
		return
	}
	err=dao.Orm.Delete(&store).Error
	if err!=nil {
		return
	}
}
func RemoveStoreUserName(storeUserName string)  {
	store,err:=QueryStoreByNumber(storeUserName)
	if err!=nil{
		return
	}
	err=dao.Orm.Delete(&store).Error
	if err!=nil {
		return
	}
}
func RemoveStorePNeedSum(storeNeedSum int)  {
	store,err:=QueryStoreByNeedSum(storeNeedSum)
	if err!=nil{
		return
	}
	err=dao.Orm.Delete(&store).Error
	if err!=nil {
		return
	}
}
func RemoveStoreProductName(storeProductName string)  {
	store,err:=QueryStoreByNumber(storeProductName)
	if err!=nil{
		return
	}
	err=dao.Orm.Delete(&store).Error
	if err!=nil {
		return
	}
}
func RemoveStoreProductType(storeProductType string)  {
	store,err:=QueryStoreByProductType(storeProductType)
	if err!=nil{
		return
	}
	err=dao.Orm.Delete(&store).Error
	if err!=nil {
		return
	}
}