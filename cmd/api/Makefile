include ../common.mk

SERVICE_NAME = api

.PHONY: new
new:
	hz new \
	-module $(MODULE)/cmd/$(SERVICE_NAME) \
	-service "$(SERVICE_NAME)" \
	hz update -idl ../../idl/api.thrift


.PHONY: update
update:
	hz update -idl ../../idl/api.thrift

.PHONY: build
build:
	sh build.sh
	
.PHONY: server
server:
	make build
	cd output && sh bootstrap.sh