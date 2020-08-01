package tanyibot_test

import(
	"testing"
	"time"
	"fmt"
	"encoding/json"
	"context"
	"github.com/xopenapi/tanyibot"
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
)


type conf struct {
	AppKey 		string   	`yaml:"appKey"` //yaml：yaml格式 enabled：属性的为enabled
	AppSecret   string 		`yaml:"appSecret"`
	TenantSign	string		`yaml:"tenantSign"`
}

func getClient() *tanyibot.APIClient {
	var conf conf
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	
	cfg := tanyibot.NewConfiguration()	
	client := tanyibot.NewAPIClient(cfg)

	timestamp := fmt.Sprintf("%d", time.Now().UnixNano() / (1000*1000))
	version := "v1"

	params := tanyibot.KVS{
		tanyibot.KV{ K: "appKey", V: conf.AppKey },
		tanyibot.KV{ K: "appSecret", V: conf.AppSecret },
		tanyibot.KV{ K: "tenantSign", V: conf.TenantSign },
		tanyibot.KV{ K: "version", V: version },
		tanyibot.KV{ K: "timestamp", V: timestamp },
	}
	signature := tanyibot.SignArr(params, conf.TenantSign)
	cfg.AddDefaultHeader("appKey", conf.AppKey)
	cfg.AddDefaultHeader("appSecret", conf.AppSecret)
	cfg.AddDefaultHeader("tenantSign", conf.TenantSign)
	cfg.AddDefaultHeader("version", version)
	cfg.AddDefaultHeader("timestamp", timestamp)
	cfg.AddDefaultHeader("signature", signature)

	return client
}

func TestGetIsvList(t *testing.T) {
	client := getClient()
	r1, _, err := client.IsvApi.GetIsvList(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	jsonBytes1, _ := json.Marshal(r1)
	t.Logf("isvList: %v", string(jsonBytes1))
}

func TestGetPhoneList(t *testing.T) {
	client := getClient()
	r2, _, err := client.PhoneApi.GetPhoneList(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	jsonBytes2, _ := json.Marshal(r2)
	t.Logf("%s", string(jsonBytes2))
}

func TestLog(t *testing.T) {
	log.Printf("hello\n")
}

func TestCreateJob(t *testing.T) {
	client := getClient()

	var phone tanyibot.Phone
	{
		r, _, err := client.PhoneApi.GetPhoneList(context.Background())
		if err != nil {
			t.Error(err)
			return
		}	
		
		if len(r.Data) == 0 {
			return
		}
	
		phone = r.Data[0]
	}

	var dialogFlowId int32
	{
		r, _, err := client.DialogFlowApi.GetDialogFlowList(context.Background())
		if err != nil {
			t.Error(err)
			return
		}	
		
		if len(r.Data) == 0 {
			return
		}

		dialogFlowId = r.Data[0].DialogFlowId
	}
	
	createOpts := tanyibot.Job{
		ConcurrencyQuota: 1,
		JobPhoneNumberIdList: []int64{
			phone.TenantPhoneNumberId,
		},
		RobotCallJob: tanyibot.RobotCallJob{
			Name: "API测试任务",
			DailyStartTime: "09:00",
			DailyEndTime: "20:00",
			IsPrior: true,
			Mode: "MANUAL",
			PhoneType: phone.PhoneType,
			RobotCount: 1,
			DialogFlowId: dialogFlowId,
		},
	}

	var robotCallJobId int64
	{
		r, _, err := client.JobApi.Create(context.Background(), createOpts)
		if err != nil {
			t.Error(err)
			return
		}

		if r.Code != 200 {
			t.Error(r.ResultMsg)
			return
		}

		robotCallJobId = r.Data.RobotCallJobId
	}
	
	{
		req := tanyibot.ImportCustomerReq{
			RobotCallJobId: robotCallJobId,
			CustomerPersons: []tanyibot.CustomerPerson{
				tanyibot.CustomerPerson{
					Name: "x先生",
					PhoneNumber: "18311112222",
				},
			},
		}

		r, _, err := client.JobApi.ImportCustomer(context.Background(), req)
		if err != nil {
			t.Error(err)
			return
		}

		if r.Code != 200 {
			t.Error(r.ResultMsg)
			return
		}
	}
	
	{
		r, _, err := client.JobApi.Start(context.Background(), tanyibot.InlineObject{
			RobotCallJobId: robotCallJobId,
		})

		if err != nil {
			t.Error(err)
			return
		}

		if r.Code != 200 {
			t.Error(r.ResultMsg)
			return
		}
	}

	t.Logf("finish")
}