package service

import (
	"net/http"

	"github.com/takama/bit"

	"goinaction/k8sapp/config"
	"goinaction/k8sapp/handlers"
	"goinaction/k8sapp/logger"
	stdlog "goinaction/k8sapp/logger/logrus"
	"goinaction/k8sapp/version"
)

func Setup(cfg *config.Config) (r bit.Router, log logger.Logger, err error) {
	log = stdlog.New(&logger.Config{
		Level: cfg.LogLevel,
		Time:  true,
		UTC:   true,
	})

	log.Info("Version:", version.RELEASE)
	log.Warnf("%s log level is used", logger.LevelDebug.String())
	log.Infof("Service %s listened on %s:%d", config.SERVICENAME, cfg.LocalHost, cfg.LocalPort)

	h := handlers.New(log, cfg)

	r = bit.NewRouter()

	r.SetupNotFoundHandler(h.Base(notFound))

	r.SetupMiddleware(h.Base)
	r.GET("/", h.Root)
	r.GET("/healthz", h.Health)
	r.GET("/readyz", h.Ready)
	r.GET("/info", h.Info)

	return
}

func notFound(c bit.Control) {
	c.Code(http.StatusNotFound)
	c.Body("Method not found for " + c.Request().URL.Path)
}
