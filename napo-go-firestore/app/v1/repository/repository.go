package repository

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	GetAll() (result []*entities.User, err error)
	GetById(id int) (result entities.User, err error)
	GetByCredential(username string, password string) (result *entities.User, err error)
	Insert(param *param.UserCreate) (err error, tx *gorm.DB)
	Update(user *entities.User) (err error)
	Delete(user *entities.User) (err error)
}

type TopicTypeRepo interface {
	GetAll() (result []*entities.TopicType, err error)
	GetByLabel(label string) (result entities.TopicType, err error)
	Insert(data *entities.TopicType) (err error)
	Inserts(data []*entities.TopicType) (err error, tx *gorm.DB)
}

type UserFollowingTopicRepo interface {
	Insert(data *entities.UserFollowingTopic) (err error, tx *gorm.DB)
	Inserts(data []*entities.UserFollowingTopic) (err error, tx *gorm.DB)
}

type CredentialRepo interface {
	GetByKeySignature(key string, signature string) (result entities.Credential, err error)
	Insert(param *param.UserCreate) (err error, tx *gorm.DB)
}
