package mysql

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/injection"
	"app/app/v1/services"
	"fmt"
)

type UserServiceImpl struct {
	Repository *injection.RepositoryInjection
}

func (instance *UserServiceImpl) GetUserForAuth(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) GetUsers() (result []*entities.User, err error) {
	/*err = instance.Db.
	Preload("FollowingTopic").
	Preload("FollowingTopic.TopicType").
	Find(&result).Error*/
	result, err = instance.Repository.MysqlUserRepo.GetAll()
	return
}

func (instance *UserServiceImpl) GetUser(id int) (result entities.User, err error) {
	/*err = instance.Db.
	Preload("FollowingTopic").
	Preload("FollowingTopic.TopicType").
	Where("id=?", id).First(&result).Error*/
	result, err = instance.Repository.MysqlUserRepo.GetById(id)
	return
}

func (instance *UserServiceImpl) AddUser(param *param.UserCreate) (err error) {
	/*tx = instance.Db.Begin()
	user := param.User
	err = tx.Create(&user).Error*/

	mysqlUserRepo := instance.Repository.MysqlUserRepo
	mysqlTopicRepo := instance.Repository.MysqlTopicTypeRepo
	mysqlUserFollowingTopicRepo := instance.Repository.MysqlUserFollowingTopicRepo
	firestoreUserRepo := instance.Repository.FirestoreUserRepo
	firestoreTopicTypeRepo := instance.Repository.FirestoreTopicTypeRepo

	//Check Topic Type master table
	var topicTypeTmpToCreate []*entities.TopicType
	var topicTypeTmpExist []*entities.TopicType
	for i := range param.FollowingTopic {
		label := *param.FollowingTopic[i]
		topicExist, err := mysqlTopicRepo.GetByLabel(label)
		if err != nil {
			tmp := entities.TopicType{
				Id:    0,
				Label: label}
			topicTypeTmpToCreate = append(topicTypeTmpToCreate, &tmp)
		} else {
			topicTypeTmpExist = append(topicTypeTmpExist, &topicExist)
		}
	}

	err, txInsertTopic := mysqlTopicRepo.Inserts(topicTypeTmpToCreate)
	if err != nil {
		txInsertTopic.Rollback()
		return
	} else {
		topicTypeTmpExist = append(topicTypeTmpExist, topicTypeTmpToCreate...)
	}

	//Insert User
	err, txInsertUser := mysqlUserRepo.Insert(param)
	if err != nil {
		txInsertTopic.Rollback()
		txInsertUser.Rollback()
		return
	}

	//Generate UserFollowingTopic & insert
	var tmpUserFollowingTopic []*entities.UserFollowingTopic
	for i := range topicTypeTmpExist {
		tmp := entities.UserFollowingTopic{
			UserId:      param.User.Id,
			TopicTypeId: topicTypeTmpExist[i].Id,
		}
		fmt.Println(tmp)
		tmpUserFollowingTopic = append(tmpUserFollowingTopic, &tmp)
		fmt.Println(tmpUserFollowingTopic)
	}
	err, txInsertUserFollowingTopic := mysqlUserFollowingTopicRepo.Inserts(tmpUserFollowingTopic)
	if err != nil {
		txInsertTopic.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		return
	}

	//Insert FirestoreServicesInjected User
	err, _ = firestoreUserRepo.Insert(param)
	if err != nil {
		txInsertTopic.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		return
	}

	err, _ = firestoreTopicTypeRepo.Inserts(topicTypeTmpExist)
	if err != nil {
		txInsertTopic.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		return
	}

	txInsertTopic.Commit()
	txInsertUser.Commit()
	txInsertUserFollowingTopic.Commit()
	return
}

func (instance *UserServiceImpl) UpdateUser(user *entities.User) (err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) DeleteUser(user *entities.User) (err error) {
	panic("implement me")
}

func NewInstanceMysqlUserServices(repository *injection.RepositoryInjection) services.UserServices {
	return &UserServiceImpl{Repository: repository}
}
