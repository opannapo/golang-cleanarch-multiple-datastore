package firestore_services

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/injection/repository"
	"app/app/v1/services"
)

type UserServiceImpl struct {
	Repository *repository.RepositoryInjection
}

func (instance *UserServiceImpl) GetUserForAuth(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) GetUsers() (result []*entities.User, err error) {
	result, err = instance.Repository.FirestoreUserRepo.GetAll()
	return
}

func (instance *UserServiceImpl) GetUser(id int) (result entities.User, err error) {
	panic("Implement me")
}

func (instance *UserServiceImpl) AddUser(param *param.UserCreate) (err error) {
	err, _ = instance.Repository.FirestoreUserRepo.Insert(param)
	return
}

func (instance *UserServiceImpl) UpdateUser(user *entities.User) (err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) DeleteUser(user *entities.User) (err error) {
	panic("implement me")
}

func NewInstanceFirestoreUserService(repository *repository.RepositoryInjection) services.UserServices {
	return &UserServiceImpl{Repository: repository}
}
