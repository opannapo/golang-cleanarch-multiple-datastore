package injection

import (
	serviceInterface "app/app/v1/services"
	serviceImplFirestore "app/app/v1/services/firestore"
	serviceImplMysql "app/app/v1/services/mysql"
	"cloud.google.com/go/firestore"
	"github.com/jinzhu/gorm"
)

type ServiceInjection struct {
	*mysqlServicesInjected
	*firestoreServicesInjected
	db              *gorm.DB
	firestoreClient *firestore.Client
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

func NewInstanceServiceInjection(db *gorm.DB, firestoreClient *firestore.Client) *ServiceInjection {
	ms := mysqlServicesInjected{
		MysqlUserService:               serviceImplMysql.NewInstanceMysqlUserServices(db),
		MysqlTopicTypeService:          serviceImplMysql.NewInstanceMysqlTopicTypeServices(db),
		MysqlUserFollowingTopicService: serviceImplMysql.NewInstanceMysqlUserFollowingTopicService(db),
	}

	fs := firestoreServicesInjected{
		FirestoreUserService:      serviceImplFirestore.NewInstanceFirestoreUserService(firestoreClient),
		FirestoreTopicTypeService: serviceImplFirestore.NewInstanceFirestoreTopicTypeService(firestoreClient),
	}

	return &ServiceInjection{
		mysqlServicesInjected:     &ms,
		firestoreServicesInjected: &fs,
		db:                        db,
		firestoreClient:           firestoreClient,
	}
}
