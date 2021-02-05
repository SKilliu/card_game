package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

func Init(log *logrus.Entry) {
	// Set server logger
	setLogger(log)

	loadConfigFromEnvs()
	logger.Info("Server succesfully inited")
}

func Start() error {
	router := NewRouter(logger)

	httpServer := http.Server{
		Addr:           fmt.Sprintf("%s:%d", configuration.Host, configuration.Port),
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Infof("Listening on port %s:%d", configuration.Host, configuration.Port)
	if err := httpServer.ListenAndServe(); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}

	return nil
}
