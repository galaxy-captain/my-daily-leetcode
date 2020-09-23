.PHONY: test
.SILENT: test

FILES := $(shell go list ./... | grep -v helper)

test:
	go test ${FILES} -count=1