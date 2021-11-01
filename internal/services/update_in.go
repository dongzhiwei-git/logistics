package services

import "inherited/internal/dao"

func UpdateStoreId(storeId int,value string)  {
	store,err:=QueryStoreById(storeId)
	if err!=nil{
		return
	}
	err=dao.Orm.Model(&store).Update("Id",value).Error
	if err!=nil {
		return
	}
}
func UpdateStoreNumber(storeNumbers string,value string)  {
	store,err:=QueryStoreByNumber(storeNumbers)
	if err!=nil{
		return
	}
	err=dao.Orm.Model(&store).Update("number",value).Error
	if err!=nil {
		return
	}
}
func UpdateStoreUserName(storeUserName string,value string)  {
	store,err:=QueryStoreByNumber(storeUserName)
	if err!=nil{
		return
	}
	err=dao.Orm.Model(&store).Update("user_name",value).Error
	if err!=nil {
		return
	}
}
func UpdateStorePNeedSum(storeNeedSum int,value int)  {
	store,err:=QueryStoreByNeedSum(storeNeedSum)
	if err!=nil{
		return
	}
	err=dao.Orm.Model(&store).Update("need_sum",value).Error
	if err!=nil {
		return
	}
}
func UpdateStoreProductName(storeProductName string,value string)  {
	store,err:=QueryStoreByNumber(storeProductName)
	if err!=nil{
		return
	}
	err=dao.Orm.Model(&store).Update("product_name",value).Error
	if err!=nil {
		return
	}
}
func UpdateStoreProductType(storeProductType string,value string)  {
	store,err:=QueryStoreByProductType(storeProductType)
	if err!=nil{
		return
	}
	err=dao.Orm.Model(&store).Update("product_type",value).Error
	if err!=nil {
		return
	}
}