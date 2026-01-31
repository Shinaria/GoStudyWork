package main

import (
	"log"
	userServer "user_system/kitex_gen/GoStudyWork/userSystem/user/userservice"
)

func main() {
	svr := userServer.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
