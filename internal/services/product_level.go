package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type ProductLevel struct {
}

// 批量存数据库
func (pl *ProductLevel) CreateLevelInfo(levelInfo []models.ProductLevel) (err error) {
	for i := 0; i < len(levelInfo)-1; i++ {
		err = dao.Orm.Create(levelInfo[i+1]).Error
	}

	return err
}
