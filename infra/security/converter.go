package security

import (
	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
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
