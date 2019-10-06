package remindbirthday

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/sillyhatxu/environment-config"
	"github.com/sillyhatxu/gin-utils/codes"
	"github.com/sillyhatxu/gin-utils/ginerrors"
	"github.com/sillyhatxu/gocache-client"
	"github.com/sillyhatxu/message-backend/grpc/grpcclient"
	"github.com/sillyhatxu/message-backend/grpc/message"
	"github.com/sillyhatxu/remind-backend/constants"
	"github.com/sillyhatxu/remind-backend/dao"
	"github.com/sillyhatxu/remind-backend/dto"
	"github.com/sillyhatxu/remind-backend/enums"
	"github.com/sillyhatxu/remind-backend/utils"
	"github.com/sirupsen/logrus"
	"time"
)

var birthdayCacheClient *cacheclient.CacheClient
var messageClient *grpcclient.Client

func InitialBirthdayClient() error {
	messageClient = grpcclient.NewGRPCMessageClient(envconfig.Conf.HostMessageGRPC, grpcclient.Timeout(30*time.Second))
	client, err := cacheclient.NewCacheClient()
	if err != nil {
		return err
	}
	birthdayCacheClient = client
	return nil
}

type RemindBirthday struct {
	Id         string             `json:"id"`
	Fallback   string             `json:"fallback"`
	Text       string             `json:"text"`
	AuthorName string             `json:"author_name"`
	Status     enums.RemindStatus `json:"status"`
}

const remindBirthdayKey = "REMIND_BIRTHDAY_KEY"

func Birthday(slackDTO dto.SlackDTO) error {
	if slackDTO.Actions == nil || len(slackDTO.Actions) == 0 {
		return ginerrors.Error(codes.FailedPrecondition, "actions is nil")
	}
	logrus.Infof("slackDTO : %#v", slackDTO)
	//action := slackDTO.Actions[0]
	return nil
}

func GetBirthdayListFromDB() ([]RemindBirthday, error) {
	array, err := dao.FindAllRemindList()
	if err != nil {
		return nil, err
	}
	var resultArray []RemindBirthday
	for _, remind := range array {
		if remind.RemindType == enums.RemindTypeBirthday {
			resultArray = append(resultArray, RemindBirthday{
				Id: utils.GeneratorId(),
				//Fallback
				//Text
				//AuthorName
				//Status
			})
		}
	}
	return resultArray, nil
}

func GetBirthdayList() ([]RemindBirthday, error) {
	var resultArray []RemindBirthday
	err := birthdayCacheClient.GetObj(remindBirthdayKey, &resultArray)
	if err != nil {
		return nil, err
	}
	if resultArray != nil {
		return resultArray, nil
	}
	return GetBirthdayListFromDB()
}

func AddBirthdayRemind(input dto.BirthdayRemindDTO) error {

	return nil
}

func SendBirthdayRemind(input dto.BirthdayRemindDTO) error {
	gotItAction := &message.AttachmentAction{
		Name:  "ACTION_001",
		Text:  "I got it",
		Type:  constants.AttachmentActionButton,
		Value: input.Id,
	}
	return messageClient.SendMessage(&message.AttachmentRequest{
		Fallback:       input.Fallback,
		Text:           input.Text,
		Color:          constants.Pink,
		AuthorName:     input.AuthorName,
		AuthorIcon:     constants.AuthorIconBirthday,
		AttachmentType: constants.AttachmentType,
		CallbackId:     constants.CallbackIdBirthday,
		Actions:        []*message.AttachmentAction{gotItAction},
		Timestamp:      ptypes.TimestampNow(),
	})
}
