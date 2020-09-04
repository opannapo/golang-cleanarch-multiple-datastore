package firestorerepository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"cloud.google.com/go/firestore"
	"context"
	"strconv"
)

//TopicTypeRepoImpl implementation of interface
type TopicTypeRepoImpl struct {
	FirestoreClient *firestore.Client
}

//GetAll TopicTypeRepoImpl get all data, result as slice and err
func (instance *TopicTypeRepoImpl) GetAll() (result []*entities.TopicType, err error) {
	panic("implement me")
}

//GetByLabel TopicTypeRepoImpl get by label
func (instance *TopicTypeRepoImpl) GetByLabel(label string) (result entities.TopicType, err error) {
	panic("implement me")
}

//Insert TopicTypeRepoImpl, insert one
func (instance *TopicTypeRepoImpl) Insert(data *entities.TopicType) (err error) {
	panic("implement me")
}

//Inserts TopicTypeRepoImpl, insert multiple
/*func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (tx *gorm.DB, err error) {
	for i := range data {
		tmpData := data[i]
		id := strconv.FormatInt(tmpData.ID, 10)
		topicTypeMap := map[string]interface{}{
			"id":    tmpData.ID,
			"label": tmpData.Label,
		}
		_, err = instance.FirestoreClient.Collection("topic_type").
			Doc(id).
			Set(context.Background(), topicTypeMap)

		if err != nil {
			break
		}
	}

	return
}*/

//Inserts TopicTypeRepoImpl, insert multiple to Firestore
func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (tx *repository.TransactionStruct, err error) {
	for i := range data {
		tmpData := data[i]
		id := strconv.FormatInt(tmpData.ID, 10)
		topicTypeMap := map[string]interface{}{
			"id":    tmpData.ID,
			"label": tmpData.Label,
		}
		_, err = instance.FirestoreClient.Collection("topic_type").
			Doc(id).
			Set(context.Background(), topicTypeMap)

		if err != nil {
			break
		}
	}

	return
}

//NewInstanceFirestoreTopicTypeRepoImpl new instance of TopicTypeRepoImpl
func NewInstanceFirestoreTopicTypeRepoImpl(firestoreClient *firestore.Client) repository.TopicTypeRepo {
	return &TopicTypeRepoImpl{FirestoreClient: firestoreClient}
}
