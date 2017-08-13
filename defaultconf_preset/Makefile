GOAPP		= github.com/Horgix/packer-builder-defaultconf/defaultconf_preset
CWD		= `pwd`

SOURCES		= main.go defaultconf/

build::
	go build -o packer-builder-defaultconf_preset

fmt::
	gofmt -w ${SOURCES}

test::
	PACKER_LOG=true packer build test_packer.json
