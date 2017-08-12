package main

import (
	"github.com/Horgix/packer-builder-defaultconf/defaultconf"
	"github.com/hashicorp/packer/packer/plugin"
	"log"
)

func main() {
	log.Println("Starting...")
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterBuilder(defaultconf.NewBuilder())
	server.Serve()
}
