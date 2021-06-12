package implementation

import (
	"context"
	"encoding/json"

	"github.com/VarthanV/kitchen/inmemorydb"
	"github.com/VarthanV/kitchen/queue"
	"github.com/VarthanV/kitchen/shared"
	"github.com/golang/glog"
)

type orderinmemoryimpelementation struct {
	repo inmemorydb.OrderRequestInMemoryRepo
}

func NewOrderInmemoryService(repo inmemorydb.OrderRequestInMemoryRepo) inmemorydb.OrderRequestInMemoryService {
	return &orderinmemoryimpelementation{
		repo: repo,
	}
}

func (o orderinmemoryimpelementation) SetOrder(ctx context.Context, key string, request queue.OrderQueueRequest) error {
	ordersInQueue := inmemorydb.NewQueue()
	var jsonByte []byte
	/*
		1) Get item from the queue
		2) Enqueue the order to the queue
		3) Again set to the inmemory db
	*/
	orderStr := o.repo.GetOrder(ctx, shared.RedisKeyForOrders)
	err := json.Unmarshal([]byte(orderStr), &ordersInQueue.Requests)
	if err != nil {
		glog.Errorf("Unable to unmarshal the json...", err)
		// if err then there are no orders in our system so creating oen
		ordersInQueue = ordersInQueue.Enqueue(ctx, request)
		glog.Infof("%v", ordersInQueue)
	} else {

		// Set it in the key again
		ordersInQueue = ordersInQueue.Enqueue(ctx, request)

		if err != nil {
			glog.Error("Error in marshalling the json...", err)
		}

	}
	jsonByte, err = json.Marshal(ordersInQueue.Requests)
	glog.Info("Orders in queue is...", ordersInQueue.Requests)
	err = o.repo.SetOrder(ctx, shared.RedisKeyForOrders, string(jsonByte))
	if err != nil {
		glog.Info("Successfully set the order to the queue...")
	}
	glog.Info(err)
	return nil
}

func (o orderinmemoryimpelementation) GetOrder(ctx context.Context) (*queue.OrderQueueRequest, error) {
	var ordersInQueue *inmemorydb.Queue
	ordersInQueue = inmemorydb.NewQueue()
	orderStr := o.repo.GetOrder(ctx, shared.RedisKeyForOrders)
	err = json.Unmarshal([]byte(orderStr), &ordersInQueue.Requests)
	if err != nil {
		return nil, err
	}
	if len(ordersInQueue.Requests) == 0 {
		return nil, nil
	}
	firstOrder := ordersInQueue.Requests[0]
	ordersInQueue.Requests = ordersInQueue.Requests[1:]
	// Set the q to inmemory db
	go func() {
		jsonByte, err := json.Marshal(ordersInQueue.Requests)
		if err != nil {
			glog.Errorf("Unable to set order in the queue")
		}
		err = o.repo.SetOrder(ctx, shared.RedisKeyForOrders, string(jsonByte))
	}()
	return &firstOrder, err
}
