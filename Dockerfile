FROM golang
ADD . /go/src/rest-in-go/
RUN go install /go/src/rest-in-go/
ENTRYPOINT /go/bin/rest-in-go
EXPOSE 8009
