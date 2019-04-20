# Bitbot

A minimal IRC bot that was built for me, not you.

## Installation

`go get github.com/bbriggs/bitbot && go install github.com/bbriggs/bitbot`

## Usage

Assuming that `$GOPATH/bin` is in your `$PATH`: 

```
bitbot run --server=irc.someserver.com:6697 --channels="#foo,bar" --nick=bitbot --nickserv='th15ism3' --operUser=bbriggs --operPass='s0upers3cr3t' --ssl
```

Or with Docker:

```
docker run -d -it bbriggs/bitbot run --server=irc.someserver.com:6697 --channels="#foo,bar" --nick=bitbot --nickserv='th15ism3' --operUser=bbriggs --operPass='s0upers3cr3t' --ssl
```
