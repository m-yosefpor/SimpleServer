package main

import (
	"github.com/m-yosefpor/simpleserver/internal/cli"
	"github.com/m-yosefpor/simpleserver/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	if err := config.Read("config.example.yaml"); err != nil {
		log.Fatal(err)
	}

	c := cli.NewRootCommand()
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}
