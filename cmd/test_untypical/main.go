package main

import (
	"flag"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"test_untypical/config"
	"test_untypical/internal/test_untypical/server"
	"test_untypical/internal/test_untypical/server/handlers"
	mapStorage "test_untypical/internal/test_untypical/service/map_service"
)

func main() {

	configPath := new(string)
	flag.StringVar(configPath, "config-path", "config/config.yaml", "path to yaml config")
	flag.Parse()

	cnfFile, err := os.Open(*configPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "err with os.Open config"))
	}

	cnf := config.Config{}
	if err := yaml.NewDecoder(cnfFile).Decode(&cnf); err != nil {
		log.Fatal(errors.Wrap(err, "err with Decode config"))
	}
	srv := mapStorage.NewMapService()
	handls := handlers.NewHandlers(srv)

	log.Println("Start server")
	server.StartServer(handls, cnf.ServerPort)
}
