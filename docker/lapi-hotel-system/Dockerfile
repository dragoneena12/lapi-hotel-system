FROM golang:1.22 as builder

WORKDIR /work
COPY . .
RUN go build -o app

# runtime image
FROM gcr.io/distroless/base
COPY --from=builder /work/app /app
ENTRYPOINT ["/app"]
