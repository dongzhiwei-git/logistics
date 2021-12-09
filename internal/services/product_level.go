package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
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

func (pl *ProductLevel) GetProductLevel() (levelInfo *[]models.ProductLevel, err error) {
	levelInfo = new([]models.ProductLevel)
	err = dao.Orm.Find(levelInfo).Error
	if err != nil {
		log.Println("[services.GetProductLevel], ", err)
	}

	return levelInfo, err
}

func (pl *ProductLevel) UpdateProductLevel(id int, totalAmount, general, forwardDate int64) (level string,
	err error) {
	levelInfo := models.ProductLevel{}
	err = dao.Orm.Model(&levelInfo).Where("id=?", id).Update(models.ProductLevel{
		TotalAmount: totalAmount,
		General:     general,
		ForwardDate: forwardDate,
	}).Error
	if err != nil {
		return "", err
	}
	var levelProductInfo []models.ProductLevel
	err = dao.Orm.Find(&levelProductInfo).Error
	if err != nil {
		log.Println("Find error")
	}
	var (
		amountSum  int64
		generalSum int64
		forwardSum int64
	)
	for _, v := range levelProductInfo {
		amountSum += v.TotalAmount
		generalSum += v.General
		forwardSum += v.ForwardDate
	}
	amount := float64(totalAmount) / float64(amountSum)
	generalFloat := float64(general) / float64(generalSum)
	forward := float64(forwardDate) / float64(forwardSum)
	sum := 0.65*amount + 0.25*generalFloat + 0.1*forward
	log.Println("sum = ", sum)
	if sum > 0.03 {
		return "A1", nil
	} else if sum > 0.02 {
		return "A1", nil

	} else if sum > 0.01 {
		return "B", nil
	} else {
		return "C", nil
	}

}
