package ip

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ipipdotnet/ipdb-go"
)

func TestIpIp(t *testing.T) {
	db, err := ipdb.NewCity("/tmp/ipip_v4.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	//db.Reload("/path/to/city.ipv4.ipdb") // 更新 ipdb 文件后可调用 Reload 方法重新加载内容

	t.Logf("%v\n", db.IsIPv4())                      // check database support ip type
	t.Logf("%v\n", db.IsIPv6())                      // check database support ip type
	t.Logf("%v\n", db.BuildTime())                   // database build time
	t.Logf("%v\n", db.Languages())                   // database support language
	t.Logf("%v\n", db.Fields())                      // database support fields
	res, err := db.FindInfo("114.114.114.114", "CN") // return CityInfo
	assert.Equal(t, nil, err)
	t.Logf("%+v\n", res)

	res1, err := db.FindInfo("114.114.114.114", "EN") // return CityInfo
	assert.Equal(t, nil, err)
	t.Logf("%+v\n", res1)

	res2, err := db.Find("1.1.1.1", "CN") // return []string
	assert.Equal(t, nil, err)
	t.Logf("%+v\n", res2)

	res3, err := db.FindMap("118.28.8.8", "CN") // return map[string]string
	assert.Equal(t, nil, err)
	t.Logf("%+v\n", res3)

	res4, err := db.FindInfo("127.0.0.1", "CN") // return CityInfo
	assert.Equal(t, nil, err)
	t.Logf("%+v\n", res4)

}
