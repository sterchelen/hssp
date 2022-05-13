.PHONY: build
build: lint
	go build

.PHONY: install
install:
	go install

.PHONY: doc
doc:
	godoc -http=:6060

.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml

.PHONY: clear
clean:
	go clean
