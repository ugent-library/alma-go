package alma

import (
	"context"
)

type GetPrintoutsParams struct {
	Letter     string `query:"letter"`
	Status     string `query:"status"`
	PrinterID  string `query:"printer_id"`
	PrintedBy  string `query:"printed_by"`
	PrintoutID string `query:"printout_id"`
	Limit      int    `query:"limit"`
	Offset     int    `query:"offset"`
}

func (c *Client) RawGetPrintouts(ctx context.Context, params GetPrintoutsParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/task-lists/printouts", params, nil)
}
