FROM golang:1.21-alpine as buildbase

WORKDIR /go/src/app
COPY . .
RUN go mod tidy

RUN GOOS=linux go build -o /usr/local/bin/core


FROM alpine:3.19

RUN apk add --no-cache ca-certificates
COPY --from=buildbase /usr/local/bin/core /usr/local/bin/core
COPY --from=buildbase /go/src/app/app.env /usr/local/bin/app.env
COPY --from=buildbase /go/src/app/app.env /app.env

ENTRYPOINT ["core"]