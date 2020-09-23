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
	"strings"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
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
	brokers []string
}

// NewKafkaEventBus creates a new kafka producer and returns client as EventBus
// A comma separated list of server addresses is to be passed in 'brokers'
func NewKafkaEventBus(brokers string) EventBus {
	return &kafkaEventBus{
		brokers: strings.Split(brokers, ","),
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
		go func() {
			writer := kafka.NewWriter(kafka.WriterConfig{
				Brokers:  p.brokers,
				Topic:    event.Topic,
				Balancer: &kafka.LeastBytes{},
				// TO DO: Add Custom Dialer when cofigs available
				// From PBI Infra to be able to use SSL. For now, its default
			})
			defer writer.Close()

			schemaIDBytes := make([]byte, 4)
			binary.BigEndian.PutUint32(schemaIDBytes, event.SchemaID)

			var payload []byte
			payload = append(payload, byte(0))
			payload = append(payload, schemaIDBytes...)
			payload = append(payload, event.Data...)

			msg := kafka.Message{
				Key:   event.Key,
				Value: payload,
			}

			switch err := writer.WriteMessages(context.Background(), msg).(type) {
			case nil:
				log.WithFields(log.Fields{
					"topic": event.Topic,
					"key":   string(event.Key),
				}).Info("kafka: delivered message successfully")
			default:
				log.WithError(err).WithFields(log.Fields{"topic": event.Topic, "key": string(event.Key),
					"broker": p.brokers}).Error("kafka: failed to deliver message ")
			}
		}()
	}

	return nil
}

// Suscribe is implemented for the sake of complying with eventbus interface
func (p *kafkaEventBus) Subscribe(topic string, handler func(event EventMessage) error) {
	log.Error("unimplemented function 'Subscribe' called")
}
