package entities

//Credential --> Table Fields Entity
type Credential struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Key       string `json:"key"`
	Signature string `json:"signature"`
	Type      int64  `json:"type"`
	UserID    int64  `json:"user_id"`
	User      *User  `json:"user" gorm:"foreignkey:ID;association_foreignkey:UserID"`
}

//TableName --> Table Name
func (Credential) TableName() string {
	return "credential"
}
