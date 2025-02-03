package alma

import (
	"context"
	"fmt"
)

// TODO linked_record_id, holdings, enrichment_source, enrichment_workflow, cataloging_level, brief_level, warnings, requests
type Bib struct {
	MmsID                      string   `json:"mms_id,omitempty"`
	RecordFormat               string   `json:"record_format,omitempty"`
	Title                      string   `json:"title,omitempty"`
	Author                     string   `json:"author,omitempty"`
	ISBN                       string   `json:"isbn,omitempty"`
	CompleteEdition            string   `json:"complete_edition,omitempty"`
	NetworkNumbers             []string `json:"network_numbers,omitempty"`
	PlaceOfPublication         string   `json:"place_of_publication,omitempty"`
	DateOfPublication          string   `json:"date_of_publication,omitempty"`
	PublisherConst             string   `json:"publisher_const,omitempty"`
	CreatedBy                  string   `json:"created_by,omitempty"`
	CreatedDate                string   `json:"created_date,omitempty"`
	LastModifiedBy             string   `json:"last_modified_by,omitempty"`
	LastModifiedDate           string   `json:"last_modified_date,omitempty"`
	SuppressFromPublishing     string   `json:"suppress_from_publishing,omitempty"`
	SuppressFromExternalSearch string   `json:"suppress_from_external_search,omitempty"`
	SuppressFromMetaDoor       string   `json:"suppress_from_metadoor,omitempty"`
	Rank                       string   `json:"rank,omitempty"`
	SyncWithOCLC               string   `json:"sync_with_oclc,omitempty"`
	SyncWithLibrariesAustralia string   `json:"sync_with_libraries_australia,omitempty"`
	OriginatingSystem          string   `json:"originating_system,omitempty"`
	OriginatingSystemID        string   `json:"originating_system_id,omitempty"`
	Anies                      []string `json:"anies,omitempty"`
}

func (bib *Bib) Record() string {
	if len(bib.Anies) > 0 {
		return bib.Anies[0]
	}
	return ""
}

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

func (c *Client) GetBib(ctx context.Context, mmsID string, params GetBibParams) (*Bib, error) {
	resData := &Bib{}
	if err := c.request(ctx, "GET", fmt.Sprintf("/bibs/%s", mmsID), params, nil, resData); err != nil {
		return nil, err
	}
	return resData, nil
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
