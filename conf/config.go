package conf

import (
	"fmt"
	"mginmall/dao"
	"os"

	"gopkg.in/yaml.v3"
)

var Config config

type config struct {
	Service serviceConfig `yaml:"service"`
	Mysql   mysqlConfig   `yaml:"mysql"`
	Redis   redisConfig   `yaml:"redis"`
	Email   emailConfig   `yaml:"email"`
	Path    pathConfig    `yaml:"path"`
}

type serviceConfig struct {
	AppMode  string `yaml:"AppMode"`
	HttpPort string `yaml:"HttpPort"`
}

type mysqlConfig struct {
	DB         string `yaml:"DB"`
	Dbhost     string `yaml:"Dbhost"`
	Dbport     string `yaml:"Dbport"`
	Dbuser     string `yaml:"Dbuser"`
	DbPassword string `yaml:"DbPassword"`
	Dbname     string `yaml:"Dbname"`
	DsnRead    string `yaml:"DsnRead"`
	DsnWrite   string `yaml:"DsnWrite"`
}

type redisConfig struct {
	RedisDb       string `yaml:"RedisDb"`
	RedisAddr     string `yaml:"RedisAddr"`
	RedisPassword string `yaml:"RedisPassword"`
	RedisDbName   string `yaml:"RedisDbName"`
}

type emailConfig struct {
	ValidEmail string `yaml:"ValidEmail"`
	SmtpEmail  string `yaml:"SmtpEmail"`
	SmtpHost   string `yaml:"SmtpHost"`
	SmtpPass   string `yaml:"SmtpPass"`
}

type pathConfig struct {
	Host        string `yaml:"Host"`
	ProductPath string `yaml:"ProductPath"`
	AvataPath   string `yaml:"AvataPath"`
}

func InitConfig() {
	data, err := os.ReadFile("conf/config.yml")
	if err != nil {
		panic(fmt.Errorf("read config.yml failed: %w", err))
	}

	if err := yaml.Unmarshal(data, &Config); err != nil {
		panic(fmt.Errorf("parse config.yml failed: %w", err))
	}

	pathRead := Config.Mysql.DsnRead
	pathWrite := Config.Mysql.DsnWrite
	dao.Database(pathRead, pathWrite)
}
