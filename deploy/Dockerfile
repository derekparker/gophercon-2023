FROM golang:latest as builder

COPY . /src

WORKDIR /src

ENV CGO_ENABLED=0

RUN go build -o /bin/count ./cmd/count

FROM alpine:latest

COPY --from=builder /bin/count /count

EXPOSE 3333

CMD ["/count"]
