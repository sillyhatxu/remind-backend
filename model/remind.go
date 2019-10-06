package model

import (
	"github.com/sillyhatxu/remind-backend/enums"
	"time"
)

type Remind struct {
	Id               int64                 `json:"id"`
	RemindTime       time.Time             `json:"remind_time"`
	RemindType       enums.RemindType      `json:"remind_type"`
	RemindData       string                `json:"remind_data"`
	RemindAdvance    enums.RemindAdvance   `json:"remind_advance"`
	RemindFrequency  enums.RemindFrequency `json:"remind_frequency"`
	IsLunar          bool                  `json:"is_lunar"`
	Status           enums.RemindStatus    `json:"status"`
	LastModifiedTime time.Time             `json:"last_modified_time"`
	CreatedTime      time.Time             `json:"created_time"`
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
