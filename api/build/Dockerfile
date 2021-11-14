ARG base_img=golang:1.17-alpine

FROM ${base_img} as builder

WORKDIR /build

COPY go.* .
COPY cmd cmd

RUN go build cmd/*

FROM ${base_img}

WORKDIR /app

COPY --from=builder /build/api .

CMD ["api"]
