package main

import (
	"errors"
	"strings"

	"net/http"

	"io/ioutil"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ID         string `yaml:"id"`
	Secret     string `yaml:"secret"`
	RegionID   string `yaml:"regionid"`
	RR         string `yaml:"rr"`
	DomainName string `yaml:"domain"`
	RecordType string `yaml:"recordType"`
}

func main() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		logrus.Errorln(err)
	}
	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		logrus.Errorln(err)
	}
	client, err := alidns.NewClientWithAccessKey(config.RegionID, config.ID, config.Secret)
	if err != nil {
		logrus.Errorln(err)
	}
	updateDomainRecordByDomainName(client, config.DomainName, config.RR, config.RecordType)
}

/*
updateDomainRecordByDomainName 修改域名的子域名记录
如果解析的是主域名，rr = @
*/
func updateDomainRecordByDomainName(client *alidns.Client, domainName, rr, typ string) {
	allRecord, err := queryDomainDescribeRecord(client, domainName)
	if err != nil {
		logrus.Errorln(err)
	}
	value, err := getPublicIPAddress()
	if err != nil {
		logrus.Errorln(err)
	}
	for _, record := range allRecord.DomainRecords.Record {
		if record.RR == rr && record.Value != value {
			resp, err := updateDomainRecord(client, record.RecordId, rr, typ, value)
			if err != nil {
				logrus.Errorln(err)
			}
			logrus.Println(resp.RecordId)
		}
	}
}

/*
addDomainRecord 添加域名解析记录
rr 子域名
type 解析类型(A, CNAME...)
value 解析目标ip
*/
func addDomainRecord(client *alidns.Client, domainName, rr, typ, value string) (*alidns.AddDomainRecordResponse, error) {
	req := alidns.CreateAddDomainRecordRequest()
	req.Scheme = "https"
	req.DomainName = domainName
	req.RR = rr
	req.Type = typ
	req.Value = value
	return client.AddDomainRecord(req)
}

/*
queryDomainDescribeRecord 查询域名的解析记录
*/
func queryDomainDescribeRecord(client *alidns.Client, domainName string) (*alidns.DescribeDomainRecordsResponse, error) {
	req := alidns.CreateDescribeDomainRecordsRequest()
	req.Scheme = "https"
	req.DomainName = domainName
	return client.DescribeDomainRecords(req)
}

// querySubDomainDescribeRecord 查询子域名的解析记录
func querySubDomainDescribeRecord(client *alidns.Client, subDomainName string) (*alidns.DescribeSubDomainRecordsResponse, error) {
	req := alidns.CreateDescribeSubDomainRecordsRequest()
	req.Scheme = "https"
	req.SubDomain = subDomainName
	return client.DescribeSubDomainRecords(req)
}

// getPublicIPAddress 获取公网ip地址
func getPublicIPAddress() (string, error) {
	resp, err := http.Get("https://api.ip.sb/ip")
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// updateDomainRecord 更新域名记录
func updateDomainRecord(client *alidns.Client, recordID, rr, typ, value string) (*alidns.UpdateDomainRecordResponse, error) {
	req := alidns.CreateUpdateDomainRecordRequest()
	req.Scheme = "https"
	req.RecordId = recordID
	req.RR = rr
	req.Type = typ
	req.Value = value
	return client.UpdateDomainRecord(req)
}
