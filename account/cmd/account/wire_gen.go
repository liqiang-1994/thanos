// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"account/internal/biz"
	"account/internal/conf"
	"account/internal/data"
	"account/internal/server"
	"account/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, auth *conf.Auth, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserCase(auth, userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, auth, userService, logger, tracerProvider)
	httpServer := server.NewHTTPServer(confServer, auth, userService, logger)
	registrar := data.NewRegistry(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
