package repositories

import (
	"context"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

type VirtualSecurity interface {
	Orders(ctx context.Context, token string, req *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error)
	Positions(ctx context.Context, token string, req *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error)
	SendOrderStock(ctx context.Context, token string, req *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error)
	SendPrice(ctx context.Context, req *kabuspb.Board) error
}
