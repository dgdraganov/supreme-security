package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	clientRootCA := os.Getenv("CLIENT_ROOT_CA")

	serverCertFile := os.Getenv("SERVER_CRT")
	serverKeyFile := os.Getenv("SERVER_KEY")

	caCert, err := os.ReadFile(clientRootCA)
	if err != nil {
		log.Fatalf("os read file: %s", err)
	}

	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(caCert)
	if !ok {
		log.Fatalln("could not append certs to cert pool")
	}

	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	http.HandleFunc("/hellos", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		w.Write([]byte("This is a secure 'Hello from the server' becasue it is using mutual TLS!"))
		if err != nil {
			log.Printf("response write: %s", err)
		}

	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := &http.Server{
		Addr:      port,
		TLSConfig: tlsConfig,
	}

	log.Println("Starting service...")
	if err := server.ListenAndServeTLS(serverCertFile, serverKeyFile); err != nil {
		log.Fatalf("listen and server: %s", err)
	}
}
