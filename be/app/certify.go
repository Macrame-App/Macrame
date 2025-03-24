package app

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func generateCertificate() tls.Certificate {
	fmt.Println("Generating certificate")
	var certError = false
	// Generate a private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		// fmt.Println(err)
		certError = true
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Macrame-Server"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(1, 0, 0),

		IsCA:        true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		// fmt.Println(err)
		certError = true
	}

	// Encode the certificate and private key to PEM
	certPEM := pem.EncodeToMemory(
		&pem.Block{Type: "CERTIFICATE", Bytes: certBytes},
	)
	keyPEM := pem.EncodeToMemory(
		&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)},
	)

	// Save the certificate and private key to files
	if err := os.WriteFile("server.crt", certPEM, 0644); err != nil {
		// fmt.Println(err)
		certError = true
	}
	if err := os.WriteFile("server.key", keyPEM, 0644); err != nil {
		// fmt.Println(err)
		certError = true
	}

	if !certError {
		return Certify()
	}

	return tls.Certificate{}
}

func Certify() tls.Certificate {
	fmt.Println("Loading certificate")
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")

	if err != nil {
		return generateCertificate()
	}

	return cert
}
