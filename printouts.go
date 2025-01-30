package alma

import (
	"context"
)

type GetPrintoutsParams struct {
	Letter     string `url:"letter,omitempty"`
	Status     string `url:"status,omitempty"`
	PrinterID  string `url:"printer_id,omitempty"`
	PrintedBy  string `url:"printed_by,omitempty"`
	PrintoutID string `url:"printout_id,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
}

func (c *Client) RawGetPrintouts(ctx context.Context, params GetPrintoutsParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/task-lists/printouts", params, nil)
}
