package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:18082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	cli := kabuspb.NewKabusServiceClient(conn)

	// 先物 直近の限月
	{
		res, err := cli.GetFutureSymbolCodeInfo(
			context.Background(),
			&kabuspb.GetFutureSymbolCodeInfoRequest{
				FutureCode:      kabuspb.FutureCode_FUTURE_CODE_NK225_MINI,
				DerivativeMonth: timestamppb.New(time.Time{}),
			})
		log.Println(res, err)
	}

	// オプション 直近の限月
	{
		res, err := cli.GetOptionSymbolCodeInfo(
			context.Background(),
			&kabuspb.GetOptionSymbolCodeInfoRequest{
				DerivativeMonth: timestamppb.New(time.Time{}),
				CallOrPut:       kabuspb.CallPut_CALL_PUT_CALL,
				StrikePrice:     0,
			})
		log.Println(res, err)
	}
}
