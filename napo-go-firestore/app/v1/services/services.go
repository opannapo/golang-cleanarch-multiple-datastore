package services

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
)

//UserServices interface layer
type UserServices interface {
	GetUsers() (result []*entities.User, err error)
	GetUser(id int) (result entities.User, err error)
	GetUserForAuth(username string, password string) (result *entities.User, err error)
	AddUser(user *param.UserCreate) (err error)
	UpdateUser(user *entities.User) (err error)
	DeleteUser(user *entities.User) (err error)
}

//TopicTypeServices interface layer
type TopicTypeServices interface {
	GetAll() (result []*entities.TopicType, err error)
	GetOneByLabel(label string) (result entities.TopicType, err error)
	Insert(data *entities.TopicType) (err error)
	Inserts(data []*entities.TopicType) (err error)
}

//UserFollowingTopicServices interface layer
type UserFollowingTopicServices interface {
	Insert(data *entities.UserFollowingTopic) (err error)
	Inserts(data []*entities.UserFollowingTopic) (err error)
}

//AuthServices interface layer
type AuthServices interface {
	ValidateCredential(param *param.AuthParam) (result entities.Credential, token string, err error)
}
