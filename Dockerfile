FROM golang:alpine AS build

WORKDIR /go/src/app
COPY . /go/src/app

RUN apk add --no-cache --upgrade --virtual .requirements build-base && \
    CGO_ENABLED=1 go build -v -o elliot main.go && \
    apk del .requirements

# ================================

FROM alpine

RUN addgroup -S elliot && adduser -s /sbin/nologin -S elliot -G elliot
USER elliot

WORKDIR /app
COPY --from=build /go/src/app/elliot /bin/elliot

CMD ["/bin/elliot"]