package mysqlservices

import (
	"app/app/v1/entities"
	"app/app/v1/injection/repositories"
	"app/app/v1/services"
)

//UserFollowingTopicService implement
type UserFollowingTopicService struct {
	Repository *repositories.RepositoryInjection
}

//Insert insert one
func (instance *UserFollowingTopicService) Insert(data *entities.UserFollowingTopic) (err error) {
	panic("implement me")
}

//Inserts insert multiple
func (instance *UserFollowingTopicService) Inserts(data []*entities.UserFollowingTopic) (err error) {
	/*tx = instance.db.Begin()
	for i := range data {
		err = tx.Create(&data[i]).Error
		if err != nil {
			break
		}
	}*/
	return
}

//NewInstanceMysqlUserFollowingTopicService new instance of
func NewInstanceMysqlUserFollowingTopicService(repository *repositories.RepositoryInjection) services.UserFollowingTopicServices {
	return &UserFollowingTopicService{Repository: repository}
}
