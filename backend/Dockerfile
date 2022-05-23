FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go get -u github.com/gin-gonic/gin
RUN go build -o go-book

EXPOSE 8080

CMD ./go-book