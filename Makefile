# コマンド一覧
.PHONY: help
help:
	@cat Makefile

.PHONY: gicom
gicom:
	go build .
	./gicom --version