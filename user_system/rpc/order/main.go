package main

import (
	"log"
	order "user_system/kitex_gen/GoStudyWork/userSystem/order/orderservice"
)

func main() {
	svr := order.NewServer(new(OrderServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
