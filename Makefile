upES:
	docker-compose -f deployments/docker-compose.yaml up -d
downES:
	docker-compose -f deployments/docker-compose.yaml down
build:
	go build -o cmd/app/main cmd/app/main.go
generate:
	swag init -g ./cmd/app/main.go -o ./docs
run: generate build
	./cmd/app/main