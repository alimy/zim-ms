#fix cgo compile error
export LC_ALL   = en_US.UTF-8
export LANG     = en_US.UTF-8
#-------------------------------------------------------------------------------
GOFMT ?= gofmt -s -w
GO      = ${GOROOT}/bin/go
GOFILES := $(shell find . -name "*.go" -type f)
#----------------------------------------------------------------------------------
.PHONY: test
test: fmt
	go test ./...
#----------------------------------------------------------------------------------
.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
