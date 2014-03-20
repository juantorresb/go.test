GOROOT := /usr/local/go
GOPATH := /var/lib/jenkins/workspace/go/project2


myhostname := $(shell hostname)
#ifeq (${myhostname}, laptop)
#    GOPATH := /home/sonia/go/project2
#else ifeq (${myhostname}, testmachine)
#    GOPATH := /home/u1234/go/project2
#    GOROOT := /usr/local/go
#endif
 
build: build-stamp
build-stamp: file1.go file2.go file3.go
    # always format code
    GOPATH=$(GOPATH) go fmt $^
    # binary
    GOPATH=$(GOPATH) go build -o project2 -v $^
    # docs
    markdown README.mkd > README.html
    help2man --no-info --include=help2man.roff --name "Project2" ./project2 > project2.roff
    man -Tps -l project2.roff > project2-man.ps
    ps2pdf project2-man.ps project2-man.pdf
    # mark as done
    touch $@pwd


