FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/mstrangfeld/open-credit-score-chain/backend
COPY . ./
RUN go get -v -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/mstrangfeld/open-credit-score-chain/backend/backend /usr/local/bin/backend
ENTRYPOINT [ "/usr/local/bin/backend" ]
