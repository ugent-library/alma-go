package alma

import (
	"bytes"
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

func (c *Client) rawRequest(ctx context.Context, method, path string, body []byte) ([]byte, error) {
	var reqBody io.Reader

	if body != nil {
		reqBody = bytes.NewBuffer(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.config.URL+path, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "apikey "+c.config.ApiKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		return resBody, fmt.Errorf("http error %d: %s", res.StatusCode, resBody)
	}

	return resBody, nil
}

func (c *Client) request(ctx context.Context, method, path string, reqData any, resData any) error {
	var reqBody []byte

	if reqData != nil {
		b, err := json.Marshal(reqData)
		if err != nil {
			return err
		}
		reqBody = b
	}

	resBody, err := c.rawRequest(ctx, method, path, reqBody)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resBody, resData); err != nil {
		return err
	}

	return nil
}
