package utils

import (
	"github.com/sillyhatxu/id-generator"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var client *idgenerator.GeneratorClient

func InitialIdGenerator() error {
	c, err := idgenerator.NewGeneratorClient("REMIND", idgenerator.SequenceFormat("%04d"))
	if err != nil {
		return err
	}
	client = c
	return nil
}

func GeneratorId() string {
	if client == nil {
		logrus.Error("id generator client nil.")
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	id, err := client.GeneratorId()
	if err != nil {
		logrus.Error("generator id error. %v", err)
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return id
}
