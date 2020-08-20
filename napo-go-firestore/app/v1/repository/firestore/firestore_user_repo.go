package firestore_repository

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

type UserRepoImpl struct {
	FirestoreClient *firestore.Client
}

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
			Id:        data["id"].(int64),
			FirstName: data["firstName"].(string),
			LastName:  data["lastName"].(string),
		}
		result = append(result, &user)
	}

	return
}

func (instance *UserRepoImpl) GetById(id int) (result entities.User, err error) {
	panic("implement me")
}

func (instance *UserRepoImpl) GetByCredential(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

func (instance *UserRepoImpl) Insert(param *param.UserCreate) (err error, tx *gorm.DB) {
	user := param.User
	id := strconv.FormatInt(user.Id, 10)
	userMap := map[string]interface{}{
		"id":              user.Id,
		"firstName":       user.FirstName,
		"lastName":        user.LastName,
		"following_topic": param.FollowingTopic,
	}
	_, err = instance.FirestoreClient.Collection("users").
		Doc(id).
		Create(context.Background(), userMap)
	return
}

func (instance *UserRepoImpl) Update(user *entities.User) (err error) {
	panic("implement me")
}

func (instance *UserRepoImpl) Delete(user *entities.User) (err error) {
	panic("implement me")
}

func NewInstanceFirestoreUserRepoImpl(firestoreClient *firestore.Client) repository.UserRepo {
	return &UserRepoImpl{FirestoreClient: firestoreClient}
}
