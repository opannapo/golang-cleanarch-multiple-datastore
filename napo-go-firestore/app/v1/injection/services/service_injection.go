package services

import (
	"app/app/v1/injection/repositories"
	serviceInterface "app/app/v1/services"
	serviceImplFirestore "app/app/v1/services/firestore"
	serviceImplGeneral "app/app/v1/services/generalservices"
	serviceImplMysql "app/app/v1/services/mysql"
)

//ServiceInjection struct
type ServiceInjection struct {
	*mysqlServicesInjected
	*firestoreServicesInjected
	*generalServicesInjected
	repository *repositories.RepositoryInjection
}

type mysqlServicesInjected struct {
	MysqlUserService               serviceInterface.UserServices
	MysqlTopicTypeService          serviceInterface.TopicTypeServices
	MysqlUserFollowingTopicService serviceInterface.UserFollowingTopicServices
	MysqlAuthService               serviceInterface.AuthServices
}

type firestoreServicesInjected struct {
	FirestoreUserService      serviceInterface.UserServices
	FirestoreTopicTypeService serviceInterface.TopicTypeServices
}

type generalServicesInjected struct {
	AuthService serviceInterface.AuthServices
}

//NewInstanceServiceInjection new instance of ServiceInjection, & generate all services Instance
func NewInstanceServiceInjection(repository *repositories.RepositoryInjection) *ServiceInjection {
	ms := mysqlServicesInjected{
		MysqlUserService:               serviceImplMysql.NewInstanceMysqlUserServices(repository),
		MysqlTopicTypeService:          serviceImplMysql.NewInstanceMysqlTopicTypeServices(repository),
		MysqlUserFollowingTopicService: serviceImplMysql.NewInstanceMysqlUserFollowingTopicService(repository),
	}

	fs := firestoreServicesInjected{
		FirestoreUserService:      serviceImplFirestore.NewInstanceFirestoreUserService(repository),
		FirestoreTopicTypeService: serviceImplFirestore.NewInstanceFirestoreTopicTypeService(repository),
	}

	gs := generalServicesInjected{
		AuthService: serviceImplGeneral.NewInstanceAuthService(repository),
	}

	return &ServiceInjection{
		mysqlServicesInjected:     &ms,
		firestoreServicesInjected: &fs,
		generalServicesInjected:   &gs,
		repository:                repository,
	}
}
