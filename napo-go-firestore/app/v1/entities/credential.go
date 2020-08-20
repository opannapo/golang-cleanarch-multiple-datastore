package entities

type Credential struct {
	Id        int64  `json:"id" gorm:"primary_key"`
	Key       string `json:"key"`
	Signature string `json:"signature"`
	Type      int64  `json:"type"`
	UserId    int64  `json:"user_id"`
	User      *User  `json:"user" gorm:"foreignkey:Id;association_foreignkey:UserId"`
}

func (Credential) TableName() string {
	return "credential"
}
