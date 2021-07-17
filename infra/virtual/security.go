package virtual

import (
	"context"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
	vs "gitlab.com/tsuchinaga/kabus-virtual-security"
)

func NewSecurity(virtual vs.VirtualSecurity) repositories.Security {
	return &security{virtual: virtual}
}

type security struct {
	repositories.Security
	virtual vs.VirtualSecurity
}

func (s *security) SendOrderStock(_ context.Context, _ string, req *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error) {
	res, err := s.virtual.StockOrder(toStockOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return fromOrderResult(res), nil
}

func (s *security) Orders(_ context.Context, _ string, _ *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error) {
	res, err := s.virtual.StockOrders()
	if err != nil {
		return nil, err
	}
	return fromStockOrders(res), nil
}

func (s *security) Positions(_ context.Context, _ string, _ *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error) {
	res, err := s.virtual.StockPositions()
	if err != nil {
		return nil, err
	}
	return fromStockPositions(res), nil
}
