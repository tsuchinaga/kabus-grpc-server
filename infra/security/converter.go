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

func fromCurrentPriceChangeStatus(status kabus.CurrentPriceChangeStatus) string {
	switch status {
	case kabus.CurrentPriceChangeStatusUnspecified:
		return "0000"
	case kabus.CurrentPriceChangeStatusNoChange:
		return "0056"
	case kabus.CurrentPriceChangeStatusUp:
		return "0057"
	case kabus.CurrentPriceChangeStatusDown:
		return "0058"
	case kabus.CurrentPriceChangeStatusOpenPriceAfterBreak:
		return "0059"
	case kabus.CurrentPriceChangeStatusTradingSessionClose:
		return "0060"
	case kabus.CurrentPriceChangeStatusClose:
		return "0061"
	case kabus.CurrentPriceChangeStatusBreakClose:
		return "0062"
	case kabus.CurrentPriceChangeStatusDownClose:
		return "0063"
	case kabus.CurrentPriceChangeStatusTarnOverClose:
		return "0064"
	case kabus.CurrentPriceChangeStatusSpecialQuoteClose:
		return "0066"
	case kabus.CurrentPriceChangeStatusReservationClose:
		return "0067"
	case kabus.CurrentPriceChangeStatusStopClose:
		return "0068"
	case kabus.CurrentPriceChangeCircuitBreakerClose:
		return "0069"
	case kabus.CurrentPriceChangeDynamicCircuitBreakerClose:
		return "0431"
	}
	return ""
}

func fromCurrentPriceStatus(status kabus.CurrentPriceStatus) int32 {
	switch status {
	case kabus.CurrentPriceStatusCurrentPrice:
		return 1
	case kabus.CurrentPriceStatusNoContinuousTicks:
		return 2
	case kabus.CurrentPriceStatusItayose:
		return 3
	case kabus.CurrentPriceStatusSystemError:
		return 4
	case kabus.CurrentPriceStatusPause:
		return 5
	case kabus.CurrentPriceStatusStopTrading:
		return 6
	case kabus.CurrentPriceStatusRestart:
		return 7
	case kabus.CurrentPriceStatusClosePrice:
		return 8
	case kabus.CurrentPriceStatusSystemStop:
		return 9
	case kabus.CurrentPriceStatusRoughQuote:
		return 10
	case kabus.CurrentPriceStatusReference:
		return 11
	case kabus.CurrentPriceStatusInCircuitBreak:
		return 12
	case kabus.CurrentPriceStatusRestoration:
		return 13
	case kabus.CurrentPriceStatusReleaseCircuitBreak:
		return 14
	case kabus.CurrentPriceStatusReleasePause:
		return 15
	case kabus.CurrentPriceStatusInReservation:
		return 16
	case kabus.CurrentPriceStatusReleaseReservation:
		return 17
	case kabus.CurrentPriceStatusFileError:
		return 18
	case kabus.CurrentPriceStatusReleaseFileError:
		return 19
	case kabus.CurrentPriceStatusSpreadStrategy:
		return 20
	case kabus.CurrentPriceStatusInDynamicCircuitBreak:
		return 21
	case kabus.CurrentPriceStatusReleaseDynamicCircuitBreak:
		return 22
	case kabus.CurrentPriceStatusContractedInItayose:
		return 23
	}
	return 0
}

func fromBidAskSign(sign kabus.BidAskSign) string {
	switch sign {
	case kabus.BidAskSignNoEffect:
		return "0000"
	case kabus.BidAskSignGeneral:
		return "0101"
	case kabus.BidAskSignSpecial:
		return "0102"
	case kabus.BidAskSignAttention:
		return "0103"
	case kabus.BidAskSignBeforeOpen:
		return "0107"
	case kabus.BidAskSignSpecialBeforeStop:
		return "0108"
	case kabus.BidAskSignAfterClose:
		return "0109"
	case kabus.BidAskSignNotExistsContractPoint:
		return "0116"
	case kabus.BidAskSignExistsContractPoint:
		return "0117"
	case kabus.BidAskSignContinuous:
		return "0118"
	case kabus.BidAskSignContinuousBeforeStop:
		return "0119"
	case kabus.BidAskSignMoving:
		return "0120"
	}
	return ""
}

func fromFirstBoardSign(firstBoardSign kabus.FirstBoardSign) *kabuspb.FirstQuote {
	return &kabuspb.FirstQuote{
		Time:     timestamppb.New(firstBoardSign.Time),
		Sign:     fromBidAskSign(firstBoardSign.Sign),
		Price:    firstBoardSign.Price,
		Quantity: firstBoardSign.Qty,
	}
}

func fromBoardSign(boardSign kabus.BoardSign) *kabuspb.Quote {
	return &kabuspb.Quote{
		Price:    boardSign.Price,
		Quantity: boardSign.Qty,
	}
}

func fromPutOrCallNum(putOrCall kabus.PutOrCallNum) kabuspb.CallPut {
	switch putOrCall {
	case kabus.PutOrCallNumCall:
		return kabuspb.CallPut_CALL_PUT_CALL
	case kabus.PutOrCallNumPut:
		return kabuspb.CallPut_CALL_PUT_PUT
	}
	return kabuspb.CallPut_CALL_PUT_UNSPECIFIED
}

func fromUnderlyer(underlyer kabus.Underlyer) string {
	switch underlyer {
	case kabus.UnderlyerNK225:
		return "NK225"
	case kabus.UnderlyerNK300:
		return "NK300"
	case kabus.UnderlyerMOTHERS:
		return "MOTHERS"
	case kabus.UnderlyerJPX400:
		return "JPX400"
	case kabus.UnderlyerTOPIX:
		return "TOPIX"
	case kabus.UnderlyerNKVI:
		return "NKVI"
	case kabus.UnderlyerDJIA:
		return "DJIA"
	case kabus.UnderlyerTSEREITINDEX:
		return "TSEREITINDEX"
	case kabus.UnderlyerTOPIXCORE30:
		return "TOPIXCORE30"
	}
	return ""
}

func fromPriceRangeGroup(priceRangeGroup kabus.PriceRangeGroup) string {
	switch priceRangeGroup {
	case kabus.PriceRangeGroup10000:
		return "10000"
	case kabus.PriceRangeGroup10003:
		return "10003"
	case kabus.PriceRangeGroup10118:
		return "10118"
	case kabus.PriceRangeGroup10119:
		return "10119"
	case kabus.PriceRangeGroup10318:
		return "10318"
	case kabus.PriceRangeGroup10706:
		return "10706"
	case kabus.PriceRangeGroup10718:
		return "10718"
	case kabus.PriceRangeGroup12122:
		return "12122"
	case kabus.PriceRangeGroup14473:
		return "14473"
	case kabus.PriceRangeGroup14515:
		return "14515"
	case kabus.PriceRangeGroup15411:
		return "15411"
	case kabus.PriceRangeGroup15569:
		return "15569"
	case kabus.PriceRangeGroup17163:
		return "17163"
	}
	return ""
}
