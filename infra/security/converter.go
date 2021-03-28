package security

import (
	"time"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toExchange(exchange kabuspb.Exchange) kabus.Exchange {
	switch exchange {
	case kabuspb.Exchange_EXCHANGE_TOUSHOU:
		return kabus.ExchangeToushou
	case kabuspb.Exchange_EXCHANGE_MEISHOU:
		return kabus.ExchangeMeishou
	case kabuspb.Exchange_EXCHANGE_FUKUSHOU:
		return kabus.ExchangeFukushou
	case kabuspb.Exchange_EXCHANGE_SATSUSHOU:
		return kabus.ExchangeSatsushou
	case kabuspb.Exchange_EXCHANGE_ALL_SESSION:
		return kabus.ExchangeAll
	case kabuspb.Exchange_EXCHANGE_DAY_SESSION:
		return kabus.ExchangeDaytime
	case kabuspb.Exchange_EXCHANGE_NIGHT_SESSION:
		return kabus.ExchangeEvening
	default:
		return kabus.ExchangeUnspecified
	}
}

func fromExchange(exchange kabus.Exchange) kabuspb.Exchange {
	switch exchange {
	case kabus.ExchangeToushou:
		return kabuspb.Exchange_EXCHANGE_TOUSHOU
	case kabus.ExchangeMeishou:
		return kabuspb.Exchange_EXCHANGE_MEISHOU
	case kabus.ExchangeFukushou:
		return kabuspb.Exchange_EXCHANGE_FUKUSHOU
	case kabus.ExchangeSatsushou:
		return kabuspb.Exchange_EXCHANGE_SATSUSHOU
	case kabus.ExchangeAll:
		return kabuspb.Exchange_EXCHANGE_ALL_SESSION
	case kabus.ExchangeDaytime:
		return kabuspb.Exchange_EXCHANGE_DAY_SESSION
	case kabus.ExchangeEvening:
		return kabuspb.Exchange_EXCHANGE_NIGHT_SESSION
	default:
		return kabuspb.Exchange_EXCHANGE_UNSPECIFIED
	}
}

func toFutureCode(futureCode kabuspb.FutureCode) kabus.FutureCode {
	switch futureCode {
	case kabuspb.FutureCode_FUTURE_CODE_NK225:
		return kabus.FutureCodeNK225
	case kabuspb.FutureCode_FUTURE_CODE_NK225_MINI:
		return kabus.FutureCodeNK225Mini
	case kabuspb.FutureCode_FUTURE_CODE_TOPIX:
		return kabus.FutureCodeTOPIX
	case kabuspb.FutureCode_FUTURE_CODE_TOPIX_MINI:
		return kabus.FutureCodeTOPIXMini
	case kabuspb.FutureCode_FUTURE_CODE_MOTHERS:
		return kabus.FutureCodeMOTHERS
	case kabuspb.FutureCode_FUTURE_CODE_JPX400:
		return kabus.FutureCodeJPX400
	case kabuspb.FutureCode_FUTURE_CODE_DOW:
		return kabus.FutureCodeDOW
	case kabuspb.FutureCode_FUTURE_CODE_VI:
		return kabus.FutureCodeVI
	case kabuspb.FutureCode_FUTURE_CODE_CORE30:
		return kabus.FutureCodeCore30
	case kabuspb.FutureCode_FUTURE_CODE_REIT:
		return kabus.FutureCodeREIT
	}
	return kabus.FutureCodeUnspecified
}

func toYmNum(timestamp *timestamppb.Timestamp) kabus.YmNUM {
	t := timestamp.AsTime().In(time.Local)
	if t.IsZero() {
		return kabus.YmNUMToday
	}
	return kabus.NewYmNUM(t)
}

func toPutOrCall(callPut kabuspb.CallPut) kabus.PutOrCall {
	switch callPut {
	case kabuspb.CallPut_CALL_PUT_CALL:
		return kabus.PutOrCallCall
	case kabuspb.CallPut_CALL_PUT_PUT:
		return kabus.PutOrCallPut
	}
	return kabus.PutOrCallUnspecified
}
