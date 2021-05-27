FROM golang:1.16 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm64
WORKDIR /root
COPY . .
RUN make

# runtime image
FROM alpine
COPY --from=builder /root/app /app/app
EXPOSE 4000
ENTRYPOINT ["/app/app"]