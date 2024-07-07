package config

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
)

var (
	Server        *server
	Mysql         *mySQL
	Snowflake     *snowflake
	Service       *service
	Jaeger        *jaeger
	Etcd          *etcd
	RabbitMQ      *rabbitMQ
	Redis         *redis
	OSS           *oss
	Elasticsearch *elasticsearch

	runtime_viper = viper.New()
)

func Init(path string, service string) {
	runtime_viper.SetConfigType("yaml")
	localConfigPath := "../config.yaml"

	runtime_viper.SetConfigFile(localConfigPath)

	if err := runtime_viper.ReadInConfig(); err != nil {
		klog.Fatalf("Not config.yaml: %v", err)
	}

	configMapping(service)
	klog.Infof("all keys: %v\n", runtime_viper.AllKeys())
	runtime_viper.WatchConfig()
}

func configMapping(srv string) {

	c := new(config)
	if err := runtime_viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	Snowflake = &c.Snowflake

	Server = &c.Server
	Server.Secret = []byte(runtime_viper.GetString("server.jwt-secret"))

	Jaeger = &c.Jaeger
	Mysql = &c.MySQL
	RabbitMQ = &c.RabbitMQ
	Redis = &c.Redis
	OSS = &c.OSS
	Elasticsearch = &c.Elasticsearch
	Service = GetService(srv)
}

func GetService(srvname string) *service {
	addrlist := runtime_viper.GetStringSlice("services." + srvname + ".addr")

	return &service{
		Name:     runtime_viper.GetString("services." + srvname + ".name"),
		AddrList: addrlist,
		LB:       runtime_viper.GetBool("services." + srvname + ".load-balance"),
	}
}

func InitForTest() {
	Snowflake = &snowflake{
		WorkerID:      0,
		DatancenterID: 0,
	}

	Server = &server{
		Version: "1.0",
		Name:    "DouServer",
		Secret:  []byte("MTAxNTkwMTg1Mw=="),
	}

	Jaeger = &jaeger{
		Addr: "127.0.0.1:6831",
	}

	Etcd = &etcd{
		Addr: "127.0.0.1:2379",
	}

	Mysql = &mySQL{
		Addr:     "127.0.0.1:3306",
		Database: "DouServer",
		Username: "DouServer",
		Password: "DouServer",
		Charset:  "utf8mb4",
	}

	RabbitMQ = &rabbitMQ{
		Addr:     "127.0.0.1:5672",
		Username: "DouServer",
		Password: "DouServer",
	}

	Redis = &redis{
		Addr:     "127.0.0.1:6379",
		Password: "DouServer",
	}
}
