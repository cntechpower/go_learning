package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// @model
// @tag=mysql
type MysqlMeta struct {
	// @field
	// @description=server id
	// @example=server-udp1
	ServerId string `json:"server_id"`
	// @field
	// @description=mysql group id
	// @example=mysql-group1
	GroupId string `json:"group_id"`
	// @field
	// @description=mysql instance id
	// @example=mysql-i1
	MysqlId string `json:"id" gorm:"primary_key"`
	// @field
	// @description=mysql alias
	// @example=172.20.134.1:5690
	Alias string `json:"alias"`

	// @field
	// @description=mysql install tarball
	// @example=mysql-5.7.20-linux-glibc2.12-x86_64.tar.gz
	Tarball string `json:"tarball"`
	// @field
	// @description=mysql version
	// @example=5.7.20
	Version string `json:"version"`
	// @field
	// @description=mysql installation standard
	// @example=uguard_gr
	Standard string `json:"standard"`

	// @field
	// @description=mysql my.cnf path
	// @example=/opt/udb/mysql/etc/my.cnf
	MycnfPath string `json:"mycnf_path"`
	// @field
	// @description=mysql my.cnf
	// @example=[mysqld]\ngroup_replication_group_name = OOOOOOOO-OOOO-OOOO-OOOO-OOOOOOOOO
	Mycnf string `json:"mycnf"`

	// @field
	// @description=mysql backup directory
	// @example=/mysql/backup
	BackupDir string `json:"backup_dir"`
	// @field
	// @description=mysql base path
	// @example=/mysql/base
	BasePath string `json:"base_path"`
	// @field
	// @description=mysql binlog path
	// @example=mysql/log/binlog/mysql-bin
	BinlogPath string `json:"binlog_path"`
	// @field
	// @description=mysql data path
	// @example=mysql/data
	Datapath string `json:"data_path"`
	// @field
	// @description=mysql pid path
	// @example=mysql/data/mysqld.pid
	PidPath string `json:"pid_path"`
	// @field
	// @description=mysql redolog path
	// @example=mysql/log/redolog
	RedologPath string `json:"redolog_path"`
	// @field
	// @description=mysql relaylog path
	// @example=mysql/log/relaylog/mysql-relay
	RelaylogPath string `json:"relaylog_path"`
	// @field
	// @description=mysql socket path
	// @example=mysql/data/mysqld.sock
	SockPath string `json:"sock_path"`
	// @field
	// @description=mysql tmp path
	// @example=mysql/tmp
	TmpPath string `json:"tmp_path"`
	// @field
	// @description=mysql error log path
	// @example=mysql/data/mysql-error.log
	ErrorLogPath string `json:"errorlog_path"`
	// @field
	// @description=mysql serverId
	// @example=242464
	MysqlServerId string `json:"mysql_serverid"`
	// @field
	// @description=master sequence
	// @example=100
	// @default=100
	MasterSequence string `json:"master_sequence"`
	// @field
	// @description=mysql run user
	// @example=mysql
	RunUser string `json:"run_user"`
	// @field
	// @description=mysql run user uid
	// @example=0
	Uid string `json:"uid"`
	// @field
	// @description=mysql run user group
	// @example=mysql-users
	RunUserGroup string `json:"run_user_group"`
	// @field
	// @description=mysql run user group gid
	// @example=0
	Gid string `json:"gid"`
	// @field
	// @description=mysql port
	// @example=5690
	Port string `json:"port"`
	// @field
	// @description=if mysql is readonly
	// @default=false
	ReadOnly bool `json:"readonly"`

	// @field
	// @description=mgr port
	// @example=5691
	GrPort string `json:"gr_port"`

	// @field
	// @description=quota limit
	// @example=0
	QuotaLimit int `json:"quota_limit"`
	// @field
	// @description=cpu quota percent
	// @example=200%
	CpuQuota int `json:"cpu_quota"`
	// @field
	// @description=mem limit
	// @example=0
	MemLimit int `json:"mem_limit"`
	// @field
	// @description=iops
	// @example=0
	Iops int `json:"iops"`

	// @field
	// @description=mysql scsi device
	ScsiDev string `json:"scsi_dev"`
	// @field
	// @description=mysql scsi mount point
	ScsiMountPoint string `json:"scsi_mount_point"`
	// @field
	// @description=mysql scsi readonly mount point
	ScsiReadOnlyMountPoint string `json:"scsi_read_only_mount_point"`
	// @field
	// @description=mysql scsi primary key
	ScsiPrKey string `json:"scsi_pr_key"`

	// @field
	// @description=if ha is standard
	// @default=false
	LooseIoFlushParamCheck bool `json:"loose_io_flush_param_check"`
	// @field
	// @description=if mysql should be monitored
	// @default=true
	MonitorEnable bool `json:"monitor_enable"`
	// @field
	// @description=if should watch slow log
	// @default=true
	SlowLogWatchEnable bool `json:"slow_log_watch_enable"`
	// @field
	// @description=ustats logs
	UstatsLogs string `json:"ustats_logs"`
	// @field
	// @description=mysql udelay enable
	// @example=true
	UdelayEnable bool `json:"udelay_enable"`
	// @field
	// @description=instance's useable sip
	// @example=172.14.20.111
	UseableSip string `json:"useable_sip"`
	// @field
	// @description=network segment of instance's useable sip
	// @example=172.14.20.0/24
	UseableSipNetSegment string `json:"useable_sip_net_segment"`
	// @field
	// @description=the ip that used for replication(if empty.use server ip instead)
	// @example=172.14.20.111
	ReplIp string `json:"repl_ip"`
	// @field
	// @description=instance's sip network device
	// @example=eth0
	SipDevice string `json:"sip_device"`
	// @field
	// @description=instance's snapshot task configs
	SnapshotTasks []*SnapshotTask `json:"snapshot_tasks" gorm:"foreignkey:MysqlId"`
}

