package alma

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Config struct {
	URL    string
	ApiKey string
}

type Client struct {
	config     Config
	httpClient *http.Client
}

func New(config Config) *Client {
	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) getRaw(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.config.URL+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "apikey "+c.config.ApiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		return b, fmt.Errorf("http error, status code: %d", res.StatusCode)
	}

	return b, nil
}

func decodeRaw[T any](b []byte, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	var t T
	if err := json.Unmarshal(b, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
