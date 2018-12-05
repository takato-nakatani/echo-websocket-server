FROM golang:latest

WORKDIR /go/src/github.com/takato-nakatani/ChatEcho/backend

RUN go get github.com/labstack/echo/...
RUN go get -u firebase.google.com/go
RUN go get github.com/joho/godotenv
RUN go get -u google.golang.org/api/option
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/websocket

COPY . /go/src/github.com/takato-nakatani/ChatEcho/backend

CMD ["go", "run", "server.go"]