package api

import (
	"bytes"
	"encoding/json"
	"net/url"

	network "github.com/0x1un/boxes/component/net"
	"github.com/0x1un/boxes/dingtalk/misc"
)

// Methods

func (f *FormValues) add(key, value string) {
	*f = append(*f, FormComponentValuesVo{
		Name:  key,
		Value: value,
	})
}

// 私人定制表单，非通用
// 城市，臺席號，域帳號，聯繫方式，故障類型，故障範圍，故障現象
// 這裏的工單爲現在自用的，具體實現靈活的用map或struct實現即可。
func FillFormTemplate(city, local, ad_user, contact, fault_type, fault_range, faults string) (formValues FormValues) {
	formValues.add("城市", city)
	formValues.add("台席号", local)
	formValues.add("域账号", ad_user)
	formValues.add("联系方式", contact)
	formValues.add("故障类型", fault_type)
	formValues.add("故障范围", fault_range)
	formValues.add("故障现象", faults)
	return formValues
}

func (self *DingTalkClient) SendProcess(formComponent FormValues) (*CreateProcessInstanceResp, error) {

	var (
		processResp *CreateProcessInstanceResp
	)
	self.ProcessReq.FormComponentValues = formComponent

	data, err := json.Marshal(self.ProcessReq)
	if err != nil {
		return nil, err
	}
	d := bytes.NewReader(data)
	url := "https://oapi.dingtalk.com/topapi/processinstance/create?access_token=" + self.AccessToken
	result := network.Post(url, d)

	if err := json.Unmarshal(result, &processResp); err != nil {
		return nil, err
	}
	return processResp, nil
}

func (self *DingTalkClient) GetProcessInstanceDetail(processid string) (*ProcessInstanceDetail, error) {
	processInstance := new(ProcessInstanceDetail)
	urlParam := make(url.Values)
	params := make(misc.Data)
	urlParam.Set("access_token", self.AccessToken)
	params.Set("process_instance_id", processid)
	data, err := self.Post("topapi/processinstance/get", urlParam, params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, processInstance); err != nil {
		return nil, err
	}
	return processInstance, nil
}
