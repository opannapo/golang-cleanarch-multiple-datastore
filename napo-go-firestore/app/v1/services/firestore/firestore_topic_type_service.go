package firestoreservices

import (
	"app/app/v1/entities"
	"app/app/v1/injection/repositories"
	"app/app/v1/services"
)

//TopicTypeServiceImpl struct
type TopicTypeServiceImpl struct {
	Repository *repositories.RepositoryInjection
}

//GetAll get all users
func (instance *TopicTypeServiceImpl) GetAll() (result []*entities.TopicType, err error) {
	panic("implement me")
}

//GetOneByLabel get one by label
func (instance *TopicTypeServiceImpl) GetOneByLabel(label string) (result entities.TopicType, err error) {
	panic("implement me")
}

//Insert insert new
func (instance *TopicTypeServiceImpl) Insert(data *entities.TopicType) (err error) {
	panic("implement me")
}

//Inserts insert multiple
func (instance *TopicTypeServiceImpl) Inserts(data []*entities.TopicType) (err error) {
	/*for i := range data {
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
	}*/

	return
}

//NewInstanceFirestoreTopicTypeService new instance of TopicTypeServiceImpl
func NewInstanceFirestoreTopicTypeService(repository *repositories.RepositoryInjection) services.TopicTypeServices {
	return &TopicTypeServiceImpl{Repository: repository}
}
