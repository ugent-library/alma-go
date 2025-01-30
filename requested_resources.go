package alma

import (
	"context"
)

type GetRequestedResourcesParams struct {
	Library    string `query:"library"`
	CircDesk   string `query:"circ_desk"`
	Location   string `query:"location"`
	OrderBy    string `query:"order_by"`
	Direction  string `query:"direction"`
	PickupInst string `query:"pickup_inst"`
	Reported   string `query:"reported"`
	Limit      int    `query:"limit"`
	Offset     int    `query:"offset"`
}

func (c *Client) RawGetRequestedResources(ctx context.Context, params GetRequestedResourcesParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/task-lists/requested-resources", params, nil)
}
