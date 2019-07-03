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
docker run -d -it -v $(pwd):/some/directory bbriggs/bitbot --config /some/directory/config.yaml
```
