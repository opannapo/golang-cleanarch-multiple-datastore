package mysqlrepository

import (
	"app/app/v1/entities"
	"app/app/v1/repository"
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
)

//CredentialRepoImpl implement
type CredentialRepoImpl struct {
	Db *gorm.DB
}

//Insert credential user
func (instance *CredentialRepoImpl) Insert(data *entities.Credential) (tx *gorm.DB, err error) {
	//update credential Signature to md5
	newSignature := md5.Sum([]byte(data.Signature))
	data.Signature = hex.EncodeToString(newSignature[:])
	tx = instance.Db.Begin()
	err = tx.Create(&data).Error

	return
}

//GetByKeySignature get by key and signature
func (instance *CredentialRepoImpl) GetByKeySignature(key string, signature string) (result entities.Credential, err error) {
	//convert credential signature to md5
	newSignature := md5.Sum([]byte(signature))
	err = instance.Db.
		Preload("User").
		Where("`key`=?", key).
		Where("`signature`=?", hex.EncodeToString(newSignature[:])).
		First(&result).Error

	return
}

//NewInstanceMsqlCredentialRepoImpl new instance
func NewInstanceMsqlCredentialRepoImpl(db *gorm.DB) repository.CredentialRepo {
	return &CredentialRepoImpl{Db: db}
}
