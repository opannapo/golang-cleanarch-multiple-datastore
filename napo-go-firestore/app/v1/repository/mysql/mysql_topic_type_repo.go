package mysqlrepository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"github.com/jinzhu/gorm"
)

//TopicTypeRepoImpl implement
type TopicTypeRepoImpl struct {
	Db *gorm.DB
}

//GetAll return all
func (instance *TopicTypeRepoImpl) GetAll() (result []*entities.TopicType, err error) {
	err = instance.Db.Find(&result).Error
	return
}

//GetByLabel return filter by label
func (instance *TopicTypeRepoImpl) GetByLabel(label string) (result entities.TopicType, err error) {
	err = instance.Db.Where("label=?", label).First(&result).Error
	return
}

//Insert insert one
func (instance *TopicTypeRepoImpl) Insert(data *entities.TopicType) (err error) {
	err = instance.Db.Create(&data).Error
	return
}

//Inserts insert multiple
func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (tx *gorm.DB, err error) {
	tx = instance.Db.Begin()
	for i := range data {
		err = tx.Create(&data[i]).Error
		if err != nil {
			break
		}
	}

	return
}

//NewInstanceMysqlTopicTypeRepoImpl new instance of TopicTypeRepoImpl
func NewInstanceMysqlTopicTypeRepoImpl(db *gorm.DB) repository.TopicTypeRepo {
	return &TopicTypeRepoImpl{Db: db}
}
