package firestore

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/injection"
	"app/app/v1/services"
)

type UserServiceImpl struct {
	Repository *injection.RepositoryInjection
}

func (instance *UserServiceImpl) GetUserForAuth(username string, password string) (result *entities.User, err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) GetUsers() (result []*entities.User, err error) {
	/*documentIterator := instance.FirestoreClient.
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
	}*/

	return
}

func (instance *UserServiceImpl) GetUser(id int) (result entities.User, err error) {
	panic("Implement me")
}

func (instance *UserServiceImpl) AddUser(param *param.UserCreate) (err error) {
	/*user := param.User
	id := strconv.FormatInt(user.Id, 10)
	userMap := map[string]interface{}{
		"id":              user.Id,
		"firstName":       user.FirstName,
		"lastName":        user.LastName,
		"following_topic": param.FollowingTopic,
	}
	_, err = instance.FirestoreClient.Collection("users").
		Doc(id).
		Create(context.Background(), userMap)*/
	return
}

func (instance *UserServiceImpl) UpdateUser(user *entities.User) (err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) DeleteUser(user *entities.User) (err error) {
	panic("implement me")
}

func NewInstanceFirestoreUserService(repository *injection.RepositoryInjection) services.UserServices {
	return &UserServiceImpl{Repository: repository}
}
