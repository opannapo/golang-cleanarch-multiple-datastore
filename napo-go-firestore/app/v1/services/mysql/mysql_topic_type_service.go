package mysqlservices

import (
	"app/app/v1/entities"
	"app/app/v1/injection/repositories"
	"app/app/v1/services"
)

//TopicTypeServiceImpl implement
type TopicTypeServiceImpl struct {
	Repository *repositories.RepositoryInjection
}

//GetOneByLabel get by label
func (instance *TopicTypeServiceImpl) GetOneByLabel(label string) (result entities.TopicType, err error) {
	//err = instance.Db.Where("label=?", label).First(&result).Error
	result, err = instance.Repository.MysqlTopicTypeRepo.GetByLabel(label)
	return
}

//Inserts insert multiple
func (instance *TopicTypeServiceImpl) Inserts(data []*entities.TopicType) (err error) {
	//instance.Repository.MysqlTopicTypeRepo.Inserts(data)
	return
}

//GetAll get all
func (instance *TopicTypeServiceImpl) GetAll() (result []*entities.TopicType, err error) {
	result, err = instance.Repository.MysqlTopicTypeRepo.GetAll()
	return
}

//Insert insert one
func (instance *TopicTypeServiceImpl) Insert(data *entities.TopicType) (err error) {
	//err = instance.Db.Create(&data).Error
	return
}

//NewInstanceMysqlTopicTypeServices new instance of TopicTypeServiceImpl
func NewInstanceMysqlTopicTypeServices(repository *repositories.RepositoryInjection) services.TopicTypeServices {
	return &TopicTypeServiceImpl{Repository: repository}
}
