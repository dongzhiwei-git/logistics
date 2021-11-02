package models

type Vehicle struct {
	Id               int    `gorm:"id" json:"id" form:"id"`
	VehicleNum       string `gorm:"vehicle_num" json:"vehicle_num" form:"vehicle_num"`
	ViSeverCondition string `gorm:"vi_sever_condition" json:"vi_sever_condition" form:"vi_sever_condition"`
	ViPosition       string `gorm:"vi_position" json:"vi_position" form:"vi_position"`
	ViLoadWeight     string `gorm:"vi_load_weight" json:"vi_load_weight" form:"vi_load_weight"`
	ViTransitValue   string `gorm:"vi_transit_value" json:"vi_transit_value" form:"vi_transit_value"`
}

func (Vehicle) TableName() string {
	return "vehicle"
}
