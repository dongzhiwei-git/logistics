package models

type Input struct {
	Id            int    `gorm:"id" json:"id" form:"id"`
	Area          string `gorm:"area" json:"area" from:"area"`
	ProductName   string `gorm:"product_name" json:"product_name" from:"product_name"`
	OrderSum      int    `gorm:"order_sum" json:"order_sum" from:"order_sum"`
	LeftSum       int    `gorm:"left_sum" json:"left_sum" from:"left_sum"`
	PurchaseLevel int    `gorm:"purchase_level" json:"purchase_level" from:"purchase_level"`
}

func (Input) TableName() string {
	return "purchase"
}
