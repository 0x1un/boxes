package main

import (
	"errors"
	"fmt"
	"log"
	"peeka/internal/dingtalk/api"
)

func SendToProv(c *api.DingTalkClient, resp UserReqResponse) error {
	formValueArray := ConvertUserRequest(resp)
	for _, v := range formValueArray {
		response, err := c.SendProcessForTest(v)
		if response.ErrCode != 0 && err != nil {
			return errors.New(fmt.Sprintf("%s", response.ErrMsg))
		}
		log.Printf("Sent a ticket succeed! status code: %d", response.ErrCode)
	}
	return nil
}