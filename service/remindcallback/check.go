package remindcallback

import (
	"fmt"
	"github.com/sillyhatxu/remind-backend/constants"
	"github.com/sillyhatxu/remind-backend/dto"
	"github.com/sillyhatxu/remind-backend/service/remindbirthday"
)

func Check(slackDTO dto.SlackDTO) error {
	if slackDTO.CallbackId == constants.CallbackIdBirthday {
		return remindbirthday.Birthday(slackDTO)
	} else {
		return fmt.Errorf("unsupported callback_id")
	}
}
