package api

import (
	"net/url"
)

type Unmarshalable interface {
}

func (c *DingTalkClient) httpRPC(path string, params url.Values, reqData interface{}, respData Unmarshalable) error {

	return nil
}
