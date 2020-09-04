package services

import (
	"app/app/v1/entities"
	"app/app/v1/injection/repositories"
)

//TopicTypeServiceImpl implement
type TopicTypeServiceImpl struct {
	Repository *repoinjection.RepositoryInjection
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
	//result, err = instance.Repository.MysqlTopicTypeRepo.GetAll()

	result, err = instance.Repository.RedisTopicTypeRepo.GetAll() //get from redis
	if err != nil {
		result, err = instance.Repository.MysqlTopicTypeRepo.GetAll()
		if len(result) > 0 {
			_, _ = instance.Repository.RedisTopicTypeRepo.Inserts(result) //set to redis
		}
	}
	return
}

//Insert insert one
func (instance *TopicTypeServiceImpl) Insert(data *entities.TopicType) (err error) {
	//err = instance.Db.Create(&data).Error
	return
}

//NewInstanceMysqlTopicTypeServices new instance of TopicTypeServiceImpl
func NewInstanceMysqlTopicTypeServices(repository *repoinjection.RepositoryInjection) TopicTypeServices {
	return &TopicTypeServiceImpl{Repository: repository}
}
