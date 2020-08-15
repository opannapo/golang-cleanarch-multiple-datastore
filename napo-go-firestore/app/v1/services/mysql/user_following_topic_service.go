package mysql

import (
	"app/app/v1/entities"
	"app/app/v1/services"
	"github.com/jinzhu/gorm"
)

type UserFollowingTopicService struct {
	db *gorm.DB
}

func (instance *UserFollowingTopicService) Insert(data *entities.UserFollowingTopic) (err error, tx *gorm.DB) {
	panic("implement me")
}

func (instance *UserFollowingTopicService) Inserts(data []*entities.UserFollowingTopic) (err error, tx *gorm.DB) {
	tx = instance.db.Begin()
	for i := range data {
		err = tx.Create(&data[i]).Error
		if err != nil {
			break
		}
	}
	return
}

func NewInstanceMysqlUserFollowingTopicService(db *gorm.DB) services.UserFollowingTopicServices {
	return &UserFollowingTopicService{db: db}
}
