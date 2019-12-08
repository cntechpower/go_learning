package main

import (
	"fmt"
	"regexp"
)

//var info string = `
//bash: warning: setlocale: LC_ALL: cannot change locale (en_US.UTF-8)\nbash: warning: setlocale: LC_ALL: cannot change locale (en_US.UTF-8)\n\nprotected-mode yes\ntcp-backlog 511\ntimeout 0\ntcp-keepalive 300\ndaemonize yes\nsupervised systemd\ndatabases 16\nsave 900 1\nsave 300 10\nsave 60 10000\nstop-writes-on-bgsave-error yes\nrdbcompression yes\nrdbchecksum yes\ndbfilename \"dump.rdb\"\nslave-serve-stale-data yes\nslave-read-only yes\nrepl-diskless-sync no\nrepl-diskless-sync-delay 5\nrepl-disable-tcp-nodelay no\nslave-priority 100\nappendonly no\nappendfilename \"appendonly.aof\"\nappendfsync everysec\nno-appendfsync-on-rewrite no\nauto-aof-rewrite-percentage 100\nauto-aof-rewrite-min-size 64mb\naof-load-truncated yes\nlua-time-limit 5000\nslowlog-log-slower-than 10000\nslowlog-max-len 128\nlatency-monitor-threshold 0\nnotify-keyspace-events \"\"\nhash-max-ziplist-entries 512\nhash-max-ziplist-value 64\nlist-max-ziplist-size -2\nlist-compress-depth 0\nset-max-intset-entries 512\nzset-max-ziplist-entries 128\nzset-max-ziplist-value 64\nhll-sparse-max-bytes 3000\nactiverehashing yes\nclient-output-buffer-limit normal 0 0 0\nclient-output-buffer-limit slave 256mb 64mb 60\nclient-output-buffer-limit pubsub 32mb 8mb 60\nhz 10\naof-rewrite-incremental-fsync yes\nport 3333\nbind *\nmaxmemory 536870912\ndir /opt/redis-3.2.12/data\npidfile /opt/redis-3.2.12/data/redis.pid\nlogfile /opt/redis-3.2.12/data/redis.log\nrequirepass 1
//`

var info string = `bash: warning: setlocale: LC_ALL: cannot change locale (en_US.UTF-8)\nbash: warning: setlocale: LC_ALL: cannot change locale (en_US.UTF-8)\n\nprotected-mode yes\ntcp-backlog 511\ntimeout 0\ntcp-keepalive 300\ndaemonize yes\nsupervised systemd\ndatabases 16\nsave 900 1\nsave 300 10\nsave 60 10000\nstop-writes-on-bgsave-error yes\nrdbcompression yes\nrdbchecksum yes\ndbfilename \"dump.rdb\"\nslave-serve-stale-data yes\nslave-read-only yes\nrepl-diskless-sync no\nrepl-diskless-sync-delay 5\nrepl-disable-tcp-nodelay no\nslave-priority 100\nappendonly no\nappendfilename \"appendonly.aof\"\nappendfsync everysec\nno-appendfsync-on-rewrite no\nauto-aof-rewrite-percentage 100\nauto-aof-rewrite-min-size 64mb\naof-load-truncated yes\nlua-time-limit 5000\nslowlog-log-slower-than 10000\nslowlog-max-len 128\nlatency-monitor-threshold 0\nnotify-keyspace-events \"\"\nhash-max-ziplist-entries 512\nhash-max-ziplist-value 64\nlist-max-ziplist-size -2\nlist-compress-depth 0\nset-max-intset-entries 512\nzset-max-ziplist-entries 128\nzset-max-ziplist-value 64\nhll-sparse-max-bytes 3000\nactiverehashing yes\nclient-output-buffer-limit normal 0 0 0\nclient-output-buffer-limit slave 256mb 64mb 60\nclient-output-buffer-limit pubsub 32mb 8mb 60\nhz 10\naof-rewrite-incremental-fsync yes\nport 3333\nbind *\nmaxmemory 536870912\ndir /opt/redis-3.2.12/data\npidfile /opt/redis-3.2.12/data/redis.pid\nlogfile /opt/redis-3.2.12/data/redis.log\nrequirepass 1`

func RemoveBashLocaleWarning(s string) string {
	re := regexp.MustCompile(`bash: warning: setlocale: LC_ALL: cannot change locale \((.+)\)\n`)
	return re.ReplaceAllString(s, "")

}

func main() {
	fmt.Println(info)
	fmt.Println(RemoveBashLocaleWarning(info))

}
