package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)


type Certificate struct {
	Organization string `json:"organization"`
}



func getSSLCert(domain string) (*x509.Certificate, error) {
	dialer := &net.Dialer{
		Timeout: time.Duration(4) * time.Second,
	}
	domain = domain + ":443"
	conn, err := tls.DialWithDialer(dialer, "tcp", domain, &tls.Config{
		InsecureSkipVerify: true,
	})

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	return cert, nil
}






func getOrganizationFromSSL(domain string) {
	data := make(map[string]interface{})
	cert, err := getSSLCert(domain)
	if err != nil{
		log.Println(err)
		return
	}

	org := cert.Subject.Organization


	if len(org) == 0 {
			return
	}
	data[domain] = org[0]
	jsondata, err := json.Marshal(data)
	if err != nil {
		log.Println("Error in whois lookup : ", err)
	}
	fmt.Println(string(jsondata))
	
}