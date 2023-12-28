package config

import (
	"github.com/spf13/viper"
	"libra.com/common/logs"
	"log"
	"os"
	"path/filepath"
)

var C *Config

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	EC    *EtcdConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type EtcdConfig struct {
	Addrs []string
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	parentDir := filepath.Dir(workDir)
	conf.viper.AddConfigPath(parentDir + "/config/api")
	conf.viper.AddConfigPath(workDir + "/config")
	println(workDir + "/config")

	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conf.InitServerConfig()
	conf.InitEtcdConfig()
	return conf
}

func (c *Config) InitServerConfig() {
	sc := &ServerConfig{
		Name: c.viper.GetString("server.name"),
		Addr: c.viper.GetString("server.addr"),
	}
	c.SC = sc
}

func (c *Config) InitEtcdConfig() {
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalf(err.Error())
	}
	ec := &EtcdConfig{
		Addrs: addrs,
	}
	c.EC = ec
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}
