FROM golang:1.12

COPY . /app

RUN cd /app/cli && go build -o m3u8-proxy && mv m3u8-proxy $GOPATH/bin/ && rm -rf /app

ENTRYPOINT ["m3u8-proxy"]
