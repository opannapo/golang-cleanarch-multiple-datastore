package mysql

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/services"
	"github.com/jinzhu/gorm"
)

type UserServiceImpl struct {
	Db *gorm.DB
}

func (instance *UserServiceImpl) GetUsers() (result []*entities.User, err error) {
	err = instance.Db.
		Preload("FollowingTopic").
		Preload("FollowingTopic.TopicType").
		Find(&result).Error
	return
}

func (instance *UserServiceImpl) GetUser(id int) (result entities.User, err error) {
	err = instance.Db.
		Preload("FollowingTopic").
		Preload("FollowingTopic.TopicType").
		Where("id=?", id).First(&result).Error
	return
}

func (instance *UserServiceImpl) AddUser(param *param.UserCreate) (err error, tx *gorm.DB) {
	tx = instance.Db.Begin()

	user := param.User
	err = tx.Create(&user).Error
	return
}

func (instance *UserServiceImpl) UpdateUser(user *entities.User) (err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) DeleteUser(user *entities.User) (err error) {
	panic("implement me")
}

func NewInstanceMysqlUserServices(db *gorm.DB) services.UserServices {
	return &UserServiceImpl{Db: db}
}
