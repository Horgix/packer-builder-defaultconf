package main

import (
	"github.com/Horgix/packer-builder-defaultconf/defaultconf_postcmp/defaultconf"
	"github.com/hashicorp/packer/packer/plugin"
	"log"
)

func main() {
	log.Println("Starting...")
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterBuilder(new(defaultconf.Builder))
	server.Serve()
}
