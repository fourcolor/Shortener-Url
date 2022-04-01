FROM golang
ADD . /usr/local/
WORKDIR /usr/local/
RUN go mod download
RUN go build -v -o app src/app.go
ENTRYPOINT ["./app"]