package main

import (
	"crypto/tls"
	"encoding/base64"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4443"
	}

	certPem := getEnvBase64("TLS_CERT")
	keyPem := getEnvBase64("TLS_KEY")

	if len(certPem) == 0 || len(keyPem) == 0 {
		log.Fatalf("TLS_CERT_PEM or TLS_KEY_PEM is empty")
	}

	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatalf("error load x509 key pair: %v", err)
	}

	s := http.Server{
		Addr: ":" + port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		}),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	s.SetKeepAlivesEnabled(false)
	log.Printf("start server on port %s", port)
	err = s.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalf("error start server: %v", err)
	}
}

func getEnvBase64(key string) []byte {
	v := os.Getenv(key)
	b, _ := base64.StdEncoding.DecodeString(v)
	return b
}
