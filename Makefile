include .env
PROJECTNAME=$(shell basename "$(PWD)")

SHELL = /bin/bash
GO-VER = go1.13


# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

OSFAMILY=$(shell uname)
ifeq ($(OSFAMILY),Darwin)
OSFAMILY=darwin
else
OSFAMILY=linux
endif

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID=/tmp/.$(PROJECTNAME).pid

# .PHONY: security-user-name
# security-user-name:
# ifndef SECURITY_USER_NAME
# 	$(error variable SECURITY_USER_NAME not defined)
# endif

.PHONY: deps-go-binary
deps-go-binary:
	echo "Expect: $(GO-VER)" && \
		echo "Actual: $$($(GO) version)" && \
	 	go version | grep $(GO-VER) > /dev/null

clean:
	@(MAKEFILE) go-clean

go-clean:
	@echo "  >  Cleaning build cache"
	GOBIN=$(GOBIN) go clean
	rm -rf $(GOBIN)

build:
	@echo "  >  Building binary..."
	go build -o $(GOBIN)/$(PROJECTNAME) .

serve:
		@echo "  >  starting serve"
		$(GOBIN)/$(PROJECTNAME) serve
