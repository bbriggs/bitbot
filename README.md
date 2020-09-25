# Bitbot

A minimal IRC bot that was built for me, not you.

### Installation

#### Direct install
`go get github.com/bbriggs/bitbot && go install github.com/bbriggs/bitbot`

#### Docker

`docker run --rm -it bbriggs/bitbot --help`

_note_: For persistent data you may want to mount a volume to `/tmp` or whatever path you specify for the embedded DB. For certain plugins to load, bitbot expects a postgres DB. 
Also remember that if you're running in Docker, your Prometheus bind address must be 0.0.0.0 + some port and you must publish that port using `-p`

#### docker-compose

See the [example deployment](docker-compose.yml)

#### Helm

See the [helm chart](https://artifacthub.io/packages/helm/bbriggs/bitbot)

### Usage
```
A Golang IRC bot powered by Hellabot

Usage:
  bitbot [flags]

Flags:
  -c, --channels strings     channels to join
      --config string        config file (default is ./config.yaml)
      --dbHost string        Postgresql host
      --dbName string        Postgresql database name
      --dbPass string        Postgresql password
      --dbPort string        Postgresql port
      --dbSSLMode string     Postgresql SSL Mode
      --dbUser string        Postgresql user
      --embedded-db string   The path to the embedded DB
  -h, --help                 help for bitbot
  -n, --nick string          nickname
      --nickserv string      nickserv password
      --operPass string      oper password
      --operUser string      oper username
      --prom                 enable prometheus
      --promAddr string      Prometheus metrics address and port
  -s, --server string        target server
      --ssl                  enable ssl
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
prom: "true"
promAddr: "0.0.0.0:8080"
dbuser: "bitbot"
dbpass: "bitbot"
dbhost: "127.0.0.1"
dbport: "5432"
dbsslmode: "disable"
# Plugins available to load are defined in cmd/bot.go
plugins:
  - "roll"
  - "skip"
  - "info"
  - "shrug"
  - "urlReader"
```

### Contributing

See our [contributing guide](CONTRIBUTING.md)

### License

Bitbot is available under the [MIT License](LICENSE)

