FROM golang
COPY . /go/src/rest-in-go
WORKDIR /go/src/rest-in-go
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 8009
