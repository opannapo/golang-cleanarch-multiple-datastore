package redis

import (
	"app/app/v1/apis/constant"
	"app/app/v1/entities"
	"app/app/v1/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type TopicTypeRepoImpl struct {
	RedisClient *redis.Client
}

func (instance *TopicTypeRepoImpl) GetAll() (result []*entities.TopicType, err error) {
	resultAsJson, err := instance.RedisClient.Get(context.Background(), constant.RedisKeyTopicTypeAll).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}

	err = json.Unmarshal([]byte(resultAsJson), &result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Load from Redis")

	return
}

func (instance *TopicTypeRepoImpl) GetByLabel(label string) (result entities.TopicType, err error) {
	panic("implement me")
}

func (instance *TopicTypeRepoImpl) Insert(data *entities.TopicType) (err error) {
	panic("implement me")
}

/*func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (tx *gorm.DB, err error) {
	instance.RedisClient.Set(context.Background(),constant.RedisKeyTopicTypeAll,data,10000)
}*/

func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (tx *repository.TransactionStruct, err error) {
	tx = &repository.TransactionStruct{
		GormTX:  nil,
		RedisTX: nil,
	}

	dataJson, _ := json.Marshal(data)
	result, err := instance.RedisClient.Set(context.Background(), constant.RedisKeyTopicTypeAll, dataJson, 0).Result()
	fmt.Printf("result %+v err %+v", result, err)
	if err != nil {
		fmt.Printf("result %+v err %+v", result, err)
		panic(err)
	}

	return
}

func NewInstanceTopicTypeRepoImpl(redisClient *redis.Client) repository.TopicTypeRepo {
	return &TopicTypeRepoImpl{RedisClient: redisClient}
}