FROM golang:latest

WORKDIR /application

COPY go.mod go.sum ./

RUN go mod download


ADD . /application

#RUN go get /application/...
RUN go build -o /application/build/aviasales .

CMD ["/appication/build/aviasales"]