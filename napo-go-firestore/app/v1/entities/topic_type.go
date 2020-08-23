package entities

//TopicType --> Table Fields Entity
type TopicType struct {
	ID    int64  `gorm:"primary_key"  json:"id"`
	Label string `json:"label"`
}

//TableName TopicType --> Table name
func (*TopicType) TableName() string {
	return "topic_type"
}
