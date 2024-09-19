package main

import (
	"ecomm/db"
	"ecomm/kafka/consumer"
	"ecomm/router"
	"ecomm/rpc/user"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	startService()

	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.StartConsumer("user", &wg)

	r := gin.New()
	router.InitRouter(r)
	r.Run(":8082")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server shutdown ...")
	wg.Wait()
}

func startService() {
	go db.ConnectMysql()
	go db.ConnectRedis()
	s := user.Server{}
	s.StartRpcService()
}
