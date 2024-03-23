FROM golang:latest as builder
WORKDIR /sources
COPY . .
RUN go build

FROM ubuntu:latest
WORKDIR /usr/local/nbaq
COPY --from=builder /sources/nbaq .
COPY data .
EXPOSE 8000
CMD ["/usr/local/nbaq/nbaq", "-dir", "2023"]