package mysql_repository

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/repository"
	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func (instance *UserRepoImpl) GetAll() (result []*entities.User, err error) {
	err = instance.Db.
		Preload("FollowingTopic").
		Preload("FollowingTopic.TopicType").
		Find(&result).Error
	return
}

func (instance *UserRepoImpl) GetById(id int) (result entities.User, err error) {
	err = instance.Db.
		Preload("FollowingTopic").
		Preload("FollowingTopic.TopicType").
		Where("id=?", id).First(&result).Error
	return
}

func (instance *UserRepoImpl) GetByCredential(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

func (instance *UserRepoImpl) Insert(param *param.UserCreate) (err error, tx *gorm.DB) {
	tx = instance.Db.Begin()

	user := param.User
	err = tx.Create(&user).Error
	return
}

func (instance *UserRepoImpl) Update(user *entities.User) (err error) {
	panic("implement me")
}

func (instance *UserRepoImpl) Delete(user *entities.User) (err error) {
	panic("implement me")
}

func NewInstanceMysqlUserRepoImpl(db *gorm.DB) repository.UserRepo {
	return &UserRepoImpl{Db: db}
}
