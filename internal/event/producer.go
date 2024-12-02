package event

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokerAddress, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
	return &KafkaProducer{writer: writer}
}

func (p *KafkaProducer) Produce(eventType, message string) error {
	msg := kafka.Message{
		Key:   []byte(eventType),
		Value: []byte(message),
	}
	if err := p.writer.WriteMessages(context.Background(), msg); err != nil {
		log.Printf("failed to write message to Kafka: %v", err)
		return err
	}
	return nil
}

func (p *KafkaProducer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Printf("failed to close Kafka writer: %v", err)
	}
}