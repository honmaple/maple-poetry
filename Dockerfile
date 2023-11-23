FROM node:lts-alpine AS web-builder

WORKDIR /src

COPY web/yarn.lock  web/*.json  web/*.js /src
RUN yarn global add @quasar/cli && yarn

COPY web/index.html /src/index.html
COPY web/src /src/src
COPY web/public /src/public
RUN quasar build


FROM golang:1.19-alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache gcc musl-dev

WORKDIR /src

COPY go.mod go.sum /src
RUN GOPROXY=https://proxy.golang.com.cn,direct go mod download

COPY internal /src/internal
COPY main.go /src/main.go
COPY --from=web-builder /src/dist /src/web/dist
RUN go build .

FROM alpine:latest

WORKDIR /opt/poetry

COPY --from=builder /src/poetry /usr/bin

CMD ["/usr/bin/poetry"]