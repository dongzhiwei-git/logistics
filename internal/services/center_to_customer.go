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
