.PHONY: help
help:
	@cat Makefile

.PHONY: gicom
gicom:
	go build -o gitcomicome 
	./gicom --version