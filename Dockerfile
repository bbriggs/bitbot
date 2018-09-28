FROM golang:latest 
ADD . /usr/local/go/src/github.com/bbriggs/bitbot
WORKDIR /usr/local/go/src/github.com/bbriggs/bitbot
RUN go install
ENTRYPOINT ["/usr/local/go/bin/bitbot"]

