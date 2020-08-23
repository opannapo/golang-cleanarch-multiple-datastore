package repository

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"github.com/jinzhu/gorm"
)

//UserRepo interface layer
type UserRepo interface {
	GetAll() (result []*entities.User, err error)
	GetByID(id int) (result entities.User, err error)
	GetByCredential(username string, password string) (result *entities.User, err error)
	Insert(param *param.UserCreate) (tx *gorm.DB, err error)
	Update(user *entities.User) (err error)
	Delete(user *entities.User) (err error)
}

//TopicTypeRepo interface layer
type TopicTypeRepo interface {
	GetAll() (result []*entities.TopicType, err error)
	GetByLabel(label string) (result entities.TopicType, err error)
	Insert(data *entities.TopicType) (err error)
	Inserts(data []*entities.TopicType) (tx *gorm.DB, err error)
}

//UserFollowingTopicRepo interface layer
type UserFollowingTopicRepo interface {
	Insert(data *entities.UserFollowingTopic) (tx *gorm.DB, err error)
	Inserts(data []*entities.UserFollowingTopic) (tx *gorm.DB, err error)
}

//CredentialRepo interface layer
type CredentialRepo interface {
	GetByKeySignature(key string, signature string) (result entities.Credential, err error)
	Insert(data *entities.Credential) (tx *gorm.DB, err error)
}
