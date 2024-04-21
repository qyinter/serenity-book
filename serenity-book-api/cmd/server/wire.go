//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"serenity-book-api/internal/handler"
	"serenity-book-api/internal/repository"
	"serenity-book-api/internal/server"
	"serenity-book-api/internal/service"
	"serenity-book-api/pkg/log"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewUserRepository,
	repository.NewSystemRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewSystemService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewSystemHandler,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
