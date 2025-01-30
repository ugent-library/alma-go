package alma

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/schema"
)

type Config struct {
	URL    string
	ApiKey string
}

type Client struct {
	config       Config
	httpClient   *http.Client
	queryEncoder *schema.Encoder
}

func New(config Config) *Client {
	queryEncoder := schema.NewEncoder()
	queryEncoder.SetAliasTag("query")

	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		queryEncoder: queryEncoder,
	}
}

func (c *Client) rawRequest(ctx context.Context, method, path string, params any, body []byte) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewBuffer(body)
	}

	reqURL, err := c.url(path, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, reqBody)
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

func (c *Client) request(ctx context.Context, method, path string, params, reqData, resData any) error {
	var reqBody []byte

	if reqData != nil {
		b, err := json.Marshal(reqData)
		if err != nil {
			return err
		}
		reqBody = b
	}

	resBody, err := c.rawRequest(ctx, method, path, params, reqBody)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resBody, resData); err != nil {
		return err
	}

	return nil
}

func (c *Client) url(path string, params any) (string, error) {
	u := c.config.URL + path

	if params != nil {
		q := url.Values{}

		if err := c.queryEncoder.Encode(params, q); err != nil {
			return "", err
		}

		// cleanup and remove empty query params
		for k, vals := range q {
			var newVals []string
			for _, v := range vals {
				v = strings.TrimSpace(v)
				if v != "" {
					newVals = append(newVals, v)
				}
			}
			if len(newVals) == 0 {
				delete(q, k)
			}
		}

		if len(q) > 0 {
			u += "?" + q.Encode()
		}
	}

	return u, nil
}
