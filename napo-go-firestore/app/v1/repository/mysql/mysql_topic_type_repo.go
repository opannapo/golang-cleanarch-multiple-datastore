package mysqlrepository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"fmt"
	"github.com/jinzhu/gorm"
)

//TopicTypeRepoImpl implement
type TopicTypeRepoImpl struct {
	Db *gorm.DB
}

//Upserts ignore on this
func (instance *TopicTypeRepoImpl) Upserts(data []*entities.TopicType) (tx *repository.DbTransactionType, err error) {
	panic("implement me")
}

//GetAll return all
func (instance *TopicTypeRepoImpl) GetAll() (result []*entities.TopicType, err error) {
	err = instance.Db.Find(&result).Error
	fmt.Println("Load from MySql")
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
func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (tx *repository.DbTransactionType, err error) {
	tx = &repository.DbTransactionType{
		GormTX:  instance.Db.Begin(),
		RedisTX: nil,
	}

	for i := range data {
		err = tx.GormTX.Create(&data[i]).Error
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
