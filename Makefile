GOAPP		= github.com/Horgix/packer-builder-defaultconf
CWD		= `pwd`

SOURCES		= main.go defaultconf/

build::
	go build -o `basename ${GOAPP}`

fmt::
	gofmt -w ${SOURCES}

test::
	PACKER_LOG=true packer build test_packer.json
