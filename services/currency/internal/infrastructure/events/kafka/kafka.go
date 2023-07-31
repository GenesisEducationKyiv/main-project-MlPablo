package kafka

import (
	"github.com/Shopify/sarama"
)

type Kafka struct {
	p sarama.SyncProducer
}

func CreateKafka(c *Config) (*Kafka, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{c.Address}, cfg)
	if err != nil {
		return nil, err
	}

	return &Kafka{p: producer}, nil
}

func (k *Kafka) Publish(topic, value string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	}

	_, _, err := k.p.SendMessage(msg)

	return err
}
