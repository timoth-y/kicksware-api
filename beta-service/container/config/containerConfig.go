package config

import (
	"github.com/timoth-y/kicksware-api/service-common/container"

	"go.kicksware.com/api/beta-service/container/factory"
	"go.kicksware.com/api/beta-service/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(factory.ProvideDataService).
		BindSingleton(factory.ProvideAuthService).

		BindSingleton(factory.ProvideRESTGatewayHandler).
		BindSingleton(factory.ProvideGRPCGatewayHandler).

		BindTransient(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
