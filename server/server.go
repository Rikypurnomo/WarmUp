package server

import (
	"context"
	"net/http"

	"github.com/Rikypurnomo/warmup/config"
	"github.com/Rikypurnomo/warmup/pkg/logger"
	"github.com/Rikypurnomo/warmup/pkg/telemetry/metric"
	"github.com/Rikypurnomo/warmup/pkg/telemetry/trace"
	"github.com/gin-gonic/gin"
)

var server *Kiyora

type Kiyora struct {
	httpServer *http.Server

	traceProviderCloseFn  []trace.CloseFunc
	metricProviderCloseFn []metric.CloseFunc
}

func ShutdownServer(ctx context.Context) {
	server.httpServer.Shutdown(ctx)
}

func StartServer(ctx context.Context, route *gin.Engine) {

	logger.Debug("StartServer: bootstraped routes from etcd")

	httpServer := &http.Server{
		Addr:         config.GetHost(),
		Handler:      route,
		ReadTimeout:  config.ServerReadTimeoutInSecond(),
		WriteTimeout: config.ServerWriteTimeoutInSecond(),
		IdleTimeout:  config.ServerIdleTimeoutInSecond(),
	}
	server = &Kiyora{
		httpServer: httpServer,
	}

	if config.IsEnabledOpenTelemetry() {
		InitGlobalProvider(config.NameServices(), config.OtelToAddr())
	}
	// InitGlobalProvider("kiyora", "http://34.101.227.55:14278/api/traces")

	logger.Debugf("StartServer: starting kiyora on %s", server.httpServer.Addr)

	if err := server.httpServer.ListenAndServe(); err != nil {
		logger.Fatalf("StartServer: starting kiyora failed with %s", err)
	}

	if config.IsEnabledOpenTelemetry() {
		for _, closeFn := range server.traceProviderCloseFn {
			go func() {
				err := closeFn(ctx)
				if err != nil {
					logger.Errorf("StartServer: failed to close trace provider: %v", err)
				}
			}()
		}
	}
}
