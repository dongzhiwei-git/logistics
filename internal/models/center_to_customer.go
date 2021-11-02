package models

type CenterInfo struct {
	Id          int    `json:"id" gorm:"id"`
	Number      string `json:"number" gorm:"number"`
	UserName    string `json:"user_name" gorm:"user_name"`
	ProductName string `json:"product_name" gorm:"product_name"`
	NeedSum     int    `json:"need_sum" gorm:"need_sum"`
	ProductType string `json:"product_type" gorm:"product_type"`
}

func (CenterInfo) TableName() string {
	return "center_to_customer"
}
