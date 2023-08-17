package app

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"testTask/pkg/app_utils"
	"testTask/pkg/handler_helpers"

	"testTask/internal/handlers"

	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"

	. "github.com/NGRsoftlab/ngr-logging"
)

func Run(httpPath string){
	rootCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	eg, ctx := errgroup.WithContext(rootCtx)

	httpSettings, err := handler_helpers.NewHttpConfig(httpPath)
	if err != nil {
		Logger.Fatal("HTTP httpConfig init failed. Reason:", err.Error())
	}

	routerHTTP := handlers.Router()
	c := cors.AllowAll()

	srvHTTP := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", httpSettings.Host(), httpSettings.PortHttp()),
		Handler: c.Handler(routerHTTP),
	}

	eg.Go(func() error {
		return srvHTTP.ListenAndServe()
	})

	if err = app_utils.WaitSignals(ctx); err != nil {
		Logger.Error("Failed to wait OS signal. Reason:", err.Error())
	}

	Logger.Info("Stopping the service...")

	err = srvHTTP.Shutdown(context.Background())
	if err != nil {
		Logger.Error("ERROR of shutting down http", err.Error())
	}

	Logger.Info("Finished")
}
