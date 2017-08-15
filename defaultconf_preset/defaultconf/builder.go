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
	config *Config
}

func NewBuilder() *Builder {
	log.Println("NewBuilder(): Creating new Builder...")
	var b *Builder = new(Builder)
	log.Println("NewBuilder(): Initializing new Config...")
	b.config = new(Config)
	log.Println("NewBuilder(): Assigning value Create = true")
	b.config.Create = true
	log.Println("NewBuilder(): Assigning value Number = 42")
	b.config.Number = 42
	log.Println("NewBuilder(): Created new Builder. Returning.")
	return b
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

	log.Println("Prepare(): Prepared. Returning.")
	return nil, nil
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	log.Println("Hello I'm a custom builder")
	return nil, nil
}

func (b *Builder) Cancel() {
}
