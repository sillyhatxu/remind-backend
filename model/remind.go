package model

import (
	"github.com/sillyhatxu/remind-backend/enums"
	"time"
)

type Remind struct {
	Id               int64                 `json:"id" mapstructure:"id"`
	RemindTime       time.Time             `json:"remind_time" mapstructure:"remind_time"`
	RemindType       enums.RemindType      `json:"remind_type" mapstructure:"remind_type"`
	RemindData       string                `json:"remind_data" mapstructure:"remind_data"`
	RemindAdvance    enums.RemindAdvance   `json:"remind_advance" mapstructure:"remind_advance"`
	RemindFrequency  enums.RemindFrequency `json:"remind_frequency" mapstructure:"remind_frequency"`
	IsLunar          bool                  `json:"is_lunar" mapstructure:"is_lunar"`
	Status           enums.RemindStatus    `json:"status" mapstructure:"status"`
	LastModifiedTime time.Time             `json:"last_modified_time" mapstructure:"last_modified_time"`
	CreatedTime      time.Time             `json:"created_time" mapstructure:"created_time"`
}

type RemindDataBirthday struct {
	UserName    string `json:"user_name"`
	RemindTitle string `json:"remind_title"`
	RemindText  string `json:"remind_text"`
}

//Fallback:   "Happy Birthday : Shikuan Xu",
//Text:       "生日快乐 阴历 : 五月十四",
//AuthorName: "徐士宽",

//type RemindDataBirthday struct {
//	UserName    string `json:"user_name"`    //author_name
//	RemindTitle string `json:"remind_title"` //fallback
//	RemindText  string `json:"remind_text"`  //text
//	RemindTime  int64  `json:"remind_time"`
//}
