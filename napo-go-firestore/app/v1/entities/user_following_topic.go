package entities

//UserFollowingTopic --> Table Fields Entity
type UserFollowingTopic struct {
	ID          int64      `gorm:"primary_key"  json:"id"`
	UserID      int64      `gorm:"foreignkey"  json:"user_id"`
	TopicTypeID int64      `json:"topic_type_id"`
	TopicType   *TopicType `gorm:"foreignkey:TopicTypeID";association_foreignkey:"ID" json:"topic_type"`
}

//TableName UserFollowingTopic
func (*UserFollowingTopic) TableName() string {
	return "user_following_topic"
}
