package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, msgChan chan *kafka.Message) error {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "imersao12-go-esquenta",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return err
	}

	err = kafkaConsumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	defer func() {
		// Desinscrever tópicos antes de encerrar
		_ = kafkaConsumer.Unsubscribe()

		// Encerrar o consumidor
		_ = kafkaConsumer.Close()
	}()

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err != nil {
			return err
		}

		msgChan <- msg
	}
}
