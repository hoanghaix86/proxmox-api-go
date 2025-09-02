package client

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	TLSInsecure bool
	Timeout     time.Duration
}

type Session struct {
	TokenId     string
	TokenSecret string
}

type Client struct {
	apiUrl     string
	config     *Config
	session    *Session
	httpClient *http.Client
}

func NewConfig() *Config {
	tlsInsecure := os.Getenv("PROXMOX_TLS_INSECURE") == "true"
	return &Config{
		TLSInsecure: tlsInsecure,
		Timeout:     10 * time.Second,
	}
}

func NewClient(apiUrl *string, config *Config) *Client {
	endpoint := os.Getenv("PROXMOX_API_URL")
	if apiUrl != nil {
		endpoint = *apiUrl
	}

	if config == nil {
		config = NewConfig()
	}

	httpClient := buildHttpClient(config)

	return &Client{
		apiUrl:     endpoint,
		config:     config,
		httpClient: httpClient,
	}
}

func (c *Client) AuthWithToken(tokenId, tokenSecret *string) {

	tkId := os.Getenv("PROXMOX_TOKEN_ID")
	tkSecret := os.Getenv("PROXMOX_TOKEN_SECRET")

	if tokenId != nil {
		tkId = *tokenId
	}

	if tokenSecret != nil {
		tkId = *tokenSecret
	}

	if tkId == "" {
		log.Fatal("Token Id invalid")
	}

	if tkSecret == "" {
		log.Fatal("Token Secret invalid")
	}

	c.session = &Session{
		TokenId:     tkId,
		TokenSecret: tkSecret,
	}
}

func buildHttpClient(config *Config) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.TLSInsecure,
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * config.Timeout,
	}

	return client
}
