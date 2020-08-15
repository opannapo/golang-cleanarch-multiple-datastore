package entities

type User struct {
	Id             int64                 `gorm:"primary_key" json:"id"`
	FirstName      string                `json:"first_name"`
	LastName       string                `json:"last_name"`
	ProfilePicture string                `json:"profile_picture,omitempty"`
	BirthDate      int64                 `json:"birth_date,omitempty"`
	Phone          string                `json:"phone,omitempty"`
	Email          string                `json:"email,omitempty"`
	FollowingTopic []*UserFollowingTopic `gorm:"foreignkey:UserId;association_foreignkey:Id" json:"following_topic"`
}

func (*User) TableName() string {
	return "user"
}
