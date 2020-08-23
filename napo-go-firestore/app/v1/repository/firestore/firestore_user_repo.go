package firestorerepository

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/repository"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/api/iterator"
	"log"
	"strconv"
)

//UserRepoImpl implementation of interface
type UserRepoImpl struct {
	FirestoreClient *firestore.Client
}

//GetAll return users
func (instance *UserRepoImpl) GetAll() (result []*entities.User, err error) {
	documentIterator := instance.FirestoreClient.
		Collection("users").
		Documents(context.Background())

	for {
		doc, err := documentIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		data := doc.Data()
		fmt.Println(data)
		user := entities.User{
			ID:        data["id"].(int64),
			FirstName: data["firstName"].(string),
			LastName:  data["lastName"].(string),
		}
		result = append(result, &user)
	}

	return
}

//GetByID return user by id
func (instance *UserRepoImpl) GetByID(id int) (result entities.User, err error) {
	panic("implement me")
}

//GetByCredential return user get by credential
func (instance *UserRepoImpl) GetByCredential(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

//Insert insert new user
func (instance *UserRepoImpl) Insert(param *param.UserCreate) (tx *gorm.DB, err error) {
	user := param.User
	id := strconv.FormatInt(user.ID, 10)
	userMap := map[string]interface{}{
		"id":              user.ID,
		"firstName":       user.FirstName,
		"lastName":        user.LastName,
		"following_topic": param.FollowingTopic,
	}
	_, err = instance.FirestoreClient.Collection("users").
		Doc(id).
		Create(context.Background(), userMap)
	return
}

//Update update user
func (instance *UserRepoImpl) Update(user *entities.User) (err error) {
	panic("implement me")
}

//Delete delete user
func (instance *UserRepoImpl) Delete(user *entities.User) (err error) {
	panic("implement me")
}

//NewInstanceFirestoreUserRepoImpl new instance of UserRepoImpl
func NewInstanceFirestoreUserRepoImpl(firestoreClient *firestore.Client) repository.UserRepo {
	return &UserRepoImpl{FirestoreClient: firestoreClient}
}
