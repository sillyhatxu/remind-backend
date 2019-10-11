package dao

import (
	"encoding/json"
	"github.com/sillyhatxu/remind-backend/enums"
	"github.com/sillyhatxu/remind-backend/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInsertRemind(t *testing.T) {
	birthday1 := model.RemindDataBirthday{
		UserName:    "徐士宽",
		RemindTitle: "Happy Birthday : XXXXX XXXX",
		RemindText:  "生日快乐 阴历 : 五月十四",
	}
	birthday2 := model.RemindDataBirthday{
		UserName:    "孙博文",
		RemindTitle: "Happy Birthday : XXXXX XXXX",
		RemindText:  "生日快乐 阳历 : 1988/12/01",
	}
	var array []model.RemindDataBirthday
	array = append(array, birthday1)
	array = append(array, birthday2)

	for _, birthday := range array {
		birthdayJSON, err := json.Marshal(birthday)
		assert.Nil(t, err)
		err = InsertRemind(&model.Remind{
			RemindTime:      time.Now(),
			RemindType:      enums.RemindTypeBirthday,
			RemindData:      string(birthdayJSON),
			RemindAdvance:   enums.RemindAdvanceSevenDay,
			RemindFrequency: enums.RemindFrequencyEveryYear,
			IsLunar:         true,
			Status:          enums.RemindStatusEnabled,
		})
		assert.Nil(t, err)
	}
}

func TestFindAllRemindList(t *testing.T) {
	res, err := FindAllRemindList()
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, len(res), 2)
}
