package firestoreservices

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/injection/repositories"
	"app/app/v1/services"
)

//UserServiceImpl implement
type UserServiceImpl struct {
	Repository *repositories.RepositoryInjection
}

//GetUserForAuth auth user by username pass
func (instance *UserServiceImpl) GetUserForAuth(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

//GetUsers return all users
func (instance *UserServiceImpl) GetUsers() (result []*entities.User, err error) {
	result, err = instance.Repository.FirestoreUserRepo.GetAll()
	return
}

//GetUser return user
func (instance *UserServiceImpl) GetUser(id int) (result entities.User, err error) {
	panic("Implement me")
}

//AddUser add user to firestore
func (instance *UserServiceImpl) AddUser(param *param.UserCreate) (err error) {
	_, err = instance.Repository.FirestoreUserRepo.Insert(param)
	return
}

//UpdateUser update one
func (instance *UserServiceImpl) UpdateUser(user *entities.User) (err error) {
	panic("implement me")
}

//DeleteUser delete one
func (instance *UserServiceImpl) DeleteUser(user *entities.User) (err error) {
	panic("implement me")
}

//NewInstanceFirestoreUserService new instance of UserServiceImpl
func NewInstanceFirestoreUserService(repository *repositories.RepositoryInjection) services.UserServices {
	return &UserServiceImpl{Repository: repository}
}
