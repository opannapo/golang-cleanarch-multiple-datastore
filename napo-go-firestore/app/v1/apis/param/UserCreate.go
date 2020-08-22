package param

import (
	"app/app/v1/entities"
)

type UserCreate struct {
	User           *entities.User       `json:"user" binding:"required"`
	FollowingTopic []*string            `json:"following_topic"`
	Credential     *entities.Credential `json:"credential,omitempty"`
}
