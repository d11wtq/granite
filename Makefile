EXE = bijou

all: parser bin

clean:
	rm -f bin/$(EXE)
	rm -f parser/$(EXE).go

parser: parser/$(EXE).go

parser/$(EXE).go:
	go generate ./parser

bin: bin/$(EXE)

bin/$(EXE): parser
	go build -o bin/$(EXE)
