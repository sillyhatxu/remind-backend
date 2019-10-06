package dao

import (
	"github.com/sillyhatxu/environment-config"
	"github.com/sillyhatxu/mysql-client"
)

var client *dbclient.MysqlClient

func InitialMysqlClient() error {
	mysqlClient, err := dbclient.NewMysqlClient(
		envconfig.Conf.DBRemindUserName,
		envconfig.Conf.DBRemindPassword,
		envconfig.Conf.DBRemindHost,
		envconfig.Conf.DBRemindPort,
		envconfig.Conf.DBRemindSchema,
		dbclient.DDLPath(envconfig.Conf.DBDDLPath),
		dbclient.Flyway(envconfig.Conf.DBFlyway),
	)
	client = mysqlClient
	return err
}
