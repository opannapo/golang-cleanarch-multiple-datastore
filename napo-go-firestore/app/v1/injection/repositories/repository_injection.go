package repoinjection

import (
	"app/app/v1/repository"
	repoImplFirestore "app/app/v1/repository/firestore"
	repoImplMysql "app/app/v1/repository/mysql"
	"cloud.google.com/go/firestore"
	"github.com/jinzhu/gorm"
)

//RepositoryInjection struct
type RepositoryInjection struct {
	*mysqlRepoInjected
	*firestoreRepoInjected
	db              *gorm.DB
	firestoreClient *firestore.Client
}

type mysqlRepoInjected struct {
	MysqlUserRepo               repository.UserRepo
	MysqlTopicTypeRepo          repository.TopicTypeRepo
	MysqlUserFollowingTopicRepo repository.UserFollowingTopicRepo
	MysqlCredentialRepo         repository.CredentialRepo
}

type firestoreRepoInjected struct {
	FirestoreUserRepo      repository.UserRepo
	FirestoreTopicTypeRepo repository.TopicTypeRepo
}

//NewInstanceRepositoryInjection new instance of RepositoryInjection struct
func NewInstanceRepositoryInjection(db *gorm.DB, firestoreClient *firestore.Client) *RepositoryInjection {
	ms := mysqlRepoInjected{
		MysqlUserRepo:               repoImplMysql.NewInstanceMysqlUserRepoImpl(db),
		MysqlTopicTypeRepo:          repoImplMysql.NewInstanceMysqlTopicTypeRepoImpl(db),
		MysqlUserFollowingTopicRepo: repoImplMysql.NewInstanceMysqlUserFollowingTopicRepoImpl(db),
		MysqlCredentialRepo:         repoImplMysql.NewInstanceMsqlCredentialRepoImpl(db),
	}

	fs := firestoreRepoInjected{
		FirestoreUserRepo:      repoImplFirestore.NewInstanceFirestoreUserRepoImpl(firestoreClient),
		FirestoreTopicTypeRepo: repoImplFirestore.NewInstanceFirestoreTopicTypeRepoImpl(firestoreClient),
	}
	return &RepositoryInjection{
		mysqlRepoInjected:     &ms,
		firestoreRepoInjected: &fs,
		db:                    db,
		firestoreClient:       firestoreClient,
	}
}
