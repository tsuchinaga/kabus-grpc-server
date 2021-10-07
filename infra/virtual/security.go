package virtual

import (
	"context"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
	vs "gitlab.com/tsuchinaga/kabus-virtual-security"
)

func NewSecurity(virtual vs.VirtualSecurity) repositories.VirtualSecurity {
	return &security{virtual: virtual}
}

type security struct {
	repositories.VirtualSecurity
	virtual vs.VirtualSecurity
}

func (s *security) SendOrderStock(_ context.Context, _ string, req *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error) {
	res, err := s.virtual.StockOrder(toStockOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return fromOrderResult(res), nil
}

func (s *security) SendOrderMargin(_ context.Context, _ string, req *kabuspb.SendMarginOrderRequest) (*kabuspb.OrderResponse, error) {
	res, err := s.virtual.MarginOrder(toMarginOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return fromOrderResult(res), nil
}

func (s *security) Orders(_ context.Context, _ string, req *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error) {
	res := make([]*kabuspb.Order, 0)

	if req.Product == kabuspb.Product_PRODUCT_ALL || req.Product == kabuspb.Product_PRODUCT_STOCK {
		stockOrders, err := s.virtual.StockOrders()
		if err != nil {
			return nil, err
		}
		orders := fromStockOrders(stockOrders)
		for _, o := range orders.Orders {
			res = append(res, o)
		}
	}

	if req.Product == kabuspb.Product_PRODUCT_ALL || req.Product == kabuspb.Product_PRODUCT_MARGIN {
		marginOrders, err := s.virtual.MarginOrders()
		if err != nil {
			return nil, err
		}
		orders := fromMarginOrders(marginOrders)
		for _, o := range orders.Orders {
			res = append(res, o)
		}
	}

	return &kabuspb.Orders{Orders: res}, nil
}

func (s *security) Positions(_ context.Context, _ string, req *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error) {
	res := make([]*kabuspb.Position, 0)

	if req.Product == kabuspb.Product_PRODUCT_ALL || req.Product == kabuspb.Product_PRODUCT_STOCK {
		stockPositions, err := s.virtual.StockPositions()
		if err != nil {
			return nil, err
		}
		positions := fromStockPositions(stockPositions)
		for _, p := range positions.Positions {
			res = append(res, p)
		}
	}

	if req.Product == kabuspb.Product_PRODUCT_ALL || req.Product == kabuspb.Product_PRODUCT_MARGIN {
		marginPositions, err := s.virtual.MarginPositions()
		if err != nil {
			return nil, err
		}
		positions := fromMarginPositions(marginPositions)
		for _, p := range positions.Positions {
			res = append(res, p)
		}
	}

	return &kabuspb.Positions{Positions: res}, nil
}

func (s *security) CancelOrder(_ context.Context, _ string, req *kabuspb.CancelOrderRequest) (*kabuspb.OrderResponse, error) {
	var err error
	prefix := string([]rune(req.OrderId)[:3])
	switch prefix {
	case "sor": // 現物
		err = s.virtual.CancelStockOrder(&vs.CancelOrderRequest{OrderCode: req.OrderId})
	case "mor": // 信用
		err = s.virtual.CancelMarginOrder(&vs.CancelOrderRequest{OrderCode: req.OrderId})
	}
	return &kabuspb.OrderResponse{}, err
}

func (s *security) SendPrice(_ context.Context, req *kabuspb.Board) error {
	if req == nil {
		return nil
	}

	return s.virtual.RegisterPrice(toRegisterPrice(req))
}
