package inmemorydb

import (
	"context"

	"github.com/VarthanV/kitchen/queue"
)

type Queue struct {
	Requests []queue.OrderQueueRequest `json:"order_requests"`
}

func (q Queue) Dequeue(ctx context.Context) *queue.OrderQueueRequest {
	if len(q.Requests) == 0 {
		return nil
	}
	order := q.Requests[0]
	q.Requests = append(q.Requests, q.Requests[1:]...)
	return &order
}

func (q Queue) Enqueue(ctx context.Context, request queue.OrderQueueRequest) {
	q.Requests = append(q.Requests, request)
}
