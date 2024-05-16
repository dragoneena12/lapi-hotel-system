FROM golang:1.22 as builder

WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 go build -o app

# runtime image
FROM gcr.io/distroless/static-debian12
COPY --from=builder /work/app /app
ENTRYPOINT ["/app"]
