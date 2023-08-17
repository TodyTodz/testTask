//Package for start the microservice
package app

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"testTask/internal/handlers"
	"testTask/pkg/app_utils"
	"testTask/pkg/handler_helpers"

	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"

	. "github.com/NGRsoftlab/ngr-logging"
)

//Run - start HTTP server and wait for shutdown signal
func Run(httpPath string){
	// root context of service, which Done with signals for stop service
	rootCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	// create error group instead of wait group for getting error if app ruins
	eg, ctx := errgroup.WithContext(rootCtx)

	// parse http config file
	httpSettings, err := handler_helpers.NewHttpConfig(httpPath)
	if err != nil {
		Logger.Fatal("HTTP httpConfig init failed. Reason:", err.Error())
	}

	routerHTTP := handlers.Router()
	c := cors.AllowAll()

	// sorry, i fixed this little char later than necessary. You can check it in commits history.
	srvHTTP := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", httpSettings.Host(), httpSettings.PortHttp()),
		Handler: c.Handler(routerHTTP),
	}

	// run http server
	eg.Go(func() error {
		return srvHTTP.ListenAndServe()
	})

	// waiting for shutdown signals
	if err = app_utils.WaitSignals(ctx); err != nil {
		Logger.Error("Failed to wait OS signal. Reason:", err.Error())
	}

	Logger.Info("Stopping the service...")

	// stop http server
	err = srvHTTP.Shutdown(context.Background())
	if err != nil {
		Logger.Error("ERROR of shutting down http", err.Error())
	}

	Logger.Info("Finished")
}
