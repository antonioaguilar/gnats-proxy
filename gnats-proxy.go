package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/buger/jsonparser"
	"github.com/julienschmidt/httprouter"
	nats "github.com/nats-io/go-nats"
)

var debug bool
var nc *nats.Conn

func proxy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	subject, err := jsonparser.GetString(body, "__subject")

	if err != nil {
		log.Printf("[ERROR] Subject property missing\n")
	} else {
		nc.Publish(subject, body)
	}

	if debug == true {
		log.Printf("[INFO] %s\n", body)
	}

	w.WriteHeader(200)
	defer r.Body.Close()
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gnats-proxy [options]\n\n")
	flag.PrintDefaults()
}

func main() {
	var err error
	var port uint
	var certFile string
	var keyFile string
	var natsURL string

	flag.UintVar(&port, "p", 8080, "`Port number`")
	flag.StringVar(&certFile, "C", "", "Server certificate `file`")
	flag.StringVar(&keyFile, "K", "", "Private key `file`")
	flag.StringVar(&natsURL, "n", "nats://0.0.0.0:4222", "NATS server `URL`")
	flag.BoolVar(&debug, "d", true, "Enable `debug` output")

	flag.Usage = usage
	flag.Parse()

	router := httprouter.New()
	router.POST("/", proxy)

	proxyPort := fmt.Sprintf(":%d", port)

	// if disconnected from NATS, try to reconnect indefinitely
	nc, err = nats.Connect(natsURL, nats.MaxReconnects(-1), nats.ReconnectWait(5*time.Second))

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("[INFO] Connected to NATS server: %s\n", natsURL)
	}

	if certFile != "" && keyFile != "" {
		log.Printf("[INFO] Server listening on: https://0.0.0.0%s\n", proxyPort)
		err = http.ListenAndServeTLS(proxyPort, certFile, keyFile, router)
	} else {
		log.Printf("[INFO] Server listening on: http://0.0.0.0%s\n", proxyPort)
		err = http.ListenAndServe(proxyPort, router)
	}
	if err != nil {
		log.Fatal(err)
	}
}
