FROM golang:1.19.0-alpine3.16 AS build

WORKDIR /app
COPY . .

RUN go build main.go


FROM alpine:3.16

COPY --from=build /app/main /app/main
EXPOSE 7788
WORKDIR /app

CMD ["./main"]