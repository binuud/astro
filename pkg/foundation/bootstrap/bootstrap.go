package bootstrap

// binu@dronasys.com

import (
	"github.com/dronasys-com/astro/pkg/foundation/config"
	"github.com/dronasys-com/trance-microservices/pkg/foundation/logger"
	"github.com/sirupsen/logrus"
)

// Bootstraps the application
// Init the logger
// Read env and exit if env is not set

func Init() {

	config.NewConfig()
	logger.NewLogger()

	if config.Config.LogVerbosity > 10 {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.TraceLevel)
	} else if config.Config.LogVerbosity > 7 {
		logrus.SetLevel(logrus.TraceLevel)
	} else if config.Config.LogVerbosity > 5 {
		logrus.SetLevel(logrus.InfoLevel)
	} else if config.Config.LogVerbosity > 3 {
		logrus.SetLevel(logrus.WarnLevel)
	}
	logrus.Infof("Config (%v)", config.Config)

}
