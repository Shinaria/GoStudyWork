package main

import (
	"log"
	user "user_system/kitex_gen/GoStudyWork/userSystem/user/userservice"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
