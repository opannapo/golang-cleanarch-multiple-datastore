package mysqlrepository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"github.com/jinzhu/gorm"
)

//CredentialRepoImpl implement
type CredentialRepoImpl struct {
	Db *gorm.DB
}

//Insert credential user
func (instance *CredentialRepoImpl) Insert(data *entities.Credential) (tx *gorm.DB, err error) {
	tx = instance.Db.Begin()
	err = instance.Db.Create(&data).Error
	return
}

//GetByKeySignature get by key and signature
func (instance *CredentialRepoImpl) GetByKeySignature(key string, signature string) (result entities.Credential, err error) {
	err = instance.Db.
		Preload("User").
		Where("`key`=?", key).
		Where("`signature`=?", signature).
		First(&result).Error

	return
}

//NewInstanceMsqlCredentialRepoImpl new instance
func NewInstanceMsqlCredentialRepoImpl(db *gorm.DB) repository.CredentialRepo {
	return &CredentialRepoImpl{Db: db}
}
