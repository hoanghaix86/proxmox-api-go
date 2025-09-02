package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type Response[T any] struct {
	Data    T              `json:"data"`
	Message string         `json:"message"`
	Errors  map[string]any `json:"errors"`
}

func ToQueryString(v any) string {
	values, err := query.Values(v)
	if err != nil {
		log.Fatalf("failed to convert struct to query string: %v", err)
	}
	encoded := values.Encode()

	// replace true to 1, false to 0
	encoded = strings.ReplaceAll(encoded, "=true", "=1")
	encoded = strings.ReplaceAll(encoded, "=false", "=0")
	return encoded
}

func ToJsonBody(b any) *bytes.Buffer {
	body, err := json.Marshal(b)
	if err != nil {
		log.Fatalf("failed to marshal body: %v", err)
	}
	return bytes.NewBuffer(body)
}

func ParseJsonBody[T any](b []byte) *Response[T] {
	var data Response[T]
	if err := json.Unmarshal(b, &data); err != nil {
		log.Fatalf("failed to unmarshal response body: %v", err)
	}
	return &data
}

func DoRequest[T any](ctx context.Context, c *Client, method string, path string, params any, body any) *T {
	url := fmt.Sprintf("%s%s", c.apiUrl, path)

	if params != nil {
		url += "?" + ToQueryString(params)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, ToJsonBody(body))
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if c.session != nil {
		req.Header.Set("Authorization", fmt.Sprintf("PVEAPIToken=%s=%s", c.session.TokenId, c.session.TokenSecret))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("failed to do request: %v", err)
	}

	defer res.Body.Close()

	log.Default().Printf("%s %s %d", method, url, res.StatusCode)

	if res.StatusCode == http.StatusUnauthorized {
		log.Fatalf("Unauthorized")
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	log.Default().Printf("RAW => %s", string(raw))

	jsonBody := ParseJsonBody[T](raw)

	if jsonBody.Message != "" {
		log.Default().Println(jsonBody.Errors)
		log.Fatal(jsonBody.Message)
	}

	return &jsonBody.Data

}

func Get[T any](ctx context.Context, c *Client, path string, params any, body any) *T {
	return DoRequest[T](ctx, c, "GET", path, params, body)
}

func Post[T any](ctx context.Context, c *Client, path string, params any, body any) *T {
	return DoRequest[T](ctx, c, "POST", path, params, body)
}

func Delete[T any](ctx context.Context, c *Client, path string, params any, body any) *T {
	return DoRequest[T](ctx, c, "DELETE", path, params, body)
}

func Put[T any](ctx context.Context, c *Client, path string, params any, body any) *T {
	return DoRequest[T](ctx, c, "PUT", path, params, body)
}

func Patch[T any](ctx context.Context, c *Client, path string, params any, body any) *T {
	return DoRequest[T](ctx, c, "PATCH", path, params, body)
}
