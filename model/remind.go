package model

import (
	"github.com/sillyhatxu/remind-backend/enums"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"time"
)

type Remind struct {
	Id         int64            `json:"id"`
	RemindTime time.Time        `json:"remind_time"`
	RemindType enums.RemindType `json:"remind_type"`

	LoginName string       `json:"login_name"`
	Password  string       `json:"password"`
	Channel   user.Channel `json:"channel"`
	Type      user.Type    `json:"type"`

	Status           enums.RemindStatus `json:"status"`
	LastModifiedTime time.Time          `json:"last_modified_time"`
	CreatedTime      time.Time          `json:"created_time"`
}
