FROM golang:1.12-alpine as builder

RUN adduser -D -g 'bitbot' bitbot
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN ./docker-build.sh
RUN touch .bolt.db

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/opt/bitbot /opt/bitbot
COPY --from=builder --chown=bitbot:bitbot /app/.bolt.db .
# Our chosen default for Prometheus
EXPOSE 8080
USER bitbot
ENTRYPOINT ["/opt/bitbot"]


