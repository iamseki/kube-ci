FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o kube-ci .

FROM scratch

COPY --from=builder /app/kube-ci /usr/bin/ 
CMD ["kube-ci","--help"]