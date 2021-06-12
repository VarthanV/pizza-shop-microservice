package inmemorydb

import (
	"context"

	"github.com/VarthanV/kitchen/queue"
)

type Queue struct {
	Requests []queue.OrderQueueRequest `json:"order_requests"`
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q Queue) Enqueue(ctx context.Context, request queue.OrderQueueRequest) *Queue {
	q.Requests = append(q.Requests, request)
	return &q
}
