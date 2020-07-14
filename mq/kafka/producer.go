package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var Producer *KafkaProducer

func InitProducer(kafkaUrls []string, topic string, config *sarama.Config) {
	p, err := NewKafkaProducer(kafkaUrls, topic, config)
	if err != nil {
		fmt.Println("New Kafka Producer Failed.")
	} else {
		Producer = p
	}
}

type KafkaProducer struct {
	sarama.AsyncProducer
	Topic string
}

func NewKafkaProducer(kafkaUrls []string, topic string, config *sarama.Config) (*KafkaProducer, error) {
	asyncProducer, err := sarama.NewAsyncProducer(kafkaUrls, config)
	if err != nil {
		return nil, err
	}

	if err := NewTopic(kafkaUrls, topic, 3); err != nil {
		fmt.Printf("NewKafkaProducer: %v\n", err)
	}

	kp := &KafkaProducer{asyncProducer, topic}

	go func() {
		for {
			select {
			case suc := <-kp.Successes():
				fmt.Printf(
					"send msg to kafka topic successfully. partition: %d, offset: %d, timestamp: %s\n",
					suc.Partition,
					suc.Offset,
					suc.Timestamp.String(),
				)
			case fail := <-kp.Errors():
				fmt.Printf(
					"send msg to kafka topic fail. error: %s",
					fail.Err.Error(),
				)
			}
		}
	}()
	return kp, nil
}

func (kp *KafkaProducer) MakeMsg(key, value string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic: kp.Topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}
	return msg
}

func (kp *KafkaProducer) SendMsg(key, value string) error {
	msg := kp.MakeMsg(key, value)
	kp.Input() <- msg
	return nil
}
