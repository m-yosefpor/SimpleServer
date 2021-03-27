package main

import (
	"github.com/m-yosefpor/simpleserver/internal/cli"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	c := cli.InitCli()

	if err := c.ReadConfig("config.example.yaml"); err != nil {
		log.Fatal(err)
	}

	if err := c.C.Execute(); err != nil {
		log.Fatal(err)
	}
}
