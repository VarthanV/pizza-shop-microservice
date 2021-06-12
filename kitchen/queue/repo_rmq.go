package queue

import (
	"context"
	"encoding/json"

	"github.com/golang/glog"
	"github.com/streadway/amqp"
)

type rabbitmqrepo struct {
	ch *amqp.Channel
}

func NewRabbitRepository(ch *amqp.Channel) QueueRepository {
	return rabbitmqrepo{
		ch: ch,
	}
}
func (r rabbitmqrepo) getNewQueue(queueName string) (amqp.Queue, error) {
	q, err := r.ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return q, err
}

func (r rabbitmqrepo) PublishOrderStatus(ctx context.Context, orderUUID string, status string) error {
	queue, err := r.getNewQueue("order-status")
	if err != nil {
		glog.Errorf("Unable to declare queue %s", err)
		return err
	}
	request := OrderStatusUpdateRequest{
		OrderUUID: orderUUID,
		Status:    status,
	}
	payload, err := json.Marshal(request)
	if err != nil {
		glog.Errorf("Error while marshalling the request %s", err)
	}
	err = r.ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		},
	)
	return nil
}

func (r rabbitmqrepo) ConsumeOrderDetails(ctx context.Context) (<-chan amqp.Delivery, error) {
	queue, err := r.getNewQueue("new-order")
	if err != nil {
		glog.Errorf("Unable to declare queue %s", err)
		return nil, err
	}
	msgs, err := r.ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	return msgs, nil
}
