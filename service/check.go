package service

import (
	"fmt"
	"github.com/sillyhatxu/remind-backend/constants"
	"github.com/sillyhatxu/remind-backend/dto"
)

func Check(slackDTO dto.SlackDTO) error {
	if slackDTO.CallbackId == constants.CallbackIdBirthday {
		return Birthday(slackDTO)
	} else {
		return fmt.Errorf("unsupported callback_id")
	}
}
