package service

import (
	"github.com/sillyhatxu/gin-utils/codes"
	"github.com/sillyhatxu/gin-utils/ginerrors"
	"github.com/sillyhatxu/remind-backend/dto"
	"github.com/sirupsen/logrus"
)

func Birthday(slackDTO dto.SlackDTO) error {
	if slackDTO.Actions == nil || len(slackDTO.Actions) == 0 {
		return ginerrors.Error(codes.FailedPrecondition, "actions is nil")
	}
	logrus.Infof("slackDTO : %#v", slackDTO)
	//action := slackDTO.Actions[0]
	return nil
}
