package core

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Codesmith28/cheatScript/api"
	"github.com/Codesmith28/cheatScript/internal"
)

type Queue struct {
	conn        *amqp.Connection
	channel     *amqp.Channel
	isConnected bool
	messages    <-chan amqp.Delivery
}

func NewQueue() *Queue {
	return &Queue{
		isConnected: false,
	}
}

func (q *Queue) Connect() error {
	if q.isConnected {
		return nil
	}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkNilErr(err)
	ch, err := conn.Channel()
	checkNilErr(err)
	q.conn = conn
	q.channel = ch
	q.isConnected = true
	_, err = ch.QueueDeclare(
		"Prompt",
		false,
		false,
		false,
		false,
		nil,
	)
	checkNilErr(err)
	messages, err := ch.Consume(
		"Prompt",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	checkNilErr(err)
	q.messages = messages
	return nil
}

func (q *Queue) Close() error {
	if q.channel != nil {
		q.channel.Close()
	}
	if q.conn != nil {
		return q.conn.Close()
	}
	return nil
}

func (q *Queue) Publish(query internal.Query) error {
	queryBytes, err := json.Marshal(query)
	checkNilErr(err)
	err = q.channel.Publish(
		"",
		"Prompt",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        queryBytes,
		},
	)
	checkNilErr(err)
	return nil
}

func (q *Queue) Consume() (string, error) {
	select {
	case d, ok := <-q.messages:
		if !ok {
			return "", fmt.Errorf("channel closed")
		}
		var query internal.Query
		err := json.Unmarshal(d.Body, &query)
		checkNilErr(err)
		content, err := api.SendPrompt(query.PromptString, query.SelectedModel, query.InputString)
		checkNilErr(err)
		return content, nil
	}
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
