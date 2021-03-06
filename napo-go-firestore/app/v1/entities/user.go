package entities

//User --> Table Fields Entity
type User struct {
	ID             int64                 `gorm:"primary_key" json:"id"`
	FirstName      string                `json:"first_name"`
	LastName       string                `json:"last_name"`
	ProfilePicture string                `json:"profile_picture,omitempty"`
	BirthDate      int64                 `json:"birth_date,omitempty"`
	Phone          string                `json:"phone,omitempty"`
	Email          string                `json:"email,omitempty"`
	FollowingTopic []*UserFollowingTopic `gorm:"foreignkey:UserID;association_foreignkey:ID" json:"following_topic,omitempty"`
	Credential     *Credential           `gorm:"foreignkey:UserID;association_foreignkey:ID" json:"credential,omitempty"`
}

//TableName User
func (*User) TableName() string {
	return "user"
}
