package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	//redigo "github.com/gomodule/redigo/redis"
)

var infos1 = `
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:8
cluster_my_epoch:1
cluster_stats_messages_sent:70406
cluster_stats_messages_received:69203
`

var infos = `
# Server
#redis_version:3.2.12
#redis_git_sha1:00000000
#redis_git_dirty:0
#redis_build_id:f8ee28ac3c8d6d36
#redis_mode:standalone
#os:Linux 3.10.0-862.14.4.el7.x86_64 x86_64
#arch_bits:64
#multiplexing_api:epoll
#gcc_version:4.8.5
#process_id:5625
#run_id:fa324cc9a54dd7e1cea793534a7768808447e9ed
#tcp_port:6379
uptime_in_seconds:248154
#uptime_in_days:2
#hz:10
#lru_clock:11960492
#executable:/data/redis3/base/redis-server
#config_file:/data/redis3/redis.conf

# Clients
connected_clients:1
client_longest_output_list:0
client_biggest_input_buf:0
blocked_clients:0

# Memory
used_memory:821760
#used_memory_human:802.50K
used_memory_rss:7278592
#used_memory_rss_human:6.94M
used_memory_peak:3983224
#used_memory_peak_human:3.80M
total_system_memory:3973763072
#total_system_memory_human:3.70G
used_memory_lua:37888
#used_memory_lua_human:37.00K
maxmemory:0
#maxmemory_human:0B
#maxmemory_policy:noeviction
mem_fragmentation_ratio:8.86
#mem_allocator:jemalloc-4.0.3

# Persistence
loading:0
rdb_changes_since_last_save:1
rdb_bgsave_in_progress:0
rdb_last_save_time:1571993426
rdb_last_bgsave_status:err
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
aof_enabled:0
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok

# Stats
total_connections_received:905
total_commands_processed:404130
instantaneous_ops_per_sec:1
total_net_input_bytes:103680849
total_net_output_bytes:450654837
instantaneous_input_kbps:0.03
instantaneous_output_kbps:3.07
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:0
evicted_keys:0
keyspace_hits:0
keyspace_misses:404088
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:282
migrate_cached_sockets:0

# Replication
role:master
connected_slaves:0
master_repl_offset:0
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0

# CPU
used_cpu_sys:187.48
used_cpu_user:98.06
used_cpu_sys_children:63.73
used_cpu_user_children:9.10

# Commandstats
cmdstat_get:calls=100000,usec=26357,usec_per_call=0.26
cmdstat_set:calls=1,usec=45,usec_per_call=45.00
cmdstat_lrange:calls=304088,usec=95596,usec_per_call=0.31
cmdstat_info:calls=36,usec=2771,usec_per_call=76.97
cmdstat_command:calls=5,usec=2344,usec_per_call=468.80

# Cluster
cluster_enabled:0

# Keyspace
db0:keys=1,expires=0,avg_ttl=0

`

func infoToMapParser(info string) (map[string]string, error) {
	infoReader := bufio.NewReader(bytes.NewReader([]byte(info)))
	infosMap := make(map[string]string, 0)
	var callTotal float64
	var usecTotal float64
	var keyTotal float64
	var keysExpired float64
	for {
		line, _, err := infoReader.ReadLine()
		if err != nil {
			break
		}
		stringLine := string(line)
		//skip empty line and comment line
		if strings.HasPrefix(stringLine, "#") || 0 == len(stringLine) || stringLine == "\r\n" {
			continue
		}
		kv := strings.Split(stringLine, ":")
		if len(kv) != 2 {
			continue
		}
		key, value := kv[0], kv[1]
		//special for cmdstat_eval:calls=10000,usec=35221,usec_per_call=3.52
		if strings.HasPrefix(key, "cmdstat_") {

			values := strings.Split(value, ",")
			for _, value := range values {
				kv1 := strings.Split(value, "=")
				if len(kv1) != 2 {
					continue
				}
				if kv1[0] == "calls" {
					num, err := strconv.ParseFloat(kv1[1], 64)
					if err != nil {
						continue
					}
					callTotal = callTotal + num
				}
				if kv1[0] == "usec" {
					num, err := strconv.ParseFloat(kv1[1], 64)
					if err != nil {
						continue
					}
					usecTotal = usecTotal + num
				}
			}
			infosMap["commands_total"] = strconv.FormatFloat(callTotal, 'f', 1, 64)
			infosMap["commands_duration_seconds_total"] = strconv.FormatFloat(usecTotal, 'f', 1, 64)
			continue
		}
		//special for db0:keys=1580112,expires=0,avg_ttl=0`
		if strings.HasPrefix(key, "db") {

			values := strings.Split(value, ",")
			for _, value := range values {
				kv1 := strings.Split(value, "=")
				if len(kv1) != 2 {
					continue
				}
				if kv1[0] == "keys" {
					num, err := strconv.ParseFloat(kv1[1], 64)
					if err != nil {
						continue
					}
					keyTotal = keyTotal + num
				}
				if kv1[0] == "expires" {
					num, err := strconv.ParseFloat(kv1[1], 64)
					if err != nil {
						continue
					}
					keysExpired = keysExpired + num
				}
			}
			infosMap["db_keys_total"] = strconv.FormatFloat(keyTotal, 'f', 1, 64)
			infosMap["db_keys_expired"] = strconv.FormatFloat(keysExpired, 'f', 1, 64)
			continue
		}
		infosMap[key] = value
	}
	return infosMap, nil
}

func main() {
	//conn, err := redigo.Dial("tcp", "10.186.64.16:6379")
	//if err != nil {
	//	panic(err)
	//}
	//reply, err := redigo.String(conn.Do("INFO", "ALL"))
	//if err != nil {
	//	panic(err)
	//}
	//conn.Close()
	infoMap, err := infoToMapParser(infos1)
	if err != nil {
		panic(err)
	}
	for k, _ := range infoMap {
		fmt.Printf(`
	"redis:info_all:%v": {
	"tags": {},
	"enable": true,
	"RequiredByUguard": false
	},`, k)
	}
	//fmt.Println(infoMap)

}
