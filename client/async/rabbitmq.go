package async

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"project/github/comics/client/json/character"
	"project/github/comics/client/json/common"
	"time"
)

func sendMQ(msg []byte) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	err = ch.Publish(
		"producer-q-exchange",
		// this should be the empty if using async/aqmp
		"1",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})

	failOnError(err, "Failed to publish a message")
	return err
}

func SendMQTopicAsyncAQMP(message string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//err = ch.ExchangeDeclare(
	//	"producer-json-golang",
	//	"topic",
	//	false,
	//	true,
	//	false,
	//	true,
	//	nil,
	//)
	//failOnError(err, "Failed to declare a queue")
	_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	err = ch.Publish(
		"producer-q-exchange",
		// this should be the empty if using async/aqmp
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func SendToMQ(dataType common.DataType, message *character.MainType) {
	responseData := &character.CharacterResponseData{
		DataType: dataType,
		RawData:  message,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		log.Println("Error converting data - ", err)
		return
	}

	err = sendMQ(jsonData)
	if err != nil {
		log.Println("Error sending data to mq - ", err)
	}
}
