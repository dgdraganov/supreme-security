package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	serverRootCA := os.Getenv("SERVER_ROOT_CA")

	clientCertFile := os.Getenv("CLIENT_CRT")
	clientKeyFile := os.Getenv("CLIENT_KEY")

	caCert, err := os.ReadFile(serverRootCA)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair(
		clientCertFile,
		clientKeyFile,
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
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	url := fmt.Sprintf("https://%s:%s/hellos", host, port)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("new request: %s", err)
	}

	for i := 0; i < 5; i++ {
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("client do: %s ", err)
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("io read all: %s ", err)
		}

		log.Printf("[CLIENT]: %s", respBody)
		<-time.After(time.Second * 1)
	}

	log.Println("[CLIENT]: shut down")

}
