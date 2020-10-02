package factory

import (
	"github.com/go-chi/chi"

	grpc "go.kicksware.com/api/beta-service/api/gRPC"
	"go.kicksware.com/api/beta-service/api/rest"
	"go.kicksware.com/api/beta-service/core/service"
	"go.kicksware.com/api/beta-service/env"
)

func ProvideRESTGatewayHandler(service service.BetaService, auth service.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideGRPCGatewayHandler(service service.BetaService, auth service.AuthService, config env.ServiceConfig) *grpc.Handler {
	return grpc.NewHandler(
		service,
		auth,
		config.Common,
	)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}