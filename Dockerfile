FROM golang:1.17-alpine

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod ./
RUN go mod vendor
RUN go mod download

RUN mkdir -p ./docker/postgres/data
RUN mkdir -p ./docker/pgadmin/data

COPY ./ ./
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]