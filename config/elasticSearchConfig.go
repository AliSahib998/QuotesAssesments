package config

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"net"
	"net/http"
	"time"
)

var ElasticDb *elasticsearch.Client

func LoadESClient() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 5,
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	conn, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic("repo search connection was not successful")
	}

	ElasticDb = conn
}
