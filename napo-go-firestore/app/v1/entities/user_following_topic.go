package entities

type UserFollowingTopic struct {
	Id          int64      `gorm:"primary_key"  json:"id"`
	UserId      int64      `gorm:"foreignkey"  json:"user_id"`
	TopicTypeId int64      `json:"topic_type_id"`
	TopicType   *TopicType `gorm:"foreignkey:TopicTypeId";association_foreignkey:"Id" json:"topic_type"`
}

func (*UserFollowingTopic) TableName() string {
	return "user_following_topic"
}
