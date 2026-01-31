package main

import (
	"log"
	"net"
	userServer "user_system/kitex_gen/GoStudyWork/userSystem/user/userservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
)

func main() {
	narcosRegistry, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GET registry success")
	svr := userServer.NewServer(
		new(UserServiceImpl),
		server.WithRegistry(narcosRegistry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "GoStudyWork.userSystem.user"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}),
	)

	if err := svr.Run(); err != nil {
		log.Println("SERVER stopped with error: ", err.Error())
	} else {
		log.Println("SERVER stopped")
	}
}
