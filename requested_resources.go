package alma

import (
	"context"
)

type GetRequestedResourcesParams struct {
	Library    string `url:"library"`
	CircDesk   string `url:"circ_desk"`
	Location   string `url:"location,omitempty"`
	OrderBy    string `url:"order_by,omitempty"`
	Direction  string `url:"direction,omitempty"`
	PickupInst string `url:"pickup_inst,omitempty"`
	Reported   string `url:"reported,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
}

func (c *Client) RawGetRequestedResources(ctx context.Context, params GetRequestedResourcesParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/task-lists/requested-resources", params, nil)
}
