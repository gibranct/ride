FROM golang:1.23-alpine AS builder

WORKDIR /src/app

# Restore dependencies
COPY . .

RUN go mod tidy

# Build executable
RUN go build -o account-app ./cmd

FROM golang:1.23-alpine

WORKDIR /src

COPY --from=builder /src/app/account-app ./

EXPOSE 3001

CMD ["/src/account-app"]