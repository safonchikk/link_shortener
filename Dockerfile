FROM golang:latest as buildbase

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/main .

FROM alpine:latest

WORKDIR /app

COPY --from=buildbase /app/main .
COPY --from=buildbase /app/app.env .
RUN chmod +x /app/main

EXPOSE 3000

ENTRYPOINT ["main"]