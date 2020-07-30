package aibot_test

import(
	"testing"
	"time"
	"fmt"
	"encoding/json"
	"context"
	"github.com/xopenapi/aibot"
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
)


type conf struct {
	AppKey 		string   	`yaml:"appKey"` //yaml：yaml格式 enabled：属性的为enabled
	AppSecret   string 		`yaml:"appSecret"`
	TenantSign	string		`yaml:"tenantSign"`
}

func TestGetPhoneList(t *testing.T) {
	var conf conf
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	
	cfg := aibot.NewConfiguration()	
	client := aibot.NewAPIClient(cfg)

	timestamp := fmt.Sprintf("%d", time.Now().UnixNano() / (1000*1000))
	version := "v1"

	params := aibot.KVS{
		aibot.KV{ K: "appKey", V: conf.AppKey },
		aibot.KV{ K: "appSecret", V: conf.AppSecret },
		aibot.KV{ K: "tenantSign", V: conf.TenantSign },
		aibot.KV{ K: "version", V: version },
		aibot.KV{ K: "timestamp", V: timestamp },
	}
	signature := aibot.SignArr(params, conf.TenantSign)
	cfg.AddDefaultHeader("appKey", conf.AppKey)
	cfg.AddDefaultHeader("appSecret", conf.AppSecret)
	cfg.AddDefaultHeader("tenantSign", conf.TenantSign)
	cfg.AddDefaultHeader("version", version)
	cfg.AddDefaultHeader("timestamp", timestamp)
	cfg.AddDefaultHeader("signature", signature)

	r1, _, err := client.IsvApi.GetIsvList(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	jsonBytes1, _ := json.Marshal(r1)
	t.Logf("isvList: %v", string(jsonBytes1))

	r2, _, err := client.PhoneApi.GetPhoneList(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	jsonBytes2, _ := json.Marshal(r2)
	t.Logf("%s", string(jsonBytes2))
}
