package mysql

import (
	"app/app/v1/entities"
	"app/app/v1/services"
	"github.com/jinzhu/gorm"
)

type TopicTypeServiceImpl struct {
	Db *gorm.DB
}

func (instance *TopicTypeServiceImpl) GetOneByLabel(label string) (result entities.TopicType, err error) {
	err = instance.Db.Where("label=?", label).First(&result).Error
	return
}

func (instance *TopicTypeServiceImpl) Inserts(data []*entities.TopicType) (err error, tx *gorm.DB) {
	tx = instance.Db.Begin()
	for i := range data {
		err = tx.Create(&data[i]).Error
		if err != nil {
			break
		}
	}

	return
}

func (instance *TopicTypeServiceImpl) GetAll() (result []*entities.TopicType, err error) {
	err = instance.Db.Find(&result).Error
	return
}

func (instance *TopicTypeServiceImpl) Insert(data *entities.TopicType) (err error) {
	err = instance.Db.Create(&data).Error
	return
}

func NewInstanceMysqlTopicTypeServices(db *gorm.DB) services.TopicTypeServices {
	return &TopicTypeServiceImpl{Db: db}
}
