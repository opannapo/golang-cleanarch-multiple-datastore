package entities

type TopicType struct {
	Id    int64  `gorm:"primary_key"  json:"id"`
	Label string `json:"label"`
}

func (*TopicType) TableName() string {
	return "topic_type"
}
