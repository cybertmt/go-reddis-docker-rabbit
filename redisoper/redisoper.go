package main

import (
	"net"
	"os"
	"os/signal"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/go-redis/redis"
)

var redis_host = os.Getenv("REDIS_HOST")
var redis_port = os.Getenv("REDIS_PORT")
var redis_db = os.Getenv("REDIS_DB")
var redis_password = os.Getenv("REDIS_PASSWORD")
var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

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

	msgChan := make(chan string, 1)

	go consume(msgChan)

	for message := range msgChan {

		err = client.Publish("10.22.22.32_lastmsg", message).Err()
		if err != nil {
			log.Println(err)
		}
		err = client.Set("10.22.22.32_lastmsg", message, 0).Err()
		if err != nil {
			log.Println(err)
		}
		log.Printf("Sent successfully: {10.22.22.32_lastmsg: %s} \n", message)

	}

	// val, err := client.Get("id1234").Result()
	// if err != nil {
	// 	log.Println(err)
	// }
	//log.Println(val)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh
	log.Println("Redisoper stopped")
}

func consume(msgChan chan string) {

	conn, err := amqp.Dial("amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/")

	if err != nil {
		log.Printf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Printf("%s: %s", "Failed to open a channel", err)
	}

	q, err := ch.QueueDeclare(
		"publisher", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		log.Printf("%s: %s", "Failed to declare a queue", err)
	}

	log.Println("Channel and Queue established")

	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Printf("%s: %s", "Failed to register consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			msgChan <- string(d.Body)

			d.Ack(false)
		}
	}()

	log.Println("Running...")
	<-forever
}
