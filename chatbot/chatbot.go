package chatbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	network "github.com/0x1un/boxes/component/net"
)

const (
	BASE_URL  = `https://oapi.dingtalk.com/robot/send?access_token=`
	FILE_NAME = "/var/log/chatbot.log"
)

type Message struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func Send(tokens, atUsers []string, notifyAll bool, text, title string) {
	logFile, _ := os.OpenFile(FILE_NAME, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	Log := log.New(logFile, "[Info]", log.Ldate|log.Ltime) // log.Ldate|log.Ltime|log.Lshortfile
	Log.Println("开始发送消息!")
	msg := &Message{}
	msg.MsgType = "markdown"
	msg.Markdown.Title = title
	msg.Markdown.Text = text
	msg.At.AtMobiles = atUsers
	msg.At.IsAtAll = notifyAll
	msgs, err := json.Marshal(msg)
	if err != nil {
		Log.Fatal(err)
	}
	fmt.Println(string(msgs))
	for _, tk := range tokens {
		fillMsgAndSent(tk, msgs, Log)
	}
}

//发送消息到钉钉
func fillMsgAndSent(token string, msg []byte, Log *log.Logger) {
	reader := bytes.NewReader(msg)
	resp := network.Post(BASE_URL+token, reader)
	Log.SetPrefix("[Info]")
	Log.Printf("消息发送完成,服务器返回内容：%s", string(resp))
}
