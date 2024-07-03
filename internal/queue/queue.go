package core

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Codesmith28/cheatScript/api"
	"github.com/Codesmith28/cheatScript/internal"
	"github.com/Codesmith28/cheatScript/internal/clipboard"
)

// Queue is a struct that holds the connection to the RabbitMQ server
type Queue struct {
	conn        *amqp.Connection
	isConnected bool
	Message     chan string
}

// NewQueue creates a new Queue struct
func NewQueue() *Queue {
	return &Queue{
		isConnected: false,
		Message:     make(chan string),
	}
}

// Connect connects to the RabbitMQ server
func (q *Queue) Connect() error {
	if q.isConnected {
		return nil
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkNilErr(err)

	q.conn = conn
	q.isConnected = true
	return nil
}

// Close closes the connection to the RabbitMQ server
func (q *Queue) Close() error {
	return q.conn.Close()
}

func (q *Queue) Publish(message string) error {
	ch, err := q.conn.Channel()
	checkNilErr(err)
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"Prompt",
		false,
		true,
		false,
		false,
		nil,
	)
	checkNilErr(err)

	// fmt.Println("Queue declared", queue)

	err = ch.Publish(
		"",
		"Prompt",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		},
	)
	checkNilErr(err)

	return nil
}

func (q *Queue) Consume(clipboard *clipboard.Clipboard) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkNilErr(err)

	ch, err := conn.Channel()
	checkNilErr(err)

	go func() {
		for {
			msgs, err := ch.Consume(
				"Prompt",
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			checkNilErr(err)

			for d := range msgs {
				var prompt internal.Prompt
				err := json.Unmarshal(d.Body, &prompt)
				checkNilErr(err)

				clipboard.Mu.Lock()
				content, _ := api.SendPrompt(prompt.PromptString, prompt.Model)
				clipboard.Mu.Unlock()
				q.Message <- content
			}
		}
	}()
}

func (q *Queue) GetMessages() (string, error) {
	return <-q.Message, nil
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
