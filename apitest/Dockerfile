FROM golang:1.10.0 As goimage
RUN mkdir -p "/var/log/apitest" && chmod -R 777 "/var/log/apitest"
ENV GOBIN=/usr/local/go
ENV GOPATH=/usr/local/goworkspace
WORKDIR $GOPATH/src/github.com/ralali/apitest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apitest main.go

CMD ["./apitest"]
