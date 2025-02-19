package es

import (
	"fmt"
	"log"

	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/pkg/constants"
	"github.com/lukanzx/DouVo/pkg/eslogrus"
	"github.com/sirupsen/logrus"

	"github.com/elastic/go-elasticsearch"
)

var (
	EsClient *elasticsearch.Client
)

func EsHookLog() *eslogrus.ElasticHook {
	hook, err := eslogrus.NewElasticHook(EsClient, config.Elasticsearch.Host, logrus.DebugLevel, constants.APIServiceName)
	if err != nil {
		panic(err)
	}

	return hook
}

// InitEs 初始化es
func Init() {
	esConn := fmt.Sprintf("http://%s", config.Elasticsearch.Addr)
	cfg := elasticsearch.Config{
		Addresses: []string{esConn},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Panic(err)
	}
	EsClient = client
}
