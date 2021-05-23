package implementation

import (
	"context"
	"errors"
	"github.com/VarthanV/pizza/message_queue"
	"github.com/VarthanV/pizza/pizza/services"
	"github.com/golang/glog"

	"github.com/VarthanV/pizza/pizza"
	"github.com/VarthanV/pizza/pizza/models"
	"github.com/google/uuid"
)

type orderservice struct {
	repo             models.OrderRepository
	cartService      services.CartService
	orderitemservice services.OrderItemService
	queueService message_queue.QueueService
}

func NewOrderService(repo models.OrderRepository, cartsvc services.CartService, orderitemsvc services.OrderItemService,quesvc message_queue.QueueService) pizza.OrderService {
	return &orderservice{
		repo:             repo,
		cartService:      cartsvc,
		orderitemservice: orderitemsvc,
		queueService: quesvc,
	}
}

func (o orderservice) CreateOrder(ctx context.Context, userID string) (err error) {
	//See if there is a cart for this given user
	cart, err := o.cartService.GetCart(ctx, userID)
	if err != nil {
		glog.Errorf("Unable to place for this user %s error getting cart items", userID)
		return err
	}
	if len(*cart) == 0 {
		glog.Errorf("The user doesnt have items in cart an order cannot be placed")
		return errors.New("no-cart")
	}
	order := models.Order{}
	// Assign a uuid to the order
	order.OrderUUID = uuid.NewString()
	order.OrderStatus= "Placed"

	/*1) Start a transaction.
	2) Insert into orders table
	3) Convert all the cart items into order item
	4) Return success or err based on the outcome
	5) Start a go_routine in parallel to make the cart_items inactive
	*/
	createErr := o.repo.CreateOrder(ctx, order, userID, cart)
	if createErr != nil {
		glog.Errorf("Unable to create order for the userId %f got error %f", createErr, userID)
		return createErr
	}
	for _, item := range *cart {
		err = o.orderitemservice.AddOrderItem(ctx, item.PizzaID, order.OrderUUID, item.Quantity, item.Price)
		if err != nil {
			glog.Errorf("Unable to create order item")
			return err
		}
		item := item
		go func() {
			err = o.cartService.MakeItemInactive(ctx, item.ID)
			if err != nil {
				glog.Errorf("Unable to make cart inactive..")
			}
		}()
	}

	//Dispatch the order details to queue
	go func() {
		//Represents a single order request entity
		var orderDetail message_queue.OrderDetail
		//The
		 request:= message_queue.OrderQueueRequest{
		 	OrderUUID: order.OrderUUID,
		}
		for _ , item := range  *cart{
			orderDetail = message_queue.OrderDetail{
				PizzaID: item.PizzaID,
				Quantity: item.Quantity,
			}
			request.Details = append(request.Details,orderDetail)
		}
		err := o.queueService.PublishOrderDetails(ctx,request)
		if err != nil {
			glog.Errorf("Unable to dispatch the order details to the queue %f",err)
		}
	}()

	return nil
}

func (o orderservice) GetOrderByUUID(ctx context.Context, uuid string) (*models.Order, error) {
	order, err := o.repo.GetOrderByUUID(ctx, uuid)
	return order, err
}

func (o orderservice) GetOrdersByUserID(ctx context.Context, userId int) (*[]models.Order, error) {
	orders, err := o.repo.GetOrdersByUserID(ctx, userId)
	return orders, err
}
