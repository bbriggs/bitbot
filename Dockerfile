FROM golang:1.11-alpine as builder

RUN apk update && apk add git build-base ca-certificates
RUN adduser -D -g 'bitbot' bitbot
WORKDIR /usr/local/go/src/github.com/bbriggs/bitbot
COPY . .
RUN ./docker-build.sh
RUN touch .bolt.db

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/bitbot /go/bin/bitbot
COPY --from=builder --chown=bitbot:bitbot /usr/local/go/src/github.com/bbriggs/bitbot/.bolt.db .
USER bitbot
ENTRYPOINT ["/go/bin/bitbot"]

