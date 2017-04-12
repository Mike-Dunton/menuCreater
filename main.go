package main

import (
	"github.com/mike-dunton/menuCreater/routes"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.Info("Starting...")
	// Start server on port 8080
	routes.GetMainRouter().Run(":8080")
}
