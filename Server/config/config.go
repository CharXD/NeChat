package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var ServerConfig = Config{}

type Config struct {
	Server Server `json:"server"`
	SQL    SQL    `json:"SQL"`
}

type Server struct {
	Port         string `json:"Port"`
	JWTSingedKey string `json:"JWTSingedKey"`
	JWTIssuer    string `json:"JWTIssuer"`
}

type SQL struct {
	User     string `json:"User"`
	Pass     string `json:"Pass"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	Database string `json:"Database"`
}

func InitConfig() {
	configFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(configFile)
	byteValue, _ := ioutil.ReadAll(configFile)
	_ = json.Unmarshal(byteValue, &ServerConfig)
}
