package kafka

import (
	"github.com/Shopify/sarama"
)

const kafkaConn = "kafka:9092"

type Kafka struct {
	p sarama.SyncProducer
}

func CreateKafka() (*Kafka, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{kafkaConn}, cfg)
	if err != nil {
		return nil, err
	}

	admin, err := sarama.NewClusterAdmin([]string{kafkaConn}, cfg)
	if err != nil {
		return nil, err
	}

	err = admin.CreateTopic("logs", &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
	if err != nil {
		return nil, err
	}

	return &Kafka{p: producer}, nil
}

func (k *Kafka) Publish() error {
	msg := &sarama.ProducerMessage{
		Topic: "logs",
		Value: sarama.StringEncoder("bibiks"),
	}

	_, _, err := k.p.SendMessage(msg)

	return err
}
