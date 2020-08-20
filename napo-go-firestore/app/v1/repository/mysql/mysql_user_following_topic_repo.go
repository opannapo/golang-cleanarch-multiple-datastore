package mysql

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"github.com/jinzhu/gorm"
)

type UserFollowingTopicRepoImpl struct {
	db *gorm.DB
}

func (instance *UserFollowingTopicRepoImpl) Insert(data *entities.UserFollowingTopic) (err error, tx *gorm.DB) {
	panic("implement me")
}

func (instance *UserFollowingTopicRepoImpl) Inserts(data []*entities.UserFollowingTopic) (err error, tx *gorm.DB) {
	tx = instance.db.Begin()
	for i := range data {
		err = tx.Create(&data[i]).Error
		if err != nil {
			break
		}
	}
	return
}

func NewInstanceMysqlUserFollowingTopicRepoImpl(db *gorm.DB) repository.UserFollowingTopicRepo {
	return &UserFollowingTopicRepoImpl{db: db}
}
