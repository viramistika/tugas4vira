package services

import (
	"context"

	cm "pnp-master/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) FastPayHandler(ctx context.Context, req cm.FastPayRequest) (res cm.FastPayResponse) {

	return
}
