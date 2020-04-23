FROM golang:alpine AS build

WORKDIR /go/src/app
COPY . /go/src/app

RUN go build -o elliot main.go

# ================================

FROM alpine

RUN addgroup -S elliot && adduser -s /sbin/nologin -S elliot -G elliot
USER elliot

WORKDIR /app
COPY --from=build /go/src/app/elliot /app/elliot

ENTRYPOINT ["/app/elliot"]