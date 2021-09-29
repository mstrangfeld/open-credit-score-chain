FROM golang:1.16 AS builder
COPY . /go/src/github.com/mstrangfeld/backend
RUN go build -o /go/src/github.com/mstrangfeld/backend/backend

FROM debian:buster
LABEL Author="Marvin Strangfeld <marvin@strangfeld.io>"
COPY --from=builder /go/src/github.com/mstrangfeld/backend/backend /usr/local/bin/backend
ENTRYPOINT [ "/usr/local/bin/backend" ]
