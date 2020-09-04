package services

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/injection/repositories"
	"fmt"
)

//UserServiceImpl implement
type UserServiceImpl struct {
	Repository *repoinjection.RepositoryInjection
}

//GetUserForAuth get user for auth
func (instance *UserServiceImpl) GetUserForAuth(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

//GetUsers get all
func (instance *UserServiceImpl) GetUsers() (result []*entities.User, err error) {
	result, err = instance.Repository.MysqlUserRepo.GetAll()
	return
}

//GetUser get one
func (instance *UserServiceImpl) GetUser(id int) (result entities.User, err error) {
	result, err = instance.Repository.MysqlUserRepo.GetByID(id)
	return
}

//AddUser insert new user include all membership model
func (instance *UserServiceImpl) AddUser(param *param.UserCreate) (err error) {
	mysqlUserRepo := instance.Repository.MysqlUserRepo
	mysqlTopicRepo := instance.Repository.MysqlTopicTypeRepo
	mysqlUserFollowingTopicRepo := instance.Repository.MysqlUserFollowingTopicRepo
	mysqlCredentialRepo := instance.Repository.MysqlCredentialRepo
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
				ID:    0,
				Label: label}
			topicTypeTmpToCreate = append(topicTypeTmpToCreate, &tmp)
		} else {
			topicTypeTmpExist = append(topicTypeTmpExist, &topicExist)
		}
	}

	txInsertTopic, err := mysqlTopicRepo.Inserts(topicTypeTmpToCreate)
	if err == nil {
		topicTypeTmpExist = append(topicTypeTmpExist, topicTypeTmpToCreate...)
	} else {
		if txInsertTopic != nil {
			txInsertTopic.GormTX.Rollback()
		}
		return
	}

	//Insert User
	txInsertUser, err := mysqlUserRepo.Insert(param)
	if err != nil {
		txInsertTopic.GormTX.Rollback()
		if txInsertUser != nil {
			txInsertUser.Rollback()
		}
		return
	}

	//Generate UserFollowingTopic & insert
	var tmpUserFollowingTopic []*entities.UserFollowingTopic
	for i := range topicTypeTmpExist {
		tmp := entities.UserFollowingTopic{
			UserID:      param.User.ID,
			TopicTypeID: topicTypeTmpExist[i].ID,
		}
		fmt.Println(tmp)
		tmpUserFollowingTopic = append(tmpUserFollowingTopic, &tmp)
		fmt.Println(tmpUserFollowingTopic)
	}
	txInsertUserFollowingTopic, err := mysqlUserFollowingTopicRepo.Inserts(tmpUserFollowingTopic)
	if err != nil {
		txInsertTopic.GormTX.Rollback()
		txInsertUser.Rollback()
		if txInsertUserFollowingTopic != nil {
			txInsertUserFollowingTopic.Rollback()
		}
		return
	}

	//Insert User Credential
	param.Credential.UserID = param.User.ID
	txInsertCredential, err := mysqlCredentialRepo.Insert(param.Credential)
	if err != nil {
		txInsertTopic.GormTX.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		if txInsertCredential != nil {
			txInsertCredential.Rollback()
		}
		return
	}

	//Insert FirestoreServicesInjected User
	_, err = firestoreUserRepo.Insert(param)
	if err != nil {
		txInsertTopic.GormTX.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		txInsertCredential.Rollback()
		return
	}

	_, err = firestoreTopicTypeRepo.Inserts(topicTypeTmpExist)
	if err != nil {
		txInsertTopic.GormTX.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		txInsertCredential.Rollback()
		return
	}

	txInsertTopic.GormTX.Commit()
	txInsertUser.Commit()
	txInsertUserFollowingTopic.Commit()
	txInsertCredential.Commit()

	return
}

//UpdateUser update user
func (instance *UserServiceImpl) UpdateUser(user *entities.User) (err error) {
	panic("implement me")
}

//DeleteUser delete one
func (instance *UserServiceImpl) DeleteUser(user *entities.User) (err error) {
	panic("implement me")
}

//NewInstanceMysqlUserServices instance of UserServiceImpl
func NewInstanceMysqlUserServices(repository *repoinjection.RepositoryInjection) UserServices {
	return &UserServiceImpl{Repository: repository}
}
