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

//TopicTypeRepoImpl implementation of interface
type TopicTypeRepoImpl struct {
	RedisClient *redis.Client
}

//GetAll TopicTypeRepoImpl get all data, result as slice and err :: from Redis
func (instance *TopicTypeRepoImpl) GetAll() (result []*entities.TopicType, err error) {
	resultAsJSON, err := instance.RedisClient.Get(context.Background(), constant.RedisKeyTopicTypeAll).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}

	err = json.Unmarshal([]byte(resultAsJSON), &result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Load from Redis")

	return
}

//GetByLabel TopicTypeRepoImpl get by label  :: from Redis
func (instance *TopicTypeRepoImpl) GetByLabel(label string) (result entities.TopicType, err error) {
	panic("implement me")
}

//Insert TopicTypeRepoImpl, insert one
func (instance *TopicTypeRepoImpl) Insert(data *entities.TopicType) (err error) {
	panic("implement me")
}

//Inserts TopicTypeRepoImpl, insert multiple to Redis :: from Redis
func (instance *TopicTypeRepoImpl) Inserts(data []*entities.TopicType) (tx *repository.DbTransactionType, err error) {
	tx = &repository.DbTransactionType{
		GormTX:  nil,
		RedisTX: nil,
	}

	dataJSON, _ := json.Marshal(data)
	result, err := instance.RedisClient.Set(context.Background(), constant.RedisKeyTopicTypeAll, dataJSON, 0).Result()
	fmt.Printf("result %+v err %+v", result, err)
	if err != nil {
		fmt.Printf("result %+v err %+v", result, err)
		panic(err)
	}

	return
}

//Upserts TopicTypeRepoImpl, Upsert multiple to existing key Redis :: from Redis
func (instance *TopicTypeRepoImpl) Upserts(data []*entities.TopicType) (tx *repository.DbTransactionType, err error) {
	tx = &repository.DbTransactionType{
		GormTX:  nil,
		RedisTX: nil,
	}

	ctx := context.Background()

	//Check Exists
	existsResult, err := instance.RedisClient.Do(ctx,
		constant.RedisCmdExists,
		constant.RedisKeyTopicTypeAll,
	).Result()
	fmt.Printf("existsResult %v ", existsResult)

	if existsResult.(int64) == 1 {
		//getAll
		resultCurrent, err := instance.GetAll()
		if err == nil {
			//Join
			resultCurrent = append(resultCurrent, data...)
			dataJSON, _ := json.Marshal(resultCurrent)

			//Replace SET
			result, err := instance.RedisClient.Set(ctx, constant.RedisKeyTopicTypeAll, dataJSON, 0).Result()
			fmt.Printf("result %+v err %+v", result, err)
			if err != nil {
				fmt.Printf("result %+v err %+v", result, err)
				panic(err)
			}
		}
	}

	return
}

//NewInstanceTopicTypeRepoImpl new instance of TopicTypeRepoImpl :: from Redis
func NewInstanceTopicTypeRepoImpl(redisClient *redis.Client) repository.TopicTypeRepo {
	return &TopicTypeRepoImpl{RedisClient: redisClient}
}
