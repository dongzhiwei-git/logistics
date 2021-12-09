package models

type ProductLevel struct {
	Id          int    `gorm:"id" json:"id" binding:"required"`
	ProductName string `gorm:"product_name" json:"product_name" `
	TotalAmount int64  `gorm:"total_amount" json:"total_amount" binding:"required"`
	General     int64  `gorm:"general" json:"general" binding:"required"`
	ForwardDate int64  `gorm:"forward_date"  json:"forward_date" binding:"required"`
	ProLevel    string `gorm:"pro_level" json:"pro_level" `
}

func (ProductLevel) TableName() string {
	return "ctr_level"
}
