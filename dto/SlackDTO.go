package dto

type SlackDTO struct {
	Type            string                  `json:"type"`
	Actions         []SlackActionDTO        `json:"actions"`
	CallbackId      string                  `json:"callback_id"`
	Team            SlackTeamDTO            `json:"team"`
	Channel         SlackChannelDTO         `json:"channel"`
	User            SlackUserDTO            `json:"user"`
	OriginalMessage SlackOriginalMessageDTO `json:"original_message"`
}

type SlackActionDTO struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type SlackTeamDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SlackChannelDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SlackUserDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SlackOriginalMessageDTO struct {
	Type        string                              `json:"type"`
	Subtype     string                              `json:"subtype"`
	Text        string                              `json:"text"`
	Timestamp   string                              `json:"ts"`
	BotId       string                              `json:"bot_id"`
	Attachments []SlackOriginalMessageAttachmentDTO `json:"attachments"`
}

type SlackOriginalMessageAttachmentDTO struct {
	Id         int                                        `json:"id"`
	CallbackId string                                     `json:"callback_id"`
	Fallback   string                                     `json:"fallback"`
	Text       string                                     `json:"text"`
	Color      string                                     `json:"color"`
	Actions    []SlackOriginalMessageAttachmentActionsDTO `json:"actions"`
}

type SlackOriginalMessageAttachmentActionsDTO struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Style string `json:"style"`
}

/**

  "original_message": {
    "type": "message",
    "subtype": "bot_message",
    "text": "Would you like to play a game?",
    "ts": "1570287687.003100",
    "bot_id": "BNMNB35S7",
    "attachments": [
      {
        "callback_id": "sillyhat-birthday",
        "fallback": "You are unable to choose a game",
        "text": "Choose a game to play",
        "id": 1,
        "color": "3AA3E3",
        "actions": [
          {
            "id": "1",
            "name": "tenminutus",
            "text": "推迟10分钟",
            "type": "button",
            "value": "sillyhat1",
            "style": ""
          },
          {
            "id": "2",
            "name": "halfhour",
            "text": "推迟半小时",
            "type": "button",
            "value": "sillyhat2",
            "style": ""
          },
          {
            "id": "3",
            "name": "onehour",
            "text": "推迟1小时",
            "type": "button",
            "value": "sillyhat3",
            "style": ""
          },
          {
            "id": "4",
            "name": "tomorrow",
            "text": "明天提醒我",
            "type": "button",
            "value": "sillyhat4",
            "style": ""
          }
        ]
      }
    ]
  },

*/
