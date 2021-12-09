package models

type CtrStock struct {
	ProductNum    string `gorm:"product_num" json:"product_num"`
	OrderQuantity int    `gorm:"order_quantity" json:"order_quantity"`
	SafeStock     int    `gorm:"safe_stock" json:"safe_stock"`
	MaxStock      int    `gorm:"max_stock" json:"max_stock"`
}

func (CtrStock) TableName() string {
	return "forecast"
}
