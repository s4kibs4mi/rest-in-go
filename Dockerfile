FROM golang
COPY . /app
WORKDIR /app/rest-in-go
ENTRYPOINT go run main.go
EXPOSE 8009
