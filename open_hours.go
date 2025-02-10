package alma

import (
	"context"
)

type GetOpenHoursParams struct {
	Scope string `url:"scope,omitempty"`
}

func (c *Client) RawGetOpenHours(ctx context.Context, params GetOpenHoursParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/conf/open-hours", params, nil)
}

type UpdateOpenHoursParams struct {
	Scope string `url:"scope,omitempty"`
}

func (c *Client) RawUpdateOpenHours(ctx context.Context, params UpdateOpenHoursParams, body []byte) ([]byte, error) {
	return c.rawRequest(ctx, "PUT", "/conf/open-hours", params, body)
}