type SnapshotTask struct {
	// @field
	// @description=tag id
	// @example=AUTO_GENERATED
	TaskId uint `json:"-" gorm:"primary_key"`
	// @field
	// @description=id , used as foreign key
	// @example=mysql-test
	MysqlId string `json:"-"`
	// @field
	// @description=snapshot task id
	// @example=snapshot-config-1
	SnapshotTaskId string `json:"snapshot_task_id"`
	// @field
	// @description=will track changes to this file, and then compare changes to keyword
	// @example=/opt/mysql/data/3306/mysql-error.log
	TrackFilePath string `json:"track_file_path"`
	// @field
	// @description=if found keyword in tracked file changes, then will generate snapshot with given bash script
	// @example=应用英文名
	TrackKeyword string `json:"track_keyword"`
	// @field
	// @description=bash script used to generate a snapshot
	// @example=bash script content
	BashScriptContent string `json:"bash_script_content"`
}

func main() {
	db, err := gorm.Open("sqlite3", "/tmp/umc.db")
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}
	strings.HasSuffix()
	defer db.Close()
	if err := db.AutoMigrate(&MysqlMeta{}, &SnapshotTask{}).Error; err != nil {
		log.Fatalf("migrate error: %v", err)
	}
	t1 := []*SnapshotTask{{
		MysqlId:           "mysql-1",
		SnapshotTaskId:    "task-1",
		TrackFilePath:     "",
		TrackKeyword:      "",
		BashScriptContent: "",
	}}
	g1 := &MysqlMeta{
		MysqlId:       "mysql-1",
		SnapshotTasks: t1,
	}
	db.Save(g1)
	g2 := &MysqlMeta{MysqlId: "mysql-1"}
	if err := db.Preload("SnapshotTasks").Find(g2).Error; err != nil {
		log.Fatalf("find error: %v", err)
	}
	fmt.Println(g2.SnapshotTasks[0])
}
