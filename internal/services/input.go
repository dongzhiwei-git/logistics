package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Input struct {
}

// 批量存数据库
func (in *Input) CreateInputInfo(allInput []models.Input) (err error) {
	for i := 0; i < len(allInput)-1; i++ {
		err = dao.Orm.Create(allInput[i+1]).Error
	}

	return err
}
