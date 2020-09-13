FROM golang:latest

RUN go get -u github.com/gorilla/mux

RUN mkdir /IcecreamShop

ADD . /IcecreamShop

WORKDIR /IcecreamShop

RUN go build -o main .

CMD ["/IcecreamShop/main"]
