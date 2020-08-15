package listener

import (
	"cloud.google.com/go/firestore"
	"github.com/jinzhu/gorm"
)

type instance struct {
	Db              *gorm.DB
	FirestoreClient *firestore.Client
}

func NewFirestoreListener(db *gorm.DB, firestoreClient *firestore.Client) *instance {
	return &instance{
		Db:              db,
		FirestoreClient: firestoreClient,
	}
}

func (instance *instance) Run() {

}
