package security

import (
	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

func toExchange(exchange kabuspb.Exchange) kabus.Exchange {
	switch exchange {
	case kabuspb.Exchange_TOUSHOU:
		return kabus.ExchangeToushou
	case kabuspb.Exchange_MEISHOU:
		return kabus.ExchangeMeishou
	case kabuspb.Exchange_FUKUSHOU:
		return kabus.ExchangeFukushou
	case kabuspb.Exchange_SATSUSHOU:
		return kabus.ExchangeSatsushou
	case kabuspb.Exchange_ALL_SESSION:
		return kabus.ExchangeAll
	case kabuspb.Exchange_DAY_SESSION:
		return kabus.ExchangeDaytime
	case kabuspb.Exchange_NIGHT_SESSION:
		return kabus.ExchangeEvening
	default:
		return kabus.ExchangeUnspecified
	}
}

func fromExchange(exchange kabus.Exchange) kabuspb.Exchange {
	switch exchange {
	case kabus.ExchangeToushou:
		return kabuspb.Exchange_TOUSHOU
	case kabus.ExchangeMeishou:
		return kabuspb.Exchange_MEISHOU
	case kabus.ExchangeFukushou:
		return kabuspb.Exchange_FUKUSHOU
	case kabus.ExchangeSatsushou:
		return kabuspb.Exchange_SATSUSHOU
	case kabus.ExchangeAll:
		return kabuspb.Exchange_ALL_SESSION
	case kabus.ExchangeDaytime:
		return kabuspb.Exchange_DAY_SESSION
	case kabus.ExchangeEvening:
		return kabuspb.Exchange_NIGHT_SESSION
	default:
		return kabuspb.Exchange_UNSPECIFIED
	}
}
