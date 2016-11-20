export GOPATH = $(PWD)
export GOBIN  = $(PWD)/bin

EXE = granite

all: parser bin

clean:
	rm -f bin/$(EXE)
	rm -f parser/$(EXE).go
	rm -rf pkg

parser: parser/$(EXE).go

parser/$(EXE).go:
	go generate ./parser

bin: bin/$(EXE)

bin/$(EXE): parser
	go build -o bin/$(EXE)

get:
	go get
