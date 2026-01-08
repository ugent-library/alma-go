package alma

import "context"

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

func (c *Client) RawGetUserFees(ctx context.Context, userID string, params GetUserFeesParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/users/"+userID+"/fees", params, nil)
}
