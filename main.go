package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mateeullahmalik/kafka-schema-reg/kafka_event_bus"
	bus "github.com/mateeullahmalik/kafka-schema-reg/kafka_event_bus"
	v1 "github.com/mateeullahmalik/kafka-schema-reg/models/v1"
	log "github.com/sirupsen/logrus"

	"github.com/golang/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	// create dummy login event
	login := &v1.UserLoggedIn{
		Uid:       "test-user",
		Ip:        "555.555.555.0000",
		Country:   "Wakanda",
		Longitude: 321.2,
		Latitude:  207.9,
	}

	event := &v1.UserAll{
		OneofType: &v1.UserAll_Login{
			Login: login,
		},
	}

	//addr := "127.0.0.1:9092"
	addr := "b-1.data-infra-msk-de.647dgl.c4.kafka.ap-southeast-1.amazonaws.com:9094,b-2.data-infra-msk-de.647dgl.c4.kafka.ap-southeast-1.amazonaws.com:9094,b-3.data-infra-msk-de.647dgl.c4.kafka.ap-southeast-1.amazonaws.com:9094"
	testLocal(event, login.Uid, addr)

}

func testLocal(obj *v1.UserAll, key string, addr string) {
	// docker-compose.yml can be used to run kafka cluster locally in a quick way
	// command: docker-compose up -d
	// this will run kafka cluster locally on localhost:9092

	// create topic if not exists
	//createTopic(addr, "vl-user-dev")

	// test producer
	p := kafka_event_bus.NewKafkaEventBus(addr)
	post(p, obj, key)

	time.Sleep(60 * time.Second)
	// confirm delivery by producer  by consuming kafka topic
	//consume(addr)
}

func post(p bus.EventBus, obj *v1.UserAll, key string) {
	data, err := proto.Marshal(obj)
	if err != nil {
		panic("unable to marshal - err: " + err.Error())
	}

	message := bus.KafkaMessage{
		Topic:    "vl-user-dev",
		SchemaID: 12,
		Data:     data,
		Key:      []byte(key),
	}
	if err := p.Send(context.Background(), message); err != nil {
		panic(err)
	}
}

func createTopic(addr, topic string) {
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": addr, "security.protocol": "SSL"})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create topics on cluster.
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}
	results, err := a.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1}},
		// Admin options
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}

	a.Close()
}

func consume(addr string) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               addr,
		"security.protocol":               "SSL",
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"group.id":                        1,
		"go.application.rebalance.enable": true,
		// Enable generation of PartitionEOF when the
		// end of a partition is reached.
		"enable.partition.eof": true,
		"auto.offset.reset":    "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)
	err = c.SubscribeTopics([]string{"vl-user-dev"}, nil)
	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false

		case ev := <-c.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				log.WithFields(log.Fields{"Message": e}).Info("Assigned partition(s)")
				c.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				log.WithFields(log.Fields{"Message": e}).Info("Revoked partition(s)")
				c.Unassign()
			case *kafka.Message:
				message := e.Value
				event := &v1.UserAll{
					OneofType: &v1.UserAll_Login{},
				}
				if err := proto.Unmarshal(message[5:], event); err != nil {
					panic("unmarshal error: " + err.Error())
				}

				fmt.Println("unmarshal success")
				fmt.Println(event)
			case kafka.PartitionEOF:
				log.Info("Reached EOF. (No more messages)")
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				log.WithError(e).Error("couldn't consume message")
			}
		}
	}

	log.Info("Closing Consumer")
	c.Close()
}
