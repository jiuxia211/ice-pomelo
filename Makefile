DIR = $(shell pwd)
CMD = $(DIR)/cmd
OUTPUT_PATH = $(DIR)/output


.PHONY: build
build:
	mkdir -p output
	cd $(CMD)/api && sh build.sh
	cd $(CMD)/user && sh build.sh
	cd $(CMD)/tiny_id && sh build.sh