package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type BootstrapYamlConfig struct {
	Service struct {
		Port string `yaml:"port"`
	} `yaml:"service"`
	Log struct {
		ErrorOutput string `yaml:"errorOutput"`
		InfoOutput  string `yaml:"infoOutput"`
	} `yaml:"log"`
	MySQL struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`
	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"redis"`
}

func GetBootstrapConfig() BootstrapYamlConfig {
	config := BootstrapYamlConfig{}

	dataBytes, err := os.ReadFile("./config/bootstrap.yml")

	if err != nil {
		fmt.Println("读取文件失败：", err)
		return config
	}

	err = yaml.Unmarshal(dataBytes, &config)

	if err != nil {
		fmt.Println("解析 YAML 文件失败：", err)
		return config
	}

	return config
}
