package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Output struct {
}

// 批量存数据库
func (out *Output) CreateOutputInfo(allOnput []models.Output) (err error) {
	for i := 0; i < len(allOnput)-1; i++ {
		err = dao.Orm.Create(allOnput[i+1]).Error
	}

	return err
}
