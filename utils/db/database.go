package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	UTC   = "UTC"
	LOCAL = "Local"
)

type PQDBInfo struct {
	SchemaName          string
	SingularTable       bool
	TimeZone            string
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBConnectionTimeout string
	DBSearchPath        string
	DBMaxOpenConn       string
	DBMaxIdleConn       string
}

func GetSqlClient(pqdbinfo *PQDBInfo, logger *logrus.Entry) *sql.DB {
	dbTimeout, err := strconv.Atoi(pqdbinfo.DBConnectionTimeout)
	if err != nil {
		logger.WithField("err", err).WithField("value", pqdbinfo.DBConnectionTimeout+"d").Error("Error in converting dbConnTimeout to integer")
		panic(err)
	}
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d search_path=%s",
		pqdbinfo.DBHost, pqdbinfo.DBPort, pqdbinfo.DBUser, pqdbinfo.DBPassword, pqdbinfo.DBName, dbTimeout, pqdbinfo.DBSearchPath)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		logger.WithField("err", err).Error("Error from sq.Open")
		panic(err)
	}
	i, err := strconv.Atoi(pqdbinfo.DBMaxOpenConn)
	if err != nil {
		logger.WithField("err", err).Error("Error in converting dbMaxOpenConns to integer")
		panic(err)
	}
	j, err := strconv.Atoi(pqdbinfo.DBMaxIdleConn)
	if err != nil {
		logger.WithField("err", err).Error("Error in converting dbMaxIdleConns to integer")
		panic(err)
	}

	db.SetMaxOpenConns(i)
	db.SetMaxIdleConns(j)
	logger.Info("Setting maxOpenConns to", i, "and maxIdleConns to", j)
	return db
}

func GetGormSqlClient(pqdbinfo *PQDBInfo, logger *logrus.Entry) *gorm.DB {
	dbTimeout, err := strconv.Atoi(pqdbinfo.DBConnectionTimeout)
	if err != nil {
		logger.WithField("err", err).WithField("value", pqdbinfo.DBConnectionTimeout+"d").Error("Error in converting dbConnTimeout to integer")
		panic(err)
	}
	i, err := strconv.Atoi(pqdbinfo.DBMaxOpenConn)
	if err != nil {
		logger.WithField("err", err).Error("Error in converting dbMaxOpenConns to integer")
		panic(err)
	}
	j, err := strconv.Atoi(pqdbinfo.DBMaxIdleConn)
	if err != nil {
		logger.WithField("err", err).Error("Error in converting dbMaxIdleConns to integer")
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d search_path=%s",
		pqdbinfo.DBHost, pqdbinfo.DBPort, pqdbinfo.DBUser, pqdbinfo.DBPassword, pqdbinfo.DBName, dbTimeout, pqdbinfo.DBSearchPath)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   pqdbinfo.SchemaName,
			SingularTable: pqdbinfo.SingularTable,
		}, Logger: gLogger.Default.LogMode(gLogger.Silent),
		NowFunc: func() time.Time {
			tz, err := time.LoadLocation(pqdbinfo.TimeZone)
			if err != nil {
				logger.WithField("err", err).Error("error in getting the timezone")
				panic(err)
			}
			return time.Now().In(tz)
		},
	})
	if err != nil {
		logger.WithField("err", err).Error("Error from sq.Open")
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.WithField("err", err).Error("Error in getting sql db from gorm db")
		panic(err)
	}
	sqlDB.SetMaxOpenConns(i)
	sqlDB.SetMaxIdleConns(j)
	logger.Info("Set maxOpenConns to", i, "and maxIdleConns to", j)
	return db
}
