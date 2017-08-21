GOAPP		= github.com/Horgix/packer-builder-defaultconf/defaultconf_postcmp
CWD		= `pwd`

SOURCES		= main.go defaultconf/

build::
	go build -o packer-builder-defaultconf_postcmp

fmt::
	gofmt -w ${SOURCES}

test::
	PACKER_LOG=true packer build test_packer.json
