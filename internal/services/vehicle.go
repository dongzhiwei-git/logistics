package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Vehicle struct {
}

// 批量存数据库
func (ve *Vehicle) CreateVehicleInfo(allVehicle []models.Vehicle) (err error) {
	for i := 0; i < len(allVehicle)-1; i++ {
		err = dao.Orm.Create(allVehicle[i+1]).Error
	}

	return err
}
