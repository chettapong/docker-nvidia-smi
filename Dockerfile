FROM nvidia/cuda:9.0-devel

RUN apt-get -y update
RUN apt-get -y install git automake libssl-dev libcurl4-openssl-dev
RUN apt-get -y install golang --no-install-recommends && \
    rm -r /var/lib/apt/lists/*

WORKDIR /go

COPY . .

ENV GOPATH /go

RUN go get github.com/basgys/goxml2json
RUN go build -v -o bin/app src/app.go

CMD ["./bin/app"]