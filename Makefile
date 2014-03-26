MyGOPATH = /home/go/devel
GOPATH=$GOPATH:/var/lib/jenkins/workspace/go/

myhostname := $(shell hostname)
  
build:
	go build go.test/first/
	go build go.test/log

install:
	go install go.test/first/
	go install go.test/log/

test:
	go test go.test/first/
	go test go.test/log/

coverage:
	go test -coverprofile=coverage.out -covermode count go.test/first
	$(MyGOPATH)/bin/gocover-cobertura < coverage.out >  coverage.xml