FROM golang:1.9-alpine AS builder

WORKDIR /go/src/github.com/Nais777/happy-pancake

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo -o app .

FROM scratch

WORKDIR /

COPY --from=builder /go/src/github.com/Nais777/happy-pancake/app /

CMD ["/app"]