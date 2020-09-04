package repoinjection

import (
	"app/app/v1/repository"
	repoImplFirestore "app/app/v1/repository/firestore"
	repoImplMysql "app/app/v1/repository/mysql"
	repoImplRedis "app/app/v1/repository/redis"
	"cloud.google.com/go/firestore"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

//RepositoryInjection struct
type RepositoryInjection struct {
	*mysqlRepoInjected
	*firestoreRepoInjected
	*redisRepoInjected
	db              *gorm.DB
	firestoreClient *firestore.Client
	redisClient     *redis.Client
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

type redisRepoInjected struct {
	RedisTopicTypeRepo repository.TopicTypeRepo
}

//NewInstanceRepositoryInjection new instance of RepositoryInjection struct
func NewInstanceRepositoryInjection(db *gorm.DB, firestoreClient *firestore.Client, redisClient *redis.Client) *RepositoryInjection {
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

	rd := redisRepoInjected{
		RedisTopicTypeRepo: repoImplRedis.NewInstanceTopicTypeRepoImpl(redisClient),
	}

	return &RepositoryInjection{
		mysqlRepoInjected:     &ms,
		firestoreRepoInjected: &fs,
		redisRepoInjected:     &rd,
		db:                    db,
		firestoreClient:       firestoreClient,
		redisClient:           redisClient,
	}
}
