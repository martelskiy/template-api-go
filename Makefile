BINARY_PATH=${EXEC_DIRECTORY}/bin/api-template-go
EXEC_DIRECTORY=cmd/api-template-go

.PHONY: build
build:
	go build -o ${BINARY_PATH} ${EXEC_DIRECTORY}/main.go

.PHONY: run
run: build
	./${BINARY_PATH}

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: dep
dep:
	go mod download

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: swag-init
swag-init:
	swag init -g ${EXEC_DIRECTORY}/main.go -o ./api/docs
