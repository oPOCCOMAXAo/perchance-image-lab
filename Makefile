build:
	go build -o bin/lab cmd/lab/main.go

lint:
	golangci-lint-v2 run
