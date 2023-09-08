package main

import (
	"crypto/tls"
	"crypto/x509"
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
	ok := caCertPool.AppendCertsFromPEM(caCert)
	if !ok {
		log.Fatalln("could not append certs from pem")
	}

	tlsConfig := &tls.Config{
		RootCAs:    caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from the server"))
	})

	server := &http.Server{
		Addr:      ":9205",
		TLSConfig: tlsConfig,
	}

	certFile := "cert/server/server.crt"
	keyFile := "cert/server/server.unencrypted.key"

	log.Println("Starting service...")
	if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
		log.Fatalf("listen and server: %s", err)
	}
}
