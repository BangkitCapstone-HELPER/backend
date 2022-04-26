package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/BangkitCapstone-HELPER/backend/internal/app/config"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/controllers"
	e "github.com/BangkitCapstone-HELPER/backend/internal/app/error"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/middlewares"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/repo"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/routes"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/services"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	e.Module,
	config.Module,
	lib.Module,
	repo.Module,
	services.Module,
	controllers.Module,
	routes.Module,
	middlewares.Module,
	// repository.Module,
	fx.Invoke(bootstrap),
)

var hook fx.Hook

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.HTTPServer,
	route routes.Routes,
	// logger *log.Logger,
	conf config.HTTPConfig,
	middleware middlewares.Middlewares,
	// database lib.Database,
) {
	hook = fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting Application")

			go func() {
				middleware.Setup()
				route.Setup()

				if err := handler.Engine().Start(conf.ListenAddr()); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						fmt.Println("Shutting down the Application")
					} else {
						fmt.Println("Error to Start Application: %v", err)
					}
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping Application")

			shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			_ = handler.Engine().Shutdown(shutdownCtx)
			// db.Close()
			return nil
		},
	}
	lifecycle.Append(hook)
}
