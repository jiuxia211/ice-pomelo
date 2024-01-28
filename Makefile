

.PHONY: build
build:
	mkdir -p output
	cd cmd/api && sh build.sh
	cd cmd/user && sh build.sh
	cd cmd/tiny_id && sh build.sh