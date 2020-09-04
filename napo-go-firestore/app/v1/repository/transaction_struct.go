package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type TransactionStruct struct {
	GormTX  *gorm.DB
	RedisTX *redis.Tx
}
