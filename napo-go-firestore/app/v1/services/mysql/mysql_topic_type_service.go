package mysql

import (
	"app/app/v1/entities"
	"app/app/v1/injection"
	"app/app/v1/services"
)

type TopicTypeServiceImpl struct {
	Repository *injection.RepositoryInjection
}

func (instance *TopicTypeServiceImpl) GetOneByLabel(label string) (result entities.TopicType, err error) {
	//err = instance.Db.Where("label=?", label).First(&result).Error
	result, err = instance.Repository.MysqlTopicTypeRepo.GetByLabel(label)
	return
}

func (instance *TopicTypeServiceImpl) Inserts(data []*entities.TopicType) (err error) {
	//instance.Repository.MysqlTopicTypeRepo.Inserts(data)

	return
}

func (instance *TopicTypeServiceImpl) GetAll() (result []*entities.TopicType, err error) {
	result, err = instance.Repository.MysqlTopicTypeRepo.GetAll()
	return
}

func (instance *TopicTypeServiceImpl) Insert(data *entities.TopicType) (err error) {
	//err = instance.Db.Create(&data).Error
	return
}

func NewInstanceMysqlTopicTypeServices(repository *injection.RepositoryInjection) services.TopicTypeServices {
	return &TopicTypeServiceImpl{Repository: repository}
}
