package serviceinjection

import (
	"app/app/v1/injection/repositories"
	serviceInterface "app/app/v1/services"
)

//ServiceInjection struct
type ServiceInjection struct {
	*servicesInjected
	repository *repoinjection.RepositoryInjection
}

type servicesInjected struct {
	UserService               serviceInterface.UserServices
	TopicTypeService          serviceInterface.TopicTypeServices
	UserFollowingTopicService serviceInterface.UserFollowingTopicServices
	AuthService               serviceInterface.AuthServices
}

//NewInstanceServiceInjection new instance of ServiceInjection, & generate all services Instance
func NewInstanceServiceInjection(repository *repoinjection.RepositoryInjection) *ServiceInjection {
	ms := servicesInjected{
		UserService:               serviceInterface.NewInstanceMysqlUserServices(repository),
		TopicTypeService:          serviceInterface.NewInstanceMysqlTopicTypeServices(repository),
		UserFollowingTopicService: serviceInterface.NewInstanceMysqlUserFollowingTopicService(repository),
		AuthService:               serviceInterface.NewInstanceAuthService(repository),
	}

	return &ServiceInjection{
		servicesInjected: &ms,
		repository:       repository,
	}
}
