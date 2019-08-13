package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ComponentGroupMeta struct {
	// @field
	// @description=component group id
	// @example=uproxy-g1
	GroupId string `json:"group_id" gorm:"primary_key"`
	// @field
	// @description=component group type
	// @example=uproxy
	Type string `json:"type"`
	// @field
	// @description=component group admin user
	// @example=admin
	AdminUser string `json:"admin_user"`
	// @field
	// @description=component group password
	// @example=aeN8MV7d2Ae6x1
	AdminPassword string `json:"admin_password"`
	// @field
	// @description=component group port(ushard, uproxy)
	// @example=13665
	Port string `json:"port"`
	// @field
	// @description=if component should be monitored
	// @default=true
	MonitorEnable bool `json:"monitor_enable"`
	// @field
	// @description=sip of component group
	// @example=172.30.10.163
	Sip string `json:"sip"`
	// @field
	// @description=sip of component group
	// @example=172.30.10.163
	SipNetSegment string `json:"sip_net_segment"`
	// @field
	// @description=if auto commit (uproxy)
	IsAutoCommit bool `json:"is_auto_commit"`
	// @field
	// @description=ignore update route error
	IgnoreUpdateRouteError bool `json:"ignore_update_route_error"`
	// @field
	// @description=ushard params
	// @example={"charset":"utf8","fakeMySQLVersion":"5.6.23","managerPort":"13310","serverPort":"13306"}
	UshardParams string `json:"ushard_params"`
	// @field
	// @description=ushard xml config files
	// @example={"cacheservice.properties":"#\\n# Copyright (C) 2016-2019 ActionTech.\\n# License: http://www.gnu.org/licenses/gpl.html...
	UshardXmlConfigFiles string `json:"ushard_xml_config_files"`
	// @field
	// @description=ushard json config
	// @example={}
	UshardJsonConfig string `json:"ushard_json_config"`
	// @field
	// @description=ushard reload error
	// @example=err
	UshardLastReloadError string `json:"ushard_last_reload_error"`
	// @field
	// @description=ushard used-work-ids
	// @example=0,1,2,3
	UshardUsedWorkIds string `json:"ushard_used_work_ids"`
}

type ComponentMeta struct {
	// @field
	// @description=server id, should not be nil
	// @example=server-udp1
	ServerId string `json:"server_id"`
	// @field
	// @description=component group id
	// @example=ushard-g1
	GroupId string `json:"component_group_id"`
	// @field
	// @description=component id
	// @example=ushard-i1
	CompId string `json:"component_id" gorm:"primary_key"`
	// @field
	// @description=component port
	// @example=7776
	Port string `json:"port"`
	// @field
	// @description=component pgrpc port
	// @example=7775
	PgrpcPort string `json:"pgrpc_port"`

	// @field
	// @description=if component should be monitored
	// @default=true
	MonitorEnable bool `json:"monitor_enable"`
	// @field
	// @description=component pid file
	// @example=/opt/comp/comp.pid
	PidFile string `json:"pid_file"`
	// @field
	// @description=component work directory
	// @example=/opt/comp
	WorkDir string `json:"work_dir"`
	// @field
	// @description=component install directory
	// @example=/opt/comp
	InstallDir string `json:"install_dir"`
	// @field
	// @description=component type, lowercase required
	// @example=uagent
	Type string `json:"type"`
	// @field
	// @description=ustats logs
	// @example=/opt/ustats/logs/brief.log
	UstatsLogs string `json:"ustats_logs"`
	// @field
	// @description=component startup command
	// @example=nohup /opt/ustats/bin/ustats 1>/opt/ustats/logs/std.log 2>&1 &
	StartCmd string `json:"start_cmd"`
	// @field
	// @description=watched by uagent
	// @default=false
	UagentWatch bool `json:"uagent_watch"`
	// @field
	// @description=guarded by uagent
	// @example=false
	UagentGuard bool `json:"uagent_guard"`
	// @field
	// @description=ucore heartbeat interval millisecond
	// @example=5000
	UcoreHeartbeatInterval int `json:"ucore_heartbeat_interval"`

	// @field
	// @description=number of symmertric multiprocessing (from uproxy)
	// @example=1
	Smp int `json:"smp"`
	// @field
	// @description=ushard params
	// @example={"dble_install_dir":"/opt/ushard","dble_version":"\\"9.9.9.9 104bf8f4ca5f674375638bcbf2ef0f465b0726fa\\"","direct_memory_max_capacity":"1024","memory_heap_initial_capacity":"1024","memory_heap_max_capacity":"4096","processors":"1","useOffHeapForMerge":"0","log4j-level":"INFO"}
	UshardParams string `json:"ushard_params"`

	// @field
	// @description=uterm proxy-ip
	// @example=172.30.10.157
	ProxyIp string `json:"proxy_ip"`
	// @field
	// @description=uterm proxy-port
	// @example=15351
	ProxyPort int `json:"proxy_port"`

	// @field
	// @description=log file limit, the unit is M
	// @example=100
	// @default=100
	LogFileLimit int `json:"log_file_limit"`
	// @field
	// @description=log files total_limit
	// @example=1024
	// @default=1024
	LogTotalLimit int `json:"log_total_limit"`
	// @field
	// @description=component tmp dir, will be applied to uguard-agent and urman-agent when it is started.
	// @example="/tmp"
	// @default=""
	TmpDir string `json:"tmp_dir"`
	// @field
	// @description=component sip device
	// @example=eth0
	SipDevice string `json:"sip_device"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "/tmp/umc.db")
	defer db.Close()
	// group := new(ComponentGroupMeta)
	// db.Joins("LEFT JOIN tag_meta on component_group_meta.group_id=tag_meta.id").
	// 	Where("type='ushard'").
	// 	First(&group)
	// json, _ := json.Marshal(group)
	// fmt.Println(string(json))
	fuzzyKeyword := "dd"
	fuzzyMatchCondition := fmt.Sprintf("component_meta.server_id like '%%%[1]v%%' "+
		"or component_meta.group_id like '%%%[1]v%%'"+
		"or component_meta.comp_id like '%%%[1]v%%'"+
		"or component_meta.port like '%%%[1]v%%'"+
		"or umc_component_statuses.status_desc like '%%%[1]v%%'"+
		"or tag_meta.attribute like '%%%[1]v%%'"+
		"or tag_meta.value like '%%%[1]v%%'", fuzzyKeyword)
	inst := new(ComponentMeta)
	db.Joins("JOIN umc_component_statuses ON umc_component_statuses.id=component_meta.comp_id LEFT JOIN tag_meta on component_meta.comp_id=tag_meta.id").
		Where(fuzzyMatchCondition).
		First(&inst)
	json, _ := json.Marshal(inst)
	fmt.Println(string(json))

}
