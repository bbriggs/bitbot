# Bitbot

A minimal IRC bot that was built for me, not you.

### Installation

`go get github.com/bbriggs/bitbot && go install github.com/bbriggs/bitbot`

### Usage
```
± bitbot --help
A Golang IRC bot powered by Hellabot

Usage:
  bitbot [flags]

Flags:
  -c, --channels strings   channels to join
      --config string      config file (default is ./config.yaml)
  -h, --help               help for bitbot
  -n, --nick string        nickname
      --nickserv string    nickserv password
      --operPass string    oper password
      --operUser string    oper username
      --prom               enable prometheus
      --promAddr string    Prometheus metrics address and port
  -s, --server string      target server
      --ssl                enable ssl
```

All flags are also supported as config file parameters.
```yaml
---
server: "irc.secops.space:6697"
nickServ: "hunter2"
operUser: "your-oper-username"
operPass: "correct-horse-battery-staple"
channels:
  - "#main"
  - "#bots"
  - "#bitbot"
nick: "bitbot"
ssl: "true"
prom: "true"
promAddr: "127.0.0.0.1:8080"
admins:
  "foo@your.irc.hostmask"
# Plugins available to load are defined in cmd/bot.go
plugins:
  - "roll"
  - "skip"
  - "info"
  - "shrug"
  - "urlReader"
```

### Running with ~scissors~ Docker

Assuming a config file named `config.yaml` in your local directory:
```
docker run --rm -it -v `pwd`/config.yaml:/app/config.yaml bbriggs/bitbot --config /app/config.yaml
```

Remember that if you're running in Docker, your Prometheus bind address must be 0.0.0.0 + some port and you must publish that port using `-p`
