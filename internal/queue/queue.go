package core

import (
	"time"

	"github.com/Codesmith28/cheatScript/api"
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
			ContentType: "text/plain",
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

	defer ch.Close()

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

	forever := make(chan string)

	go func() {
		for d := range msgs {

			clipboard.Mu.Lock()
			content, _ := api.SendPrompt(string(d.Body))
			clipboard.Mu.Unlock()

			if content == "" {
				continue
			}

			q.Message <- content

			time.Sleep(2 * time.Second)
		}
	}()

	<-forever
}

func (q *Queue) GetMessages() (string, error) {
	return <-q.Message, nil
}
