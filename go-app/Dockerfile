FROM golang:1.18-alpine

COPY . /go/src
WORKDIR /go/src

# パッケージインストール
RUN apk update && \
    apk add --no-cache mysql-client vim && \
    # install revel
    go install github.com/revel/cmd/revel@latest