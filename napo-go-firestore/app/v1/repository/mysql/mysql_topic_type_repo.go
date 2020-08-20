package mysql_repository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"github.com/jinzhu/gorm"
)

type TopicTypeRepoImpl struct {
	Db *gorm.DB
}

func (instance *TopicTypeRepoImpl) GetAll() (result []*entities.TopicType, err error) {
	err = instance.Db.Find(&result).Error
	return
}

func (instance *TopicTypeRepoImpl) GetByLabel(label string) (result entities.TopicType, err error) {
	err = instance.Db.Where("label=?", label).First(&result).Error
	return
}

func (instance *TopicTypeRepoImpl) Insert(data *entities.TopicType) (err error) {
	err = instance.Db.Create(&data).Error
	return
}

func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (err error, tx *gorm.DB) {
	tx = instance.Db.Begin()
	for i := range data {
		err = tx.Create(&data[i]).Error
		if err != nil {
			break
		}
	}

	return
}

func NewInstanceMysqlTopicTypeRepoImpl(db *gorm.DB) repository.TopicTypeRepo {
	return &TopicTypeRepoImpl{Db: db}
}
