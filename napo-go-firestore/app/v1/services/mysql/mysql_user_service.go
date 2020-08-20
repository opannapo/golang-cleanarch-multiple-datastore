package mysql

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/injection"
	"app/app/v1/services"
)

type UserServiceImpl struct {
	Repository *injection.RepositoryInjection
}

func (instance *UserServiceImpl) GetUserForAuth(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) GetUsers() (result []*entities.User, err error) {
	/*err = instance.Db.
	Preload("FollowingTopic").
	Preload("FollowingTopic.TopicType").
	Find(&result).Error*/
	result, err = instance.Repository.MysqlUserRepo.GetAll()
	return
}

func (instance *UserServiceImpl) GetUser(id int) (result entities.User, err error) {
	/*err = instance.Db.
	Preload("FollowingTopic").
	Preload("FollowingTopic.TopicType").
	Where("id=?", id).First(&result).Error*/
	result, err = instance.Repository.MysqlUserRepo.GetById(id)
	return
}

func (instance *UserServiceImpl) AddUser(param *param.UserCreate) (err error) {
	/*tx = instance.Db.Begin()

	user := param.User
	err = tx.Create(&user).Error*/
	return
}

func (instance *UserServiceImpl) UpdateUser(user *entities.User) (err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) DeleteUser(user *entities.User) (err error) {
	panic("implement me")
}

func NewInstanceMysqlUserServices(repository *injection.RepositoryInjection) services.UserServices {
	return &UserServiceImpl{Repository: repository}
}
