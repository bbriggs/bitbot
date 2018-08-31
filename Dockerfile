FROM golang:latest 
ADD . /usr/local/go/src/github.com/bbriggs/bitbot
WORKDIR /usr/local/go/src/github.com/bbriggs/bitbot
RUN go get ./...  # Because the IRC lib has stupid "private" packages that break with vendoring
RUN go install
ENTRYPOINT ["/usr/local/go/bin/bitbot"]

