FROM alpine
ADD *.go /root/install/
RUN apk update
RUN apk add --no-cache go git libc-dev
RUN cd /root/install && \
    go get github.com/buger/jsonparser && \
    go get github.com/julienschmidt/httprouter && \
    go get github.com/nats-io/go-nats && \
    go build -o /root/install/gnats-proxy
ADD script.sh /script.sh
ENTRYPOINT ["/script.sh"]
