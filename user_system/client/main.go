package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"

	"user_system/kitex_gen/GoStudyWork/userSystem/user"
	"user_system/kitex_gen/GoStudyWork/userSystem/user/userService"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

func main(){
	narcosResolver, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		log.Fatal(err)
	}
	newClient := userservice.MustNewClient(
		"GoStudyWork.userSystem.user",
		client.WithResolver(narcosResolver),
		client.WithRPCTimeout(time.Second*10),
	)
	i := 1
	for {
		respond, err:= newClient.GetUserInfo(context.Background(),&user.UserID{No: int64(i)})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(respond.Name)
		i++
		time.Sleep(time.Second)
	}

}