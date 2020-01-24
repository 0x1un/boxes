package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func ListenITOP(url string, data io.Reader) <-chan UserReqResponse {

	return nil
}

// 返回来自itop的标准门户工单数据
func FetcheFromITOP(url string, data io.Reader) UserReqResponse {
	resp, err := request(http.MethodPost, url, data)
	if err != nil {
		panic(err)
	}
	t := new(UserReqResponse)
	if err := json.Unmarshal(resp, t); err != nil {
		panic(err)
	}
	return *t

}

// 简单封装的http请求
func request(method, url string, data io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	switch data.(type) {
	case *strings.Reader:
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case *bytes.Reader:
		req.Header.Add("Content-Type", "application/json")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}