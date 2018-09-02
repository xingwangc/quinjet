package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func QueryMsgQueue(host, port, user, pwd, queue, mode string) error {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pwd, host, port)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	if mode == "sync" {
		for msg := range msgs {
			fmt.Println(string(msg.Body))
			if string(msg.Body) == "complete" {
				break
			}
		}
	} else {
		select {
		case msg := <-msgs:
			fmt.Println(string(msg.Body))
		default:
			fmt.Println("running")
		}
	}

	return nil
}
