package event

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokerAddress, topic, groupID string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupID,
	})
	return &KafkaConsumer{reader: reader}
}

func (c *KafkaConsumer) Consume() {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message from Kafka: %v", err)
			continue
		}
		log.Printf("received event: key=%s value=%s", string(msg.Key), string(msg.Value))
	}
}

func (c *KafkaConsumer) Close() {
	if err := c.reader.Close(); err != nil {
		log.Printf("failed to close Kafka reader: %v", err)
	}
}
