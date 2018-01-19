FROM nvidia/cuda

RUN apt-get update && \
    apt-get -y install git && \
    apt-get -y install golang --no-install-recommends && \
    rm -r /var/lib/apt/lists/*

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

COPY . .

RUN go get github.com/basgys/goxml2json
RUN go build -v -o bin/app src/app.go

CMD ["./bin/app"]
