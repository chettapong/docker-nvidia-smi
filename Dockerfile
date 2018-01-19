FROM nvidia/cuda

RUN apt-get update && \
    apt-get -y install golang --no-install-recommends && \
    rm -r /var/lib/apt/lists/*

WORKDIR /go

COPY . .

RUN go get github.com/basgys/goxml2json
RUN go install -v
RUN go build -v -o bin/app src/app.go

CMD ["./bin/app"]
