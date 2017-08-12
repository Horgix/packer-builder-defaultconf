GOAPP		= github.com/Horgix/packer-builder-defaultconf/defaultconf_1
CWD		= `pwd`

SOURCES		= main.go defaultconf/

build::
	go build -o packer-builder-defaultconf_1

fmt::
	gofmt -w ${SOURCES}

test::
	PACKER_LOG=true packer build test_packer.json
