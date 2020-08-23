package mysqlrepository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"github.com/jinzhu/gorm"
)

//UserFollowingTopicRepoImpl implement
type UserFollowingTopicRepoImpl struct {
	db *gorm.DB
}

//Insert insert one
func (instance *UserFollowingTopicRepoImpl) Insert(data *entities.UserFollowingTopic) (tx *gorm.DB, err error) {
	panic("implement me")
}

//Inserts insert multiple
func (instance *UserFollowingTopicRepoImpl) Inserts(data []*entities.UserFollowingTopic) (tx *gorm.DB, err error) {
	tx = instance.db.Begin()
	for i := range data {
		err = tx.Create(&data[i]).Error
		if err != nil {
			break
		}
	}
	return
}

//NewInstanceMysqlUserFollowingTopicRepoImpl new instance of UserFollowingTopicRepoImpl
func NewInstanceMysqlUserFollowingTopicRepoImpl(db *gorm.DB) repository.UserFollowingTopicRepo {
	return &UserFollowingTopicRepoImpl{db: db}
}
