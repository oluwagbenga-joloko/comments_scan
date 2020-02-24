BUILDPATH=$(CURDIR)
GO=$(shell which go)

EXENAME=main

test:
	@$(GO) test -v -cover

build: 
	@$(GO) build -o $(EXENAME)

run:
	./$(EXENAME)

all: test build run
