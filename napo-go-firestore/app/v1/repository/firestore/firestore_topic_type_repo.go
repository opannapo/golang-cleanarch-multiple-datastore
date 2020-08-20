package firestore_repository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/jinzhu/gorm"
	"strconv"
)

type TopicTypeRepoImpl struct {
	FirestoreClient *firestore.Client
}

func (instance *TopicTypeRepoImpl) GetAll() (result []*entities.TopicType, err error) {
	panic("implement me")
}

func (instance *TopicTypeRepoImpl) GetByLabel(label string) (result entities.TopicType, err error) {
	panic("implement me")
}

func (instance *TopicTypeRepoImpl) Insert(data *entities.TopicType) (err error) {
	panic("implement me")
}

func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (err error, tx *gorm.DB) {
	for i := range data {
		tmpData := data[i]
		id := strconv.FormatInt(tmpData.Id, 10)
		topicTypeMap := map[string]interface{}{
			"id":    tmpData.Id,
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

func NewInstanceFirestoreTopicTypeRepoImpl(firestoreClient *firestore.Client) repository.TopicTypeRepo {
	return &TopicTypeRepoImpl{FirestoreClient: firestoreClient}
}
