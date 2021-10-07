package di

import (
	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/infra"
	"gitlab.com/tsuchinaga/kabus-grpc-server/infra/security"
	"gitlab.com/tsuchinaga/kabus-grpc-server/infra/stores"
	"gitlab.com/tsuchinaga/kabus-grpc-server/infra/virtual"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/services"
	vs "gitlab.com/tsuchinaga/kabus-virtual-security"
)

func InjectedServer() kabuspb.KabusServiceServer {
	setting := infra.GetSetting()
	return server.NewServer(
		security.NewSecurity(
			kabus.NewRESTClient(setting.IsProduction())),
		virtual.NewSecurity(vs.NewVirtualSecurity()),
		services.NewTokenService(
			stores.GetTokenStore(),
			security.NewSecurity(
				kabus.NewRESTClient(setting.IsProduction())),
			infra.NewClock(),
			setting),
		services.NewRegisterSymbolService(
			stores.GetRegisterSymbolStore()),
		services.NewBoardStreamService(
			stores.GetBoardStreamStore(),
			security.GetBoardWS(setting.IsProduction()),
			virtual.NewSecurity(vs.NewVirtualSecurity())))
}
