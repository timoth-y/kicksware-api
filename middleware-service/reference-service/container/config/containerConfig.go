package config

import (
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/container"

	"reference-service/container/factory"
	"reference-service/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).
		BindSingleton(factory.ProvideDataService).
		BindSingleton(factory.ProvideGatewayHandler).
		BindTransient(factory.ProvideEndpointRouter).
		BindTransient(factory.ProvideServer)
}
