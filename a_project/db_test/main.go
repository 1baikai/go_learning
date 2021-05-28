package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	opengauss  *gorm.DB
	Connection *gorm.DB
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		logrus.Error("err is ", err)
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(7 * time.Hour)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	Connection = gormDB
	opengauss = gormDB

	if err = autoMigrateTable(); err != nil {
		logrus.Errorf(fmt.Sprintf("autoMigrated  error  is %v", err))
		return
	}
	// 添加信息
	//Connection.Create(&Cluster{
	//	ID:            411111111111,
	//	ClusterName:   "bbb4",
	//	ConfigMethod:  "aaaa24",
	//	AddWay:        "aaa24",
	//	ClusterStatus: 2,
	//})
	//Connection.Create(&ZoneInfo{
	//	ID:           4111111111112,
	//	ClusterID:    411111111111,
	//	ZoneName:     "cccc3",
	//	VIP:          "192.168.1.31114",
	//	PingList:     "192.168.1.3114",
	//	Arping:       "192.168.1.314",
	//})
	//Connection.Create(&HostInfo{
	//	ID:           411111111113,
	//	ZoneID:       4111111111112,
	//	HostIP:       "192.168.1.58134",
	//	HostPort:     "236000",
	//	HostUsername: "omm32",
	//	HeartbeatIP:  "",
	//})
	//Connection.Create(&DatabaseInfo{
	//	ID:                 411111111114,
	//	ClusterID:          411111111111,
	//	ZoneID:             4111111111112,
	//	HostID:             411111111113,
	//	UserID:             411111111116,
	//	Name:               "aaa4",
	//	DatabaseName:       "postgres",
	//	Tag:                nil,
	//	Host:               "192.168.1.5843",
	//	Port:               "260003",
	//	Username:           "monit23",
	//	Password:           "asd32",
	//	DbAdminUser:        "admin23",
	//	DbAdminPassword:    "asd23",
	//	Status:             true,
	//	Version:            "2.0.0",
	//	Directory:          "c/c/c/",
	//	Role:               "priamry",
	//	UpstreamLibrary:    "",
	//	SysDatabaseID:      "12312323",
	//	SyncState:          "",
	//	TurnOnMonitor:      true,
	//	IsCluster:          true,
	//	InstallHAStatus:    2,
	//	AgentPort:          "",
	//	NodeStatus:         2,
	//	IsMogSystemInstall: true,
	//})
	//res := []*DatabaseInfo{}
	//err = Connection.Table("mog_database_info").Joins("left join mog_cluster_info on mog_database_info.cluster_id = mog_cluster_info.id where mog_cluster_info.cluster_status  in ? and mog_database_info.turn_on_monitor = ? and mog_database_info.node_status in ?", []int{2, 0}, true, []int{0, 2}).Scan(&res).Error
	//if err != nil {
	//	return
	//}
	//logrus.Printf("%v", res[0].ID)
	ceshi(true, 1, 10, "", "creationTime")
	ceshi3()
	ceshi2(411111111114)
}

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s connect_timeout=%s session_timeout=0",
		"127.0.0.1", "enmotech", "Enmo@123", "54321", "baikai", "disable", "5")
	//dsn := GetDSN(host, port, username, password, dbName)
	db, err := sql.Open("opengauss", dsn)
	if err != nil {
		return nil, err
	}

	return db, err
}

func autoMigrateTable() error {

	err := opengauss.AutoMigrate(
		new(Cluster),
		new(ZoneInfo),
		new(HostInfo),
		new(DatabaseInfo),
	)
	if err != nil {
		logrus.Errorf("AutoMigrate error %v", err)
		return err
	}
	logrus.Info("all tables autoMigrated.")
	return nil
}

func ceshi3() {
	res := Connection.Model(&DatabaseInfo{}).Where("id = ?", 411111111114).Updates(DatabaseInfo{

		Name:         "dbInfoName",
		DatabaseName: "mogdb",
		Host:         "192.168.1.58",
		Port:         "26000",
		Username:    "mogdb",
		Password:     "qweca",
	})
	if res.Error != nil {
		logrus.Errorf("err%v", res.Error)
		return
	}
}




func ceshi2(id int64) {
	dbInfo := DatabaseInfo{}
	res :=Connection.First(&dbInfo, id)
	if res.Error != nil {
		return
	}
	logrus.Info(dbInfo)

}


func ceshi(Desc bool, PageNum int, PageLimit int, Search, Order string) {
	var desc string
	if Desc {
		desc = "desc"
	} else {
		desc = "asc"
	}
	databaseInfos := []*DatabaseInfo{}
	offset := (PageNum - 1) * PageLimit
	search := fmt.Sprintf("%%%s%%", Search)
	logrus.Println("search", search)

	sqlRow := fmt.Sprintf("SELECT * FROM mog_database_info WHERE is_cluster is FALSE AND (name Like ? OR database_name LIKE ? OR host LIKE ? OR port LIKE ? OR version LIKE ? OR directory LIKE ? OR sys_database_id LIKE ? OR role LIKE ? OR upstream_library LIKE ?) ORDER BY %s %s LIMIT %d OFFSET %d;\n", fmt.Sprintf("\"%s\"", Camel2Case(Order)), desc, PageLimit, offset)
	New := strings.Replace(sqlRow, "?", "'"+search+"'", -1)
	result := Connection.Raw(New).Scan(&databaseInfos)
	if result.Error != nil {
		logrus.Println(result.Error)
		return
	}

	var total int
	sqlRow = "SELECT COUNT(*) FROM mog_database_info WHERE is_cluster is FALSE AND (name Like ? OR database_name LIKE ? OR host LIKE ? OR port LIKE ? OR version LIKE ? OR directory LIKE ? OR sys_database_id LIKE ? OR role LIKE ? OR upstream_library LIKE ?) ;\n"
	New = strings.Replace(sqlRow, "?", "'"+search+"'", -1)
	totalData := Connection.Raw(New).Scan(&total)
	if totalData.Error != nil {
		return
	}
	logrus.Info(databaseInfos,total,"\n")

}
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}
