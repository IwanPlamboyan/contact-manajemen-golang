//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/IwanPlamboyan/contact-manajemen-golang/app"
	"github.com/IwanPlamboyan/contact-manajemen-golang/config"
	"github.com/IwanPlamboyan/contact-manajemen-golang/controller"
	"github.com/IwanPlamboyan/contact-manajemen-golang/middleware"
	"github.com/IwanPlamboyan/contact-manajemen-golang/repository"
	"github.com/IwanPlamboyan/contact-manajemen-golang/service"
	"github.com/google/wire"

	"github.com/go-playground/validator"
)

var DatabaseSet = wire.NewSet (
	config.ProvideDatabaseConfig,
	app.Connect,
)

var authSet = wire.NewSet(
	repository.NewRefreshTokenRepository,
	wire.Bind(new(repository.RefreshTokenRepository), new(*repository.RefreshTokenRepositoryImpl)),
	service.NewAuthService,
	controller.NewAuthController,
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewUserService,
	controller.NewUserController,
)

var contactSet = wire.NewSet(
	repository.NewContactRepository,
	wire.Bind(new(repository.ContactRepository), new(*repository.ContactRepositoryImpl)),
	service.NewContactsService,
	controller.NewContactController,
)

var addressSet = wire.NewSet(
	repository.NewAddressRepository,
	wire.Bind(new(repository.AddressRepository), new(*repository.AddressRepositoryImpl)),
	service.NewAddressService,
	controller.NewAddressController,
)

func InitializedServer() (*http.Server, error) {
	wire.Build(
		config.LoadConfig,
		config.ProvideJWTUtil,
		middleware.NewJWTMiddleware,
		DatabaseSet,
		validator.New,
		authSet,
		userSet,
		addressSet,
		contactSet,
		app.NewRouter,
		NewServer,
	)
	return nil, nil
}