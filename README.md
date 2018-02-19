# gnats-proxy
[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/antonioaguilar/gnats-proxy)](https://goreportcard.com/report/github.com/antonioaguilar/gnats-proxy)

This proxy server allows clients to arbitrarily specify a NATS.io subject and publish messages directly to NATS.io without using the default [NATS.io client libraries](https://nats.io/download/).

## Install

```bash
$ go get -u github.com/antonioaguilar/gnats-proxy
```

## Usage

```bash
$ gnats-proxy --help

usage: gnats-proxy [options]

  -C file
    	Server certificate file
  -K file
    	Private key file
  -d debug
    	Enable debug output (default: true)
  -n URL
    	NATS server URL (default: "nats://0.0.0.0:4222")
  -p Port number
    	Port number (default: 8080)
```

## Publishing data to NATS

External clients can issue HTTP post requests to the default route `/` and publish messages directly to [NATS.io](https://nats.io/), for example:

```bash
curl -s -H "Content-Type: application/json" \
-X POST -d '{"__subject":"CUSTOMER","account":"ACC-123456789","orders":"PO-123456789"}' \
http://localhost:8080/
```

This command will post the JSON data to the default route [http://localhost:8080/](http://localhost:8080/), the `gnats-proxy` server will read and parse the JSON data and will publish this data to [NATS.io](https://nats.io/) with a subject called `CUSTOMER`.


## Run as Docker container

You can run ```gnats-proxy``` in a Docker container as follows:

```bash
# pull the image
docker pull aaguilar/gnats-proxy

# run the container
docker run -it --rm -p 8080:8080 aaguilar/gnats-proxy -p 8080 -n nats://localhost:4222
```
