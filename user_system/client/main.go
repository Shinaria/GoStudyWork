package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"

	"user_system/kitex_gen/GoStudyWork/userSystem/user"
	"user_system/kitex_gen/GoStudyWork/userSystem/user/userService"
)

func main(){
	client, err := userservice.NewClient("GoStudyWork.userSystem.user",client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	i := 1
	for {
		respond, err:= client.GetUserInfo(context.Background(),&user.UserID{No: int64(i)})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(respond.Name)
		i++
		time.Sleep(time.Second)
	}

}