package client

import (
	"crypto/tls"
	"errors"
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

func (c *Client) AuthWithToken(tokenId, tokenSecret string) error {

	id := tokenId
	if id == "" {
		id = os.Getenv("PROXMOX_TOKEN_ID")
	}

	secret := tokenSecret
	if secret == "" {
		secret = os.Getenv("PROXMOX_TOKEN_SECRET")
	}

	if id == "" {
		return errors.New("token id invalid")
	}

	if secret == "" {
		return errors.New("token secret invalid")
	}

	c.session = &Session{
		TokenId:     id,
		TokenSecret: secret,
	}

	return nil
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
