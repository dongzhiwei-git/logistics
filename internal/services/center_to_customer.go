package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Center struct {
}

func (cen *Center) GetCenterInfo() (sysUser *models.CenterInfo, err error) {
	centerInfo := new(models.CenterInfo)
	err = dao.Orm.Find(centerInfo).Error

	return centerInfo, err
}

// 批量存数据库
func (cen *Center) CreateCenterInfo(allCenter []models.CenterInfo) (err error) {
	for i := 0; i < len(allCenter)-1; i++ {
		err = dao.Orm.Create(allCenter[i+1]).Error
	}

	return err
}
