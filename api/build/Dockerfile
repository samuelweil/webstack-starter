ARG base_img=golang:1.17-alpine

FROM ${base_img} as builder

WORKDIR /build

COPY go.* ./
RUN go mod download

COPY cmd cmd
COPY internal internal

RUN go build cmd/*

FROM ${base_img}

WORKDIR /app

COPY --from=builder /build/api .

EXPOSE 8080

CMD ["./api"]
