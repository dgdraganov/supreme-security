package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	// caCert, err := os.ReadFile("cert/ca/ca.crt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		// RootCAs:    caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
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
