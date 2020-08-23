package mysqlrepository

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/repository"
	"github.com/jinzhu/gorm"
)

//UserRepoImpl implement
type UserRepoImpl struct {
	Db *gorm.DB
}

//GetAll get all users
func (instance *UserRepoImpl) GetAll() (result []*entities.User, err error) {
	err = instance.Db.
		Preload("FollowingTopic").
		Preload("FollowingTopic.TopicType").
		//Preload("Credential").
		Find(&result).Error
	return
}

//GetByID get user by ID
func (instance *UserRepoImpl) GetByID(id int) (result entities.User, err error) {
	err = instance.Db.
		Preload("FollowingTopic").
		Preload("FollowingTopic.TopicType").
		//Preload("Credential").
		Where("id=?", id).First(&result).Error
	return
}

//GetByCredential get user by credential
func (instance *UserRepoImpl) GetByCredential(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

//Insert insert one
func (instance *UserRepoImpl) Insert(param *param.UserCreate) (tx *gorm.DB, err error) {
	tx = instance.Db.Begin()

	user := param.User
	err = tx.Create(&user).Error
	return
}

//Update update one
func (instance *UserRepoImpl) Update(user *entities.User) (err error) {
	panic("implement me")
}

//Delete delete one
func (instance *UserRepoImpl) Delete(user *entities.User) (err error) {
	panic("implement me")
}

//NewInstanceMysqlUserRepoImpl new instance UserRepoImpl
func NewInstanceMysqlUserRepoImpl(db *gorm.DB) repository.UserRepo {
	return &UserRepoImpl{Db: db}
}
