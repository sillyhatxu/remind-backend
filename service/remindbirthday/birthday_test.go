package remindbirthday

import (
	"fmt"
	"github.com/sillyhatxu/environment-config"
	"github.com/sillyhatxu/remind-backend/dao"
	"github.com/sillyhatxu/remind-backend/dto"
	"github.com/sillyhatxu/remind-backend/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	path := "/Users/shikuanxu/go/src/github.com/sillyhatxu/remind-backend"
	err := envconfig.LoadEnv(fmt.Sprintf("%s/%s", path, ".env"))
	if err != nil {
		panic(err)
	}
	err = envconfig.ParseDefaultConfig(envconfig.ConfigFile(fmt.Sprintf("%s/%s", path, "config-local.conf")))
	if err != nil {
		panic(err)
	}
	err = utils.InitialIdGenerator()
	if err != nil {
		panic(err)
	}
	err = dao.InitialMysqlClient()
	if err != nil {
		panic(err)
	}
	err = InitialBirthdayClient()
	if err != nil {
		panic(err)
	}
}

func TestSendBirthdayRemind(t *testing.T) {
	err := SendBirthdayRemind(dto.BirthdayRemindDTO{
		Id:         utils.GeneratorId(),
		Fallback:   "Happy Birthday : Shikuan Xu",
		Text:       "生日快乐 阴历 : 五月十四",
		AuthorName: "徐士宽",
	})
	assert.Nil(t, err)
}
