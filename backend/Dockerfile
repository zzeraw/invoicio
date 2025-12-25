FROM golang:1.22-alpine AS build

WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
RUN go build -o /bin/api ./cmd/api

FROM alpine:3.19
RUN apk add --no-cache ca-certificates curl
WORKDIR /app
COPY --from=build /bin/api /app/api
EXPOSE 8080
ENV PORT=8080
CMD ["/app/api"]
