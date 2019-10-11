package dao

import (
	"github.com/sillyhatxu/remind-backend/enums"
	"github.com/sillyhatxu/remind-backend/model"
)

const insertRemindSQL = `
	insert into remind (remind_time, remind_type, remind_data, remind_advance, remind_frequency, is_lunar, status) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
`

func InsertRemind(input *model.Remind) error {
	_, err := client.Insert(insertRemindSQL, input.RemindTime, input.RemindType, input.RemindData, input.RemindAdvance, input.RemindFrequency, input.IsLunar, input.Status)
	return err
}

const findAllRemindListSQL = `
	SELECT * FROM remind where status = ?
`

func FindAllRemindList() ([]model.Remind, error) {
	var array []model.Remind
	err := client.FindList(findAllRemindListSQL, &array, enums.RemindStatusEnabled)
	if err != nil {
		return nil, err
	}
	return array, nil
}
