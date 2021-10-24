package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Center struct {

}

func (cen *Center) GetSysUserInfo() (sysUser *models.SysUser, err error) {
	sysUserInfo := new(models.SysUser)
	err = dao.Orm.Find(sysUserInfo).Error

	return sysUserInfo, err
}
