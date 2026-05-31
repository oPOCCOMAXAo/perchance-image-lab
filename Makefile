build:
	go build -o bin/lab cmd/lab/main.go

lint:
	golangci-lint-v2 run

format-swagger:
	swag fmt --dir ./

generate-swagger: format-swagger
	swag init \
	--instanceName Full \
	--generalInfo pkg/api/swagger/full.go \
	--outputTypes go \
	--quiet \
	--dir ./

generate: generate-swagger
	go generate ./...

update-dev-deps:
	go get -u github.com/opoccomaxao/gopkg
