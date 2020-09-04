package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

//TransactionStruct dynamic return type to handling transaction commit & roolback for each database type
type TransactionStruct struct {
	GormTX  *gorm.DB
	RedisTX *redis.Tx
}
