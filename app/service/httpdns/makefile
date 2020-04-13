$(eval -include ../../../build/makefile.tars)
TARS_RELEASE_PATH := $(TARS_RELEASE_ROOT)/service/httpdns

APP       := zim
TARGET    := HttpDns
MFLAGS    :=
DFLAGS    :=
CONFIG    := service.conf
STRIP_FLAG:= N
#-------------------------------------------------------------------------------
#fix cgo compile error
export LC_ALL   = en_US.UTF-8
export LANG     = en_US.UTF-8
#-------------------------------------------------------------------------------
GOPATH ?= $(shell go env GOPATH)
GOROOT ?= $(shell go env GOROOT)
GOFMT ?= gofmt -s -w
GO      = ${GOROOT}/bin/go
TAGS = pprof
LDFLAGS += -X "gitbus.com/exlab/zim-ms/app/service/httpdns/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "gitbus.com/exlab/zim-ms/app/service/httpdns/version.GitHash=$(shell git rev-parse HEAD)"
#-------------------------------------------------------------------------------
libpath=${subst :, ,$(GOPATH)}
TARS2GO     := $(firstword $(subst :, , $(shell go env GOPATH)))/bin/tars2go
J2GO_FLAG   := -outdir=proto/gen
#-------------------------------------------------------------------------------
TARS_SRC     := $(wildcard proto/*.tars)
GO_SRC      := $(wildcard *.go)
GOFILES := $(shell find . -name "*.go" -type f)
#----------------------------------------------------------------------------------
copyfile = if test -z "$(APP)" || test -z "$(TARGET)"; then \
               echo "['APP' or 'TARGET' option is empty.]"; exit 1; \
           	else \
		       	if test ! -d $(2); then \
              		echo "[No such dir:$(2), now we create it.]";\
    				mkdir -p $(2);\
				fi; \
         		echo "[Copy file $(1) -> $(2)]"; \
         		cp -v $(1) $(2); \
			fi;

ALL: $(TARGET)
#----------------------------------------------------------------------------------
$(TARGET): gen-tars $(GO_SRC)
	$(GO) build -ldflags '$(LDFLAGS)' $(GO_BUILD_FLAG) -o $@
#----------------------------------------------------------------------------------
pprof: gen-tars $(GO_SRC)
	$(GO) build -ldflags '$(LDFLAGS)' -tags '$(TAGS)' $(GO_BUILD_FLAG) -o $(TARGET)
#----------------------------------------------------------------------------------
serve: gen-tars $(GO_SRC)
	$(GO) run -ldflags '$(LDFLAGS)' $(GO_BUILD_FLAG) main.go --config=$(CONFIG)
#----------------------------------------------------------------------------------
serve-pprof: gen-tars $(GO_SRC)
	$(GO) run -ldflags '$(LDFLAGS)' -tags '$(TAGS)' $(GO_BUILD_FLAG) main.go --config=$(CONFIG)
#----------------------------------------------------------------------------------
client: gen-tars $(GO_SRC)
	$(GO) run -ldflags '$(LDFLAGS)' internal/client/client.go --config=$(CONFIG)
#----------------------------------------------------------------------------------
.PHONY: gen-tars
gen-tars: $(TARS_SRC)
	$(TARS2GO) $(J2GO_FLAG) $(TARS_SRC)
#----------------------------------------------------------------------------------
.PHONY: linux-amd64
linux-amd64: gen-tars $(GO_SRC)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' $(GO_BUILD_FLAG) -o $(TARGET)
#----------------------------------------------------------------------------------
.PHONY: release
release: linux-amd64 $(CONFIG)
	@if [ -d $(TARGET)_tmp_dir ]; then \
		echo "dir has exist:$(TARGET)_tmp_dir, abort."; \
		exit 1; \
	else \
		mkdir -p $(TARGET)_tmp_dir/$(TARGET) $(TARS_RELEASE_PATH);\
		cp -rf $(TARGET) $(CONFIG) $(TARGET)_tmp_dir/$(TARGET)/; \
		cd $(TARGET)_tmp_dir; tar --exclude=".svn" --exclude="_svn" -czvf $(TARGET).tgz $(TARGET)/; cd -; \
		mv $(TARGET)_tmp_dir/$(TARGET).tgz $(TARS_RELEASE_PATH); \
		rm -rf $(TARGET)_tmp_dir; \
		echo "tar cvfz $(TARGET).tgz ..."; \
	fi

ifneq ($(TARS_SRC),)

SERVER_NAME := $(TARGET)

endif
#----------------------------------------------------------------------------------
.PHONY: test
test: fmt
	go test ./...
#----------------------------------------------------------------------------------
.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
#----------------------------------------------------------------------------------
clean:
	rm -vf $(DEPEND_TARS_OBJ) $(INVOKE_DEPEND_TARS_OBJ) $(LOCAL_OBJ) $(TARGET) $(TARGETS) $(DEP_FILE) ${CLEANFILE} .*.d.tmp gmon.out
#----------------------------------------------------------------------------------
cleanall:
	rm -vf $(DEPEND_TARS_H) $(DEPEND_TARS_CPP) $(DEPEND_TARS_OBJ) $(LOCAL_OBJ) $(HCE_H) $(HCE_CPP) $(TARGET) $(TARGETS) $(DEP_FILE) ${CLEANFILE} *.o .*.d.tmp .*.d gmon.out
