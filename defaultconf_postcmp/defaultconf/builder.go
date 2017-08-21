package defaultconf

import (
	"log"

	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	ctx                 interpolate.Context

	// Create: a boolean stating, for example, if we should create
	// something or not
	Create bool `mapstructure:"create"`
	// Number: let's say the number of things to create
	Number int `mapstructure:"number"`
}

type Builder struct {
	config Config
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	log.Println("Prepare(): Starting...")

	log.Println("Prepare(): Current config: ", b.config)
	log.Println("Prepare(): Create config Parameter: ", b.config.Create)
	log.Println("Prepare(): Number config Parameter: ", b.config.Number)
	log.Println("Prepare(): Calling config.Decode()...")
	err := config.Decode(&b.config, &config.DecodeOpts{Interpolate: false}, raws...)
	if err != nil {
		return nil, err
	}
	log.Println("Prepare(): Called config.Decode()")
	log.Println("Prepare(): Current config: ", b.config)
	log.Println("Prepare(): Create config Parameter: ", b.config.Create)
	log.Println("Prepare(): Number config Parameter: ", b.config.Number)

	log.Println("Prepare(): Checking for default values:")
        // Check
        if b.config.Create == false { // false is the default value of a bool
          log.Println("-- Prepare(): Create is false - assigning true")
          b.config.Create = true
        }
        // Check
        if b.config.Number == 0 { // 0 is the default value of an int
          log.Println("-- Prepare(): Number is 0 - assigning 42")
          b.config.Number = 42
        }

	log.Println("Prepare(): Current config: ", b.config)
	log.Println("Prepare(): Create config Parameter: ", b.config.Create)
	log.Println("Prepare(): Number config Parameter: ", b.config.Number)
	log.Println("Prepare(): Prepared. Returning.")
	return nil, nil
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	log.Println("Hello I'm a custom builder")
	return nil, nil
}

func (b *Builder) Cancel() {
}
