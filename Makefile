.PHONY: help
help:
	@cat Makefile

.PHONY: gicom
gicom:
	go build .
	./gicom --version