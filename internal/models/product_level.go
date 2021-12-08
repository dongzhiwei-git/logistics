package models

type ProductLevel struct {
	Id          int
	ProductName string `gorm:"product_name" json:"product_name"`
	TotalAmount int64  `gorm:"total_amount" json:"total_amount"`
	General     int64  `gorm:"general" json:"general"`
	ForwardDate int64  `gorm:"forward_date" json:"forward_date"`
}

func (ProductLevel) TableName() string {
	return "ctr_level"
}
