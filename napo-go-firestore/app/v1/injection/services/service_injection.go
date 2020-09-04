package services

import (
	"app/app/v1/injection/repositories"
	serviceInterface "app/app/v1/services"
)

//ServiceInjection struct
type ServiceInjection struct {
	*servicesInjected
	repository *repositories.RepositoryInjection
}

type servicesInjected struct {
	MysqlUserService               serviceInterface.UserServices
	MysqlTopicTypeService          serviceInterface.TopicTypeServices
	MysqlUserFollowingTopicService serviceInterface.UserFollowingTopicServices
	MysqlAuthService               serviceInterface.AuthServices
	AuthService                    serviceInterface.AuthServices
}

//NewInstanceServiceInjection new instance of ServiceInjection, & generate all services Instance
func NewInstanceServiceInjection(repository *repositories.RepositoryInjection) *ServiceInjection {
	ms := servicesInjected{
		MysqlUserService:               serviceInterface.NewInstanceMysqlUserServices(repository),
		MysqlTopicTypeService:          serviceInterface.NewInstanceMysqlTopicTypeServices(repository),
		MysqlUserFollowingTopicService: serviceInterface.NewInstanceMysqlUserFollowingTopicService(repository),
		AuthService:                    serviceInterface.NewInstanceAuthService(repository),
	}

	return &ServiceInjection{
		servicesInjected: &ms,
		repository:       repository,
	}
}
