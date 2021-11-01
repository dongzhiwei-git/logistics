package models

type Output struct {
	Id              int    `gorm:"id" json:"id" form:"id"`
	ProductName     string `gorm:"product_name" json:"product_name" from:"product_name"`
	OrderNum        string `gorm:"order_num" json:"order_num" form:"order_num"`
	OrderSum        int    `gorm:"order_sum" json:"order_sum" from:"order_sum"`
	OrderCustomer   string `gorm:"order_customer" json:"order_customer" form:"order_customer"`
	OrderCustomerID string `gorm:"order_customer_id" json:"order_customer_id" form:"order_customer_id"`
	ToArea          string `gorm:"to_area" json:"to_area" form:"to_area"`
	ArriveTime      string `gorm:"arrive_time" json:"arrive_time" form:"arrive_time"`
	TradeStatus     string `gorm:"trade_status" json:"trade_status" form:"trade_status"`
}

func (Output) TableName() string {
	return "out"
}
