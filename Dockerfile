FROM golang:1.26-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /infrum .

FROM scratch

COPY --from=builder /infrum /infrum

WORKDIR /workspace

ENTRYPOINT ["/infrum"]
