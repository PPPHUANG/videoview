package config

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	LbAddr string `json:"lb_addr"`
	Oss string `json:"oss_addr"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}
	err := decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
}

func GetlbAddr() string {
	return configuration.LbAddr
}

func GetOssAddr() string {
	return configuration.Oss
}