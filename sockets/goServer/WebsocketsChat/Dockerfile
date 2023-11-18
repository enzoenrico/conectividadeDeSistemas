FROM golang:rc

WORKDIR /app

RUN export GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o main .

EXPOSE 9999

CMD ["./main"]