package core

import (
	"encoding/json"

	"github.com/Codesmith28/cheatScript/api"
	"github.com/Codesmith28/cheatScript/internal"
	"github.com/Codesmith28/cheatScript/internal/clipboard"
	amqp "github.com/rabbitmq/amqp091-go"
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
	if err != nil {
		return err
	}

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
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"Prompt",
		false,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

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

	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) Consume(clipboard *clipboard.Clipboard) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

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

			if err != nil {
				panic(err)
			}

			for d := range msgs {
				var prompt internal.Prompt
				err := json.Unmarshal(d.Body, &prompt)
				if err != nil {
					panic(err)
				}

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
