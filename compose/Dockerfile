FROM golang:latest

WORKDIR /application

COPY go.mod go.sum ./

RUN go mod download

ENV GOPATH=/application/.go

ADD . /application

#RUN go get /application/...
RUN go build -o /application/build/aviasales /application/src/main.go

RUN ls /application

RUN ls /application/build

CMD ["./build/aviasales"]