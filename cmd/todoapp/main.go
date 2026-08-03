package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/LootKeeper/todoapp-go/internal/core/logger"
	core_http_middleware "github.com/LootKeeper/todoapp-go/internal/core/transport/http/middleware"
	core_http_server "github.com/LootKeeper/todoapp-go/internal/core/transport/http/server"
	users_transport_http_v1 "github.com/LootKeeper/todoapp-go/internal/features/users/transport/http/v1"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("failed to ini app logger: ", err)
		os.Exit(1)
	}

	defer logger.Close()

	logger.Debug("Starting app...")

	usersTransportHTTP := users_transport_http_v1.NewUsersHTTPHandler(nil)
	usersRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersRoutes...)

	middleware := []core_http_middleware.Middleware{
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	}

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		middleware...,
	)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("Error during httpServer running: ", zap.Error(err))
	}
}
