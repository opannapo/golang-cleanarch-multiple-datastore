package mysql_repository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"github.com/jinzhu/gorm"
)

type CredentialRepoImpl struct {
	Db *gorm.DB
}

func (instance *CredentialRepoImpl) GetByKeySignature(key string, signature string) (result *entities.Credential, err error) {
	err = instance.Db.
		Where("key=?", key).
		Where("signature=?", signature).
		Find(&result).Error
	return
}

func NewInstanceMsqlCredentialRepoImpl(db *gorm.DB) repository.CredentialRepo {
	return &CredentialRepoImpl{Db: db}
}
