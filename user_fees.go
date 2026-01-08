package alma

import (
	"context"
	"fmt"
)

// TODO fee array
type Fees struct {
	TotalRecordCount int     `json:"total_record_count" xml:"total_record_count,attr"`
	TotalSum         float64 `json:"total_sum" xml:"total_sum,attr"`
	Currency         string  `json:"currency" xml:"currency,attr"`
}

type GetUserFeesParams struct {
	UserIDType string `url:"user_id_type,omitempty"`
	Status     string `url:"status,omitempty"`
}

func (c *Client) RawGetUserFees(ctx context.Context, id string, params GetUserFeesParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", fmt.Sprintf("/users/%s/fees", id), params, nil)
}

func (c *Client) GetUserFees(ctx context.Context, id string, params GetUserFeesParams) (*Fees, error) {
	resData := &Fees{}
	if err := c.request(ctx, "GET", fmt.Sprintf("/users/%s", id), params, nil, resData); err != nil {
		return nil, err
	}
	return resData, nil
}
