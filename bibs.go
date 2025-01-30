package alma

import (
	"context"
	"fmt"
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

type GetBibParams struct {
	View   string   `url:"view,omitempty"`
	Expand []string `url:"expand,omitempty" del:","`
}

func (c *Client) RawGetBib(ctx context.Context, mmsID string, params GetBibParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", fmt.Sprintf("/bibs/%s", mmsID), params, nil)
}

func (c *Client) RawGetHoldings(ctx context.Context, mmsID string) ([]byte, error) {
	return c.rawRequest(ctx, "GET", fmt.Sprintf("/bibs/%s/holdings", mmsID), nil, nil)
}
