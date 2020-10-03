package factory

import (
	"go.kicksware.com/api/user-service/core/repo"
	"go.kicksware.com/api/user-service/core/service"
	"go.kicksware.com/api/user-service/env"
	"go.kicksware.com/api/user-service/usecase/business"
)

func ProvideDataService(repository repo.UserRepository, remoteRepository repo.RemoteRepository, config env.ServiceConfig) service.UserService {
	return business.NewUserService(repository, remoteRepository, config)
}

func ProvideAuthService(service service.UserService, config env.ServiceConfig) service.AuthService {
	return business.NewAuthServiceJWT(
		service,
		config.Auth,
	)
}

func ProvideMailService(service service.UserService, repo repo.SubscriptionRepository, config env.ServiceConfig) service.MailService {
	return business.NewMailService(
		service,
		repo,
		config.Mail,
		config.FallbackMail,
	)
}

func ProvideInteractService(service service.UserService, likesRepo repo.LikesRepository) service.InteractService {
	return business.NewInteractService(
		service,
		likesRepo,
	)
}