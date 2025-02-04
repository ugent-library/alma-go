package alma

import (
	"context"
)

type GetRequestedResourcesParams struct {
	CircDesk   string `url:"circ_desk"`
	Direction  string `url:"direction,omitempty"`
	Library    string `url:"library"`
	Limit      int    `url:"limit,omitempty"`
	Location   string `url:"location,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	OrderBy    string `url:"order_by,omitempty"`
	PickupInst string `url:"pickup_inst,omitempty"`
	Reported   string `url:"reported,omitempty"`
}

func (c *Client) RawGetRequestedResources(ctx context.Context, params GetRequestedResourcesParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/task-lists/requested-resources", params, nil)
}
