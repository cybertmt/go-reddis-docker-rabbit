package main

import (
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
)

var redis_host = os.Getenv("REDIS_HOST")
var redis_port = os.Getenv("REDIS_PORT")
var redis_db = os.Getenv("REDIS_DB")
var redis_password = os.Getenv("REDIS_PASSWORD")

func main() {
	redis_db_int, err := strconv.Atoi(redis_db)
	if err != nil {
		log.Printf("%s: %s", "Error converting redis_db:", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(redis_host, redis_port),
		Password: redis_password,
		DB:       redis_db_int,
	})
	subscriber := client.Subscribe("10.22.22.32_lastmsg")

	ch := subscriber.Channel()
	log.Println("Comsumer started")
	for msg := range ch {
		log.Printf("Channel %s got message %s\n", msg.Channel, msg.Payload)
		time.Sleep(30 * time.Second)
	}
	subscriber.Close()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh
	log.Println("Comsumer stopped")
}
