package message_queue

import (
	"context"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/streadway/amqp"
)

type rabbitmqimplementation struct {
	ch *amqp.Channel
}



func NewRabbitRepository(ch *amqp.Channel) QueueRepository {
	return rabbitmqimplementation{
		ch: ch,
	}
}
func (r rabbitmqimplementation) getNewQueue(queueName string) (amqp.Queue,error) {
	q ,err :=r.ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	return q,err
}

func (r rabbitmqimplementation) PublishOrderDetails(ctx context.Context, request OrderQueueRequest) error {
	queue,err := r.getNewQueue("new-order")
	if  err != nil {
		glog.Errorf("Unable to declare queue %s",err)
		return err
	}
	payload ,err := json.Marshal(request)

	err = r.ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: payload ,
		},
		)
	return nil
}

func (r rabbitmqimplementation)  ConsumeOrderStatus(ctx context.Context) (<- chan amqp.Delivery,error)  {
	queue ,err := r.getNewQueue("order-status")
	if err != nil {
		glog.Errorf("Unable to declare queue %s",err)
		return nil,err
	}
	msgs ,err:=  r.ch.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	return  msgs,nil
}