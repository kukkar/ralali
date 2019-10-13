FROM golang:1.10.0 As goimage
ARG app_env
ENV GOPATH /usr/local/personelDev/goworkspace
WORKDIR /usr/local/personelDev/goworkspace/src/github.com/ralali/apitest/
COPY . /usr/local/personelDev/goworkspace/src/github.com/ralali/apitest/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apitest main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /apitest/

COPY --from=goimage /github.com/ralali/apitest/apitest .
COPY --from=goimage /github.com/ralali/apitest/conf/ conf/
COPY --from=goimage /usr/local/go/lib/time/zoneinfo.zip .
COPY version.txt .

ENV ZONEINFO=/apitest/zoneinfo.zip

CMD ["./apitest"]
