package services

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
)

type UserServices interface {
	GetUsers() (result []*entities.User, err error)
	GetUser(id int) (result entities.User, err error)
	GetUserForAuth(username string, password string) (result *entities.User, err error)
	AddUser(user *param.UserCreate) (err error)
	UpdateUser(user *entities.User) (err error)
	DeleteUser(user *entities.User) (err error)
}

type TopicTypeServices interface {
	GetAll() (result []*entities.TopicType, err error)
	GetOneByLabel(label string) (result entities.TopicType, err error)
	Insert(data *entities.TopicType) (err error)
	Inserts(data []*entities.TopicType) (err error)
}

type UserFollowingTopicServices interface {
	Insert(data *entities.UserFollowingTopic) (err error)
	Inserts(data []*entities.UserFollowingTopic) (err error)
}

type AuthServices interface {
	InsertCredential() (err error)
}
