FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o cmd/app/main cmd/app/main.go

EXPOSE 8080

CMD [ "/cmd/app/main" ]