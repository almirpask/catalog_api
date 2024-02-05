package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://admin:admin@rabbitmq:5672")

	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Publish(ctx context.Context, ch *amqp.Channel, routingKey, body, exchange string) error {
	err := ch.PublishWithContext(ctx, exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/json",
		Body:        []byte(body),
	})

	if err != nil {
		return err
	}

	return nil
}
