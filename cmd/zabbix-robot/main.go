package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/0x1un/boxes/chatbot"
)

const (
	CONFIG_FILE = "./zabbix-robot.json"
)

type Config struct {
	RobotTokens []string `json:"robot_tokens"`
	AtUsers     []string `json:"at_users"`
	AtAll       bool     `json:"at_all"`
}

func main() {
	config := *readConfig()
	var (
		content string
		title   string
		area    string
	)
	args := os.Args
	if len(args) == 1 {
		log.Println("no args")
		return
	}
	if len(args) > 2 {
		title = args[2]
	}
	for _, t := range []string{"阿里", "滴滴", "水滴", "VIPKID", "VK", "vipkid", "vk", "温江京东", "温江小米"} {
		area = getGroupName(title, t)
		if area != title {
			break
		}
	}
	content = args[1]
	ct := bytes.NewBufferString(content)
	ct.WriteString("\n\n")
	for _, phone := range config.AtUsers {
		ct.WriteString("@" + phone)
	}
	chatbot.Send(config.RobotTokens, config.AtUsers, config.AtAll, ct.String(), area)
}

func readConfig() *Config {
	bytes, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Fatalln(err.Error())
	}
	config := &Config{}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if len(config.RobotTokens) == 0 {
		log.Fatalln("dingtalk robot token cannot be empty!")
	}
	return config
}

func getGroupName(s, target string) string {
	if strings.Contains(s, target) {
		return fmt.Sprintf("%s%s", target, "项目业务设备")
	}
	return s
}
