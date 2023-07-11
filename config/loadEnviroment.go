package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var appConfig Configure

type Configure struct {
	Version string `yaml:"Version"`
	Config  struct {
		NameServices       string `yaml:"warmup"`
		Env                string `yaml:"Env"`
		LogLevel           string `yaml:"LogLevel"`
		Host               string `yaml:"Host"`
		Port               string `yaml:"Port"`
		LocationTime       string `yaml:"LocationTime"`
		ReadTimeoutSecond  int    `yaml:"ReadTimeoutSecond"`
		WriteTimeoutSecond int    `yaml:"WriteTimeoutSecond"`
		IdleTimeoutSecond  int    `yaml:"IdleTimeoutSecond"`
		LogType            string `yaml:"LogType"`
	} `yaml:"Config"`
	OpenTelemetry struct {
		Enabled bool   `yaml:"Enabled"`
		Host    string `yaml:"Host"`
		Port    string `yaml:"Port"`
	} `yaml:"OpenTelemetry"`
	HTTPReq struct {
		Timeout int  `yaml:"Timeout"`
		Retry   int  `yaml:"Retry"`
		Debug   bool `yaml:"Debug"`
	} `yaml:"HTTPReq"`
	AdapterAz struct {
		Enable bool   `yaml:"Enable"`
		Host   string `yaml:"Host"`
		Port   string `yaml:"Port"`
	} `yaml:"AdapterAz"`
	Postgresql struct {
		Enabled  bool   `yaml:"Enabled"`
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		Database string `yaml:"Database"`
		Timeout  string `yaml:"Timeout"`
		SslMode  string `yaml:"SslMode"`
		LogLevel string `yaml:"LogLevel"`
	} `yaml:"Postgresql"`
	Redis struct {
		Enabled  bool   `yaml:"Enabled"`
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Password string `yaml:"Password"`
	} `yaml:"Redis"`
	Cors struct {
		Methods          []string `yaml:"Methods"`
		Headers          []string `yaml:"Headers"`
		Origins          []string `yaml:"Origins"`
		ExposeHeader     []string `yaml:"ExposeHeader"`
		AllowCredentials bool     `yaml:"AllowCredentials"`
		AllowAllOrigins  bool     `yaml:"AllowAllOrigins" default:"true"`
		MaxAge           int      `yaml:"MaxAge"`
	} `yaml:"Cors"`
}

func Load() {
	viper.SetConfigName("app")

	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// config := &Configure{}
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	// appConfig = config
}
