package di

import (
	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/infra"
	"gitlab.com/tsuchinaga/kabus-grpc-server/infra/security"
	"gitlab.com/tsuchinaga/kabus-grpc-server/infra/stores"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/services"
)

func InjectedServer() kabuspb.KabusServiceServer {
	setting := infra.GetSetting()
	return server.NewServer(
		security.NewSecurity(
			kabus.NewRESTClient(setting.IsProduction())),
		services.NewTokenService(
			stores.GetTokenStore(),
			security.NewSecurity(
				kabus.NewRESTClient(setting.IsProduction())),
			infra.NewClock(),
			setting),
		services.NewRegisterSymbolService(
			stores.GetRegisterSymbolStore()))
}
