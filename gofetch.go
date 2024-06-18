package gofetch

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
)

// Config holds the configuration for the HTTP client.
type Config struct {
	BaseUrl string
}

// Client is an interface for making HTTP requests.
type Client interface {
	DoRequest(method, endpoint, body string, headers map[string]string) (string, error)
}

type client struct {
	config Config
	client *http.Client
}

// New creates a new instance of the Client with the given configuration.
func New(config Config) (Client, error) {
	if config.BaseUrl == "" {
		return nil, errors.New("invalid configuration: BaseUrl and ApiKey are required")
	}
	return &client{
		config: config,
		client: &http.Client{},
	}, nil
}

// DoRequest makes an HTTP request with the specified method, endpoint, body, and headers.
func (c *client) DoRequest(method, endpoint, body string, headers map[string]string) (string, error) {
	var req *http.Request
	var err error

	url := strings.TrimRight(c.config.BaseUrl, "/") + "/" + strings.TrimLeft(endpoint, "/")

	if method == http.MethodGet || method == http.MethodDelete {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBufferString(body))
		// if err == nil {
		// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// }
	}

	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", errors.New("failed to fetch data: " + resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
