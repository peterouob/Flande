package main

import (
	"ecomm/kafka/consumer"
	"ecomm/router"
	"ecomm/rpc"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go consumer.StartConsumer("user", &wg)

	s := rpc.Server{}
	s.StartRpcService()

	r := gin.New()
	router.InitRouter(r)
	r.Run(":8082")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server shutdown ...")

	wg.Wait()
	log.Println("Kafka consumers stopped ...")
	log.Println("Server exiting !")
}
