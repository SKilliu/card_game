package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Init(log *logrus.Entry) {
	// Set server logger
	setLogger(log)

	loadConfigFromEnvs()
	logger.Info("Server succesfully inited")
}

func Start() {
	logger.Infof("Listening on port %s:%d", configuration.Host, configuration.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.Host, configuration.Port), nil)
	if err != nil {
		logger.Fatal(err)
	}
}
