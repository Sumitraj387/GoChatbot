package providers

import (
	"strconv"

	"GoChatbot/utils/db"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetGormDbClient(config AppConfig, logger *logrus.Entry) (*gorm.DB, error) {

	dbInfo := db.PQDBInfo{
		DBHost:              config.DbConfig.Host,
		DBPort:              strconv.Itoa(config.DbConfig.Port),
		DBUser:              config.DbConfig.User,
		DBPassword:          config.DbConfig.Password,
		DBName:              config.DbConfig.Name,
		DBConnectionTimeout: strconv.Itoa(config.DbConfig.Timeout),
		DBSearchPath:        config.DbConfig.SearchPath,
		DBMaxOpenConn:       strconv.Itoa(config.DbConfig.MaxOpenConn),
		DBMaxIdleConn:       strconv.Itoa(config.DbConfig.MaxIdleConn),
		SchemaName:          "public.",
	}
	return db.GetGormSqlClient(&dbInfo, logger), nil
}
