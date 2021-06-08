package implementation

import (
	"context"

	"github.com/VarthanV/kitchen/processes"
	"github.com/golang/glog"
)

var err error

type orderprocessupdateimplementation struct {
	repo processes.OrderProcessUpdateRepo
}

func NewOrderOrderProcessUpdateImplementation(repo processes.OrderProcessUpdateRepo) processes.OrderProcessUpdateService {
	return &orderprocessupdateimplementation{
		repo: repo,
	}

}

func (o orderprocessupdateimplementation) MarkOrderComplete(ctx context.Context, orderUUID string, cookID int) error {
	err = o.repo.UpdateOrderProcces(ctx, orderUUID, cookID)
	if err != nil {
		glog.Errorf("Error while updating order_process %s", err)
	}
	return err
}

func (o orderprocessupdateimplementation) MarkOrderItemComplete(ctx context.Context, pizzaID int, orderUUID string) error {
	err = o.repo.UpdateOrderItemProcess(ctx, pizzaID, orderUUID)
	if err != nil {
		glog.Errorf("Error while updating order_item_process %s", err)
	}
	return err
}
