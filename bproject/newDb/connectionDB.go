package newDb

import (
	"database/sql"
	"fmt"

	"bproject/entity"
	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s connect_timeout=%s session_timeout=0",
		"127.0.0.1", "enmotech", "Enmo@123", "54321", "enmotech", "disable", "5")
	//dsn := GetDSN(host, port, username, password, dbName)
	db, err := sql.Open("opengauss", dsn)
	if err != nil {
		return nil, err
	}

	return db, err
}

var ConnectDatabase, dbSql *gorm.DB

func InitDB() error {

	db, err := ConnectDB()
	if err != nil {
		return errors.Wrap(err, "数据库连接失败")
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "初始化grom 失败")
	}

	dbSql = gormDB
	ConnectDatabase = gormDB

	if err = autoMigrateTable(); err != nil {
		logrus.Errorf(fmt.Sprintf("autoMigrated  error  is %v", err))
		return err
	}
	return nil
}

func autoMigrateTable() error {
	err := dbSql.AutoMigrate(
		new(entity.User),
	)
	if err != nil {
		return errors.Wrap(err, "初始化表错误")
	}
	return nil
}
