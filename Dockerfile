FROM golang:1.20.4-alpine3.18 as build

WORKDIR /build

COPY . .

RUN go build -o order_service ./cmd/app/main.go


FROM alpine:3.18

WORKDIR /app

COPY --from=build /build/order_service /app/order_service

EXPOSE 9999
CMD [ "/app/order_service" ]