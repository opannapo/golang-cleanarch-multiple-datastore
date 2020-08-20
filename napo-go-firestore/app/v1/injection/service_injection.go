package injection

import (
	serviceInterface "app/app/v1/services"
	serviceImplFirestore "app/app/v1/services/firestore"
	serviceImplMysql "app/app/v1/services/mysql"
)

type ServiceInjection struct {
	*mysqlServicesInjected
	*firestoreServicesInjected
	repository *RepositoryInjection
}

type mysqlServicesInjected struct {
	MysqlUserService               serviceInterface.UserServices
	MysqlTopicTypeService          serviceInterface.TopicTypeServices
	MysqlUserFollowingTopicService serviceInterface.UserFollowingTopicServices
}

type firestoreServicesInjected struct {
	FirestoreUserService      serviceInterface.UserServices
	FirestoreTopicTypeService serviceInterface.TopicTypeServices
}

func NewInstanceServiceInjection(repository *RepositoryInjection) *ServiceInjection {
	ms := mysqlServicesInjected{
		MysqlUserService:               serviceImplMysql.NewInstanceMysqlUserServices(repository),
		MysqlTopicTypeService:          serviceImplMysql.NewInstanceMysqlTopicTypeServices(repository),
		MysqlUserFollowingTopicService: serviceImplMysql.NewInstanceMysqlUserFollowingTopicService(repository),
	}

	fs := firestoreServicesInjected{
		FirestoreUserService:      serviceImplFirestore.NewInstanceFirestoreUserService(repository),
		FirestoreTopicTypeService: serviceImplFirestore.NewInstanceFirestoreTopicTypeService(repository),
	}

	return &ServiceInjection{
		mysqlServicesInjected:     &ms,
		firestoreServicesInjected: &fs,
		repository:                repository,
	}
}
