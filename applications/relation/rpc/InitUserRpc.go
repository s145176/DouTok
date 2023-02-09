package rpc

import (
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var UserClient userservice.Client

func InitUserRpc() {
	v := conf.UserConfig
	EtcdAddress := fmt.Sprintf("%s:%d", v.GetString("etcd.address"), v.GetInt("etcd.port"))

	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := v.GetString("server.name")

	c, err := userservice.NewClient(
		ServiceName,
		//多路复用会报错
		//client.WithMuxConnection(1),                       // mux
		// tracer
		client.WithResolver(r), // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(err)
	}
	UserClient = c
}