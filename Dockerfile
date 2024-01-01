FROM golang:latest as builder
WORKDIR /sources
COPY . .
RUN go build

FROM ubuntu:latest
LABEL org.opencontainers.image.source https://github.com/stephenhu/nbaq
LABEL org.opencontainers.image.description="nba api"
LABEL org.opencontainers.image.licenses=MIT
WORKDIR /usr/local/nbaq
COPY --from=builder /sources/nbaq .
CMD ["/usr/local/nbaq/nbaq", "-src", "2023"]