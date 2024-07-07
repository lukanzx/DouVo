package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/elastic/go-elasticsearch"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/lukanzx/DouVo/cmd/user/dal"
	"github.com/lukanzx/DouVo/cmd/user/rpc"
	"github.com/lukanzx/DouVo/config"
	user "github.com/lukanzx/DouVo/kitex_gen/user/userservice"
	"github.com/lukanzx/DouVo/pkg/constants"
	"github.com/lukanzx/DouVo/pkg/eslogrus"
	"github.com/lukanzx/DouVo/pkg/tracer"
	"github.com/lukanzx/DouVo/pkg/utils"
	"github.com/sirupsen/logrus"
)

var (
	path       *string
	listenAddr string
	EsClient   *elasticsearch.Client
)

func Init() {
	path = flag.String("config", "./config", "config path")
	flag.Parse()
	config.Init(*path, constants.UserServiceName)
	dal.Init()
	tracer.InitJaeger(constants.UserServiceName)
	rpc.Init()
	EsInit()
	klog.SetLevel(klog.LevelDebug)
	klog.SetLogger(kitexlogrus.NewLogger(kitexlogrus.WithHook(EsHookLog())))
}

func EsHookLog() *eslogrus.ElasticHook {
	hook, err := eslogrus.NewElasticHook(
		EsClient,
		config.Elasticsearch.Host,
		logrus.DebugLevel,
		constants.UserServiceName)
	if err != nil {
		panic(err)
	}
	return hook
}

func EsInit() {
	esConn := fmt.Sprintf("http://%s", config.Elasticsearch.Addr)
	cfg := elasticsearch.Config{
		Addresses: []string{esConn},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	EsClient = client
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	for index, addr := range config.Service.AddrList {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}
		if index == len(config.Service.AddrList)-1 {
			klog.Fatal("not available port from config")
		}
	}
	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserServiceName,
		}),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(trace.NewDefaultServerSuite()),

		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),
	)
	if err = svr.Run(); err != nil {
		panic(err)
	}
}
