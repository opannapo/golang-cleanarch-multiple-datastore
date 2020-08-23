package param

import (
	"app/app/v1/entities"
)

//UserCreate struct for parameter user create
type UserCreate struct {
	User           *entities.User       `json:"user" binding:"required"`
	FollowingTopic []*string            `json:"following_topic"`
	Credential     *entities.Credential `json:"credential,omitempty"`
}
