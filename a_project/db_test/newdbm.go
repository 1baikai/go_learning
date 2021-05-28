package main

import (
	"database/sql/driver"
	"encoding/json"
)

type Cluster struct {
	ID            int64  `json:"id" gorm:"column:id"`
	ClusterName   string `json:"clusterName" gorm:"column:cluster_name"`
	ConfigMethod  string `json:"configMethod" gorm:"column:config_method"`
	AddWay        string `json:"addWay" gorm:"column:add_way"`
	ClusterStatus int    `json:"clusterStatus" gorm:"column:cluster_status"`                     // 0 安装失败--，1 安装中，2 安装成功 运行中，3 集群不可用 ，4有故障,5.停止    集群安装状态
	CreationTime  int64  `json:"creationTime" gorm:"column:creation_time;autoCreateTime:milli"` //creation_time    --
	UpdateTime    int64  `json:"updateTime" gorm:"column:update_time;autoUpdateTime:milli"`     //update_time      --
}

func (Cluster) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "mog_cluster_info"
}
type ZoneInfo struct {
	ID           int64  `json:"id" gorm:"column:id"`
	ClusterID    int64  `json:"cluster_id" gorm:"cluster_id"`
	ZoneName     string `json:"zoneName" gorm:"column:zone_name"`
	VIP          string `json:"vIP" gorm:"column:vip"`
	PingList     string `json:"pingList" gorm:"column:ping_list"`
	Arping       string `json:"arping" gorm:"column:arping"`
	CreationTime int64  `json:"creationTime" gorm:"column:creation_time;autoCreateTime:milli"` //creation_time    --
	UpdateTime   int64  `json:"updateTime" gorm:"column:update_time;autoUpdateTime:milli"`     //update_time      --
}

func (ZoneInfo) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "mog_zone_info"
}


type HostInfo struct {
	ID           int64  `json:"id" gorm:"column:id"`
	ZoneID       int64  `json:"zone_id" gorm:"zone_id"`
	HostIP       string `json:"host_ip" gorm:"host_ip"`   //
	HostPort     string `json:"host_port" gorm:"host_port"`  //22
	HostUsername string `json:"host_username" gorm:"host_username"` //omm
	HeartbeatIP  string `json:"heartbeatIP" gorm:"column:heartbeat_ip"`
	//DockerDBHost       string  `json:"dockerDBHost" gorm:"column:dockerDBHost"`     // 相当于主机ip
}

func (HostInfo) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "mog_host_info"
}


type DatabaseInfo struct {
	ID                 int64   `json:"id" gorm:"column:id"`
	ClusterID          int64   `json:"cluster_id" gorm:"cluster_id"`
	ZoneID             int64   `json:"zone_id" gorm:"zone_id"`
	HostID             int64   `json:"host_id" gorm:"host_id"`
	UserID             int64   `json:"userID" gorm:"column:user_id"`
	Name               string  `json:"name" gorm:"column:name"`
	DatabaseName       string  `json:"databaseName" gorm:"column:database_name"`
	Tag                TagList `json:"tag" gorm:"column:tag"`
	Host               string  `json:"host" gorm:"column:host;uniqueIndex:ip"`
	Port               string  `json:"port" gorm:"column:port;uniqueIndex:ip"`
	Username           string  `json:"username" gorm:"column:username"`
	Password           string  `json:"password" gorm:"column:password"`
	DbAdminUser        string  `json:"dbAdminUser" gorm:"column:db_admin_user"`
	DbAdminPassword    string  `json:"dbAdminPassword" gorm:"column:db_admin_password"`
	Status             bool    `json:"status" gorm:"column:status"`
	Version            string  `json:"version" gorm:"column:version"`
	Directory          string  `json:"directory" gorm:"column:directory"`
	Role               string  `json:"role" gorm:"column:role"`
	UpstreamLibrary    string  `json:"upstreamLibrary" gorm:"column:upstream_library"`
	SysDatabaseID      string  `json:"sysDatabaseID" gorm:"column:sys_database_id"`
	CreationTime       int64   `json:"creationTime" gorm:"column:creation_time;autoCreateTime:milli"` //creation_time    --
	UpdateTime         int64   `json:"updateTime" gorm:"column:update_time;autoUpdateTime:milli"`     //update_time      --
	SyncState          string  `json:"syncState" gorm:"column:sync_state"`
	TurnOnMonitor      bool    `json:"turnOnMonitor" gorm:"column:turn_on_monitor"`      //是否开启监控
	IsCluster          bool    `json:"isCluster" gorm:"column:is_cluster"`               //是否是集群
	InstallHAStatus    int     `json:"installHAStatus" gorm:"column:install_ha_status" ` // 安装ha状态】// 0 --，1 安装中，2 安装成功，3 安装失败
	AgentPort          string  `json:"agentPort" gorm:"column:agent_port"`
	NodeStatus         int     `json:"nodeStatus" gorm:"column:node_status"`
	IsMogSystemInstall bool    `json:"isMogSystemInstall" gorm:"column:is_mog_system_install"`
}
func (DatabaseInfo) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "mog_database_info"
}



type TagList []int64


// gorm 自定义结构需要实现 Value Scan 两个方法
// Value 实现方法
func (t TagList) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// Scan 实现方法
func (t *TagList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &t)
}

// TagInfo 标签
type TagInfo struct {
	ID           int64  `json:"id" gorm:"column:id"`
	Key          string `json:"key" gorm:"column:tagKey;uniqueIndex:kv"`
	Value        string `json:"value" gorm:"column:tagValue;uniqueIndex:kv"`
	CreationTime int64  `json:"creationTime" gorm:"column:creationTime;autoCreateTime:milli"` //creation_time    --
	UpdateTime   int64  `json:"updateTime" gorm:"column:updateTime;autoUpdateTime:milli"`     //update_time      --

}


func (TagInfo) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "mog_tag"
}


// InstallDBLog 日志
type InstallDBLog struct {
	ID           int64  `json:"id" gorm:"column:id"`
	ClusterID    int64  `json:"clusterID" gorm:"column:clusterID"`
	Label        string `json:"label" gorm:"column:label"`
	Data         string `json:"data" gorm:"column:data"`
	CreationTime int64  `json:"creationTime" gorm:"column:creationTime;autoCreateTime:milli"` //creation_time
}

func (InstallDBLog) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "mog_install_database_log"
}
