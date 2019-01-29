FROM golang:latest

RUN mkdir -p /home/user/go/src/ah8ad3/tamgo

WORKDIR /home/user/go/src/ah8ad3/tamgo

ADD . .

RUN export GOPATH=/home/user/go && go get ./... && go build -o main

EXPOSE 8000

CMD ["/home/user/go/src/ah8ad3/tamgo/main"]