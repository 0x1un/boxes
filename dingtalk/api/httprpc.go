package api

import (
	"net/url"
)

// Unmarshalable 数据转换接口
type Unmarshalable interface {
}

/*
httpRpc 此方法兼容get & post方法
@respData 返回的数据存储地
@reqData 请求携带的数据
@params 请求携带的参数
@path 请求的uri路径
*/
func (c *DingTalkClient) httpRPC(
	path string,
	params url.Values,
	reqData interface{},
	respData Unmarshalable) error {

	return nil
}
