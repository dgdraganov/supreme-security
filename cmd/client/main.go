package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	caCert, err := os.ReadFile("cert/ca/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair(
		"cert/client/client.crt",
		"cert/client/client.unencrypted.key",
	)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	url := "https://server:9205/hello"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("new request: %s ", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client do: %s ", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("io read all: %s ", err)
	}

	log.Printf("response from server: %s", respBody)
}
