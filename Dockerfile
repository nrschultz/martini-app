FROM golang:1.4.2

RUN go get github.com/go-martini/martini
ADD . /martini
WORKDIR /martini
EXPOSE 3000
CMD ["go", "run", "server.go"]

