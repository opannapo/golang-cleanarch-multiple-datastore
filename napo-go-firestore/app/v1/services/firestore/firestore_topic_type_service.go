package firestore

import (
	"app/app/v1/entities"
	"app/app/v1/injection"
	"app/app/v1/services"
)

type TopicTypeServiceImpl struct {
	Repository *injection.RepositoryInjection
}

func (instance *TopicTypeServiceImpl) GetAll() (result []*entities.TopicType, err error) {
	panic("implement me")
}

func (instance *TopicTypeServiceImpl) GetOneByLabel(label string) (result entities.TopicType, err error) {
	panic("implement me")
}

func (instance *TopicTypeServiceImpl) Insert(data *entities.TopicType) (err error) {
	panic("implement me")
}

func (instance *TopicTypeServiceImpl) Inserts(data []*entities.TopicType) (err error) {
	/*for i := range data {
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
	}*/

	return
}

func NewInstanceFirestoreTopicTypeService(repository *injection.RepositoryInjection) services.TopicTypeServices {
	return &TopicTypeServiceImpl{Repository: repository}
}
