FROM golang:1.21 as build

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o gql .

FROM alpine:latest AS final

WORKDIR /app

COPY --from=build /app/gql .

CMD ["./gql"]