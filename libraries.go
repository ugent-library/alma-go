package alma

import (
	"context"
	"fmt"
)

func (c *Client) RawGetLibrary(ctx context.Context, id string) ([]byte, error) {
	return c.rawRequest(ctx, "GET", fmt.Sprintf("/conf/libraries/%s", id), nil, nil)
}

func (c *Client) RawGetLibraryOpenHours(ctx context.Context, id string) ([]byte, error) {
	return c.rawRequest(ctx, "GET", fmt.Sprintf("/conf/libraries/%s/open-hours", id), nil, nil)
}
