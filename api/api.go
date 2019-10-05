package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/codes"
	"github.com/sillyhatxu/gin-utils/entity"
	"github.com/sillyhatxu/gin-utils/ginerrors"
	"github.com/sillyhatxu/gin-utils/interceptor/loginterceptor"
	"github.com/sillyhatxu/gin-utils/response"
	"github.com/sillyhatxu/remind-backend/dto"
	"github.com/sillyhatxu/remind-backend/service"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func InitialAPI(listener net.Listener) {
	logrus.Info("initial internal api start")
	router := SetupRouter()
	err := http.Serve(listener, router)
	if err != nil {
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "UP", "message": "OK"}) })
	remindGroup := router.Group("/slack/remind").Use(loginterceptor.Handle())
	{
		remindGroup.POST("", remind)
	}
	return router
}

func remind(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusOK, ginerrors.Error(codes.InvalidParameter, err.Error()))
		return
	}
	decodedValue, err := url.QueryUnescape(string(body))
	if err != nil {
		context.JSON(http.StatusOK, ginerrors.Error(codes.InvalidParameter, err.Error()))
		return
	}
	bodyArray := strings.Split(decodedValue, "=")
	if bodyArray == nil || len(bodyArray) != 2 {
		logrus.Errorf("body format error. %v", string(body))
		context.JSON(http.StatusOK, ginerrors.Errorf(codes.InvalidParameter, "body format error", entity.Data(string(body))))
		return
	}
	var slackDTO dto.SlackDTO
	err = json.Unmarshal([]byte(bodyArray[1]), &slackDTO)
	if err != nil {
		context.JSON(http.StatusOK, ginerrors.Error(codes.InvalidParameter, err.Error()))
		return
	}
	err = service.Check(slackDTO)
	if err != nil {
		context.JSON(http.StatusOK, ginerrors.Convert(err))
		return
	}
	context.JSON(http.StatusOK, response.Success(codes.OK, entity.Msg("Operator Success")))
}
