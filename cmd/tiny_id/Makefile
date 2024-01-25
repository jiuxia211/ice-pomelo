include ../common.mk

SERVICE_NAME = tiny_id

.PHONY: server
server:
	sh build.sh
	cd output && sh bootstrap.sh


.PHONY: gen
gen:
	kitex \
	-gen-path ../../kitex_gen \
	-module "$(MODULE)" \
	-service "$(SERVICE_NAME)" \
	-type thrift \
	../../idl/$(SERVICE_NAME).thrift
	go mod tidy