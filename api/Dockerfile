FROM golang:1.21.5-alpine3.19 as builder
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
COPY . /src
WORKDIR /src
RUN go build -ldflags="-w -s" -o ./build/api ./api/main.go

FROM gcr.io/distroless/static
COPY --from=builder /src/build/api /api
CMD ["/api"]
