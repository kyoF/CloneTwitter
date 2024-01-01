FROM golang:1.21.1-alpine AS builder

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/main /app/main.go


FROM builder AS runner

WORKDIR /app

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/main /app/main

EXPOSE 8888
CMD [ "/app/main" ]

