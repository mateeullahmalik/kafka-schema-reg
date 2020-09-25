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
	l0 "github.com/mateeullahmalik/kafka-schema-reg/models/v1/events_variantlabs_protobuf_l0"
	l1 "github.com/mateeullahmalik/kafka-schema-reg/models/v1/events_variantlabs_protobuf_l1"
	log "github.com/sirupsen/logrus"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var (
	kafkatopic = "vl-user-dev"
)

func main() {
	ts, _ := ptypes.TimestampProto(time.Now())
	tb, _ := ptypes.TimestampProto(time.Now().AddDate(-20, 0, 0))
	contact := &l0.UserContactInfo{
		Email:  "abc@variantlabs.com",
		Mobile: 3999988,
	}

	Addr := &l0.UserAddressInfo{
		AddressString: "31 Park Lane",
		City:          "London",
		State:         "",
		Country:       "GB",
	}
	// create dummy login event
	reg := &l0.UserRegistered{
		Uid:           "test-reg-user-2",
		FirstName:     "variant",
		LastName:      "labs",
		Username:      "variantlabs",
		DateOfBirth:   tb,
		DateOfJoining: ts,
		Language:      "en",
		AccountStatus: l0.UserAccountStatus_ACTIVE,
		Contact:       contact,
		Address:       Addr,
	}

	event := &l1.UserAll{
		OneofType: &l1.UserAll_Registration{},
	}

	//addr := "127.0.0.1:9092"
	addr := "SSL://b-1.data-infra-msk-de.647dgl.c4.kafka.ap-southeast-1.amazonaws.com:9094,SSL://b-2.data-infra-msk-de.647dgl.c4.kafka.ap-southeast-1.amazonaws.com:9094,SSL://b-3.data-infra-msk-de.647dgl.c4.kafka.ap-southeast-1.amazonaws.com:9094"
	testLocal(event, reg.Uid, addr)

}

func testLocal(obj *l1.UserAll, key string, addr string) {
	// docker-compose.yml can be used to run kafka cluster locally in a quick way
	// command: docker-compose up -d
	// this will run kafka cluster locally on localhost:9092

	// create topic if not exists
	//createTopic(addr, kafkatopic)

	// test producer
	p := kafka_event_bus.NewKafkaEventBus(addr)
	post(p, obj, key)

	time.Sleep(20 * time.Second)
	// confirm delivery by producer  by consuming kafka topic
	//consume(addr)
}

func post(p bus.EventBus, obj *l1.UserAll, key string) {
	data, err := proto.Marshal(obj)
	if err != nil {
		panic("unable to marshal - err: " + err.Error())
	}
	_, vals := obj.Descriptor()

	message := bus.KafkaMessage{
		Topic:          kafkatopic,
		SchemaID:       12,
		Data:           data,
		Key:            []byte(key),
		DescriptorVals: vals,
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

	/*results, err := a.DeleteTopics(ctx, []string{topic}, kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to delete topics: %v\n", err)
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}*/

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
	err = c.SubscribeTopics([]string{kafkatopic}, nil)
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
				event := &l1.UserAll{
					OneofType: &l1.UserAll_Login{},
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
