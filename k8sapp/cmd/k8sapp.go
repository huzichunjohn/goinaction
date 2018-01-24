package main

import (
	"fmt"
	"log"

	"goinaction/k8sapp/config"
	"goinaction/k8sapp/service"
	"goinaction/k8sapp/system"
)

func main() {
	cfg := new(config.Config)
	if err := cfg.Load(config.SERVICENAME); err != nil {
		log.Fatal(err)
	}

	router, logger, err := service.Setup(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go router.Listen(fmt.Sprintf("%s:%d", cfg.LocalHost, cfg.LocalPort))

	signals := system.NewSignals()
	if err := signals.Wait(logger, new(system.Handling)); err != nil {
		logger.Fatal(err)
	}
}
