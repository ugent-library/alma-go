package alma

import (
	"context"
)

type GetBibsParams struct {
	MmsID            []string `url:"mms_id,omitempty" del:","`
	IeID             string   `url:"ie_id,omitempty"`
	HoldingsID       string   `url:"holdings_id,omitempty"`
	RepresentationID string   `url:"representation_id,omitempty"`
	NzMmsID          string   `url:"nz_mms_id,omitempty"`
	CzMmsID          string   `url:"cz_mms_id,omitempty"`
	View             string   `url:"view,omitempty"`
	Expand           []string `url:"expand,omitempty" del:","`
	OtherSystemID    string   `url:"other_system_id,omitempty"`
	LodUri           string   `url:"lod_uri,omitempty"`
}

func (c *Client) RawGetBibs(ctx context.Context, params GetBibsParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/bibs", params, nil)
}
