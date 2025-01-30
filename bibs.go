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

func (c *Client) RawGetHolding(ctx context.Context, mmsID, holdingID string) ([]byte, error) {
	return c.rawRequest(ctx, "GET", fmt.Sprintf("/bibs/%s/holdings/%s", mmsID, holdingID), nil, nil)
}

type GetHoldingItemsParams struct {
	Limit                   int    `url:"limit,omitempty"`
	Offset                  int    `url:"offset,omitempty"`
	Expand                  string `url:"expand,omitempty"`
	UserID                  string `url:"user_id,omitempty"`
	CurrentLibrary          string `url:"current_library,omitempty"`
	CurrentLocation         string `url:"current_location,omitempty"`
	Query                   string `url:"q,omitempty"`
	OrderBy                 string `url:"order_by,omitempty"`
	Direction               string `url:"direction,omitempty"`
	CreateDateFrom          string `url:"create_date_from,omitempty"`
	CreateDateTo            string `url:"create_date_to,omitempty"`
	ModifyDateFrom          string `url:"modify_date_from,omitempty"`
	ModifyDateTo            string `url:"modify_date_to,omitempty"`
	ReceiveDateFrom         string `url:"receive_date_from,omitempty"`
	ReceiveDateTo           string `url:"receive_date_to,omitempty"`
	ExpectedReceiveDateFrom string `url:"expected_receive_date_from,omitempty"`
	ExpectedReceiveDateTo   string `url:"expected_receive_date_to,omitempty"`
	View                    string `url:"view,omitempty"`
}

func (c *Client) RawGetHoldingItems(ctx context.Context, mmsID, holdingID string, params GetHoldingItemsParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", fmt.Sprintf("/bibs/%s/holdings/%s/items", mmsID, holdingID), params, nil)
}
