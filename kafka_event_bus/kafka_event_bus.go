// Package kafka_event_bus is a wrapper above github.com/segmentio/kafka-go
// Title: kafka_event_bus
// Files: kafka_event_bus.go
//
// Author: Matee ullah
// Description: This package provides an interface implementation
// of EventBus for kafka. The process to work with this package
// involves instantiating with NewKafkaEventBus(addrs) as the
// first step and then calling Send(context,message..) function
// to publish message on the kafka topic. Send(context,message)
// spawns another goroutines & publishes the messages in async
package kafka_event_bus

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// KafkaMessage is the expected argument of Send Method
type KafkaMessage struct {
	Topic    string
	Key      []byte
	SchemaID uint32
	Data     []byte
}

// kafkaEventBus implements EventBus interface
type kafkaEventBus struct {
	brokers string
}

// NewKafkaEventBus creates a new kafka producer and returns client as EventBus
// A comma separated list of server addresses is to be passed in 'brokers'
func NewKafkaEventBus(brokers string) EventBus {

	return &kafkaEventBus{
		brokers: brokers,
	}
}

// Send publishes the messages on given topic asynchronously
// this function does not guarantee the delivery of message.
// It expects one or more KafkaMessage objects in args
func (p *kafkaEventBus) Send(ctx context.Context, e ...interface{}) error {
	for i := 0; i < len(e); i++ {
		// args must be of KafkaMessage Type, otherwise we'll get error
		event, ok := e[i].(KafkaMessage)
		if !ok {
			err := errors.New("unable to convert interface to KafkaMessage type")
			log.WithError(err).Error("invalid arg. received")
			return fmt.Errorf("convert interface arg. to KafkaMessage: %v", err)
		}

		// we do not want to disrupt the game flow
		// hence, a separate goroutine is spawned for
		// posting each message onto kafka
		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": p.brokers,
			"security.protocol": "SSL"})
		if err != nil {
			panic(err)
		}

		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()

		fmt.Println("conn success")
		schemaIDBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(schemaIDBytes, event.SchemaID)

		var payload []byte
		payload = append(payload, byte(0))
		payload = append(payload, schemaIDBytes...)
		payload = append(payload, event.Data...)

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &event.Topic, Partition: kafka.PartitionAny},
			Value:          payload,
			Key:            event.Key,
		}, nil)

		if err != nil {
			fmt.Println("unable to produce messgaes: - err: " + err.Error())
		}

		time.Sleep(10 * time.Second)
	}

	return nil
}

// Suscribe is implemented for the sake of complying with eventbus interface
func (p *kafkaEventBus) Subscribe(topic string, handler func(event EventMessage) error) {
	log.Error("unimplemented function 'Subscribe' called")
}
