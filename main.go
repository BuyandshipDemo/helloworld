package main

import (
	helloworld "github.com/BuyandshipDemo/helloworld/kitex_gen/buyandship/marketplace/helloworld/greetservice"
	"github.com/BuyandshipDemo/utils/configs"
	"log"
)

func main() {
	opts := configs.Load("/Users/user/buyandship/helloworld/kitex_info.yaml")
	
	svr := helloworld.NewServer(new(GreetServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
