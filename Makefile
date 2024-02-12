upES:
	docker-compose -f deployments/docker-compose.yaml up -d
downES:
	docker-compose -f deployments/docker-compose.yaml down
build:
	go build -o cmd/app/main cmd/app/main.go
run: build
	./cmd/app/main