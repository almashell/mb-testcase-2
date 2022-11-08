FROM golang:1.13-stretch AS lang

WORKDIR /home/mb_api
COPY . .
RUN go build -o bin/mb_api ./cmd/server/main.go

EXPOSE 3001

CMD ./bin/mb_api