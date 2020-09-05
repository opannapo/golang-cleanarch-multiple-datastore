package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

//DbTransactionType dynamic return type to handling transaction commit & roolback for each database type
type DbTransactionType struct {
	GormTX  *gorm.DB
	RedisTX *redis.Tx
}
