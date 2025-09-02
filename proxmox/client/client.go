package client

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
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
	return &Config{
		TLSInsecure: false,
		Timeout:     10 * time.Second,
	}
}

func NewClient(apiUrl string, config *Config) *Client {
	if config == nil {
		config = NewConfig()
	}

	httpClient := buildHttpClient(config)

	return &Client{
		apiUrl:     apiUrl,
		config:     config,
		httpClient: httpClient,
	}
}

func (c *Client) AuthWithToken(tokenId, tokenSecret string) {
	if tokenId == "" || tokenSecret == "" {
		log.Fatalf("token id or token secret is empty")
	}

	c.session = &Session{
		TokenId:     tokenId,
		TokenSecret: tokenSecret,
	}

	DoRequest[any](context.Background(), c, "GET", "/version", nil, nil)
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
