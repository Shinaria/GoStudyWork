package main

import (
	"log"
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
	log.Println(narcosRegistry)
	svr := userServer.NewServer(
		new(UserServiceImpl),
		server.WithRegistry(narcosRegistry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "GoStudyWork.userSystem.user"}),
	)

	if err := svr.Run(); err != nil {
		log.Println("server stopped with error: ", err.Error())
	} else {
		log.Println("server stopped")
	}
}
