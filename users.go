package alma

import (
	"context"
	"encoding/xml"
	"fmt"
)

// TODO requests, loans, fees, user_roles, user_blocks, user_notes, user_statistics, proxy_for_users, library_notices, is_researcher, researcher
type User struct {
	XMLName                xml.Name         `json:"-" xml:"user"`
	AccountType            *Code            `json:"account_type,omitempty" xml:"account_type,omitempty"`
	BirthDate              string           `json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	CampusCode             *Code            `json:"campus_code,omitempty" xml:"campus_code,omitempty"`
	CatalogerLevel         *Code            `json:"cataloger_level,omitempty" xml:"cataloger_level,omitempty"`
	ContactInfo            *ContactInfo     `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
	CreatedBy              string           `json:"created_by,omitempty" xml:"created_by,omitempty"`
	CreatedDate            string           `json:"created_date,omitempty" xml:"created_date,omitempty"`
	ExpiryDate             string           `json:"expiry_date,omitempty" xml:"expiry_date,omitempty"`
	ExternalID             string           `json:"external_id,omitempty" xml:"external_id,omitempty"`
	FirstName              string           `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ForcePasswordChange    string           `json:"force_password_change,omitempty" xml:"force_password_change,omitempty"`
	FullName               string           `json:"full_name,omitempty" xml:"full_name,omitempty"`
	Gender                 *Code            `json:"gender,omitempty" xml:"gender,omitempty"`
	JobCategory            *Code            `json:"job_category,omitempty" xml:"job_category,omitempty"`
	JobDescription         string           `json:"job_description,omitempty" xml:"job_description,omitempty"`
	LastModifiedBy         string           `json:"last_modified_by,omitempty" xml:"last_modified_by,omitempty"`
	LastModifiedDate       string           `json:"last_modified_date,omitempty" xml:"last_modified_date,omitempty"`
	LastName               string           `json:"last_name,omitempty" xml:"last_name,omitempty"`
	LastPatronActivityDate string           `json:"last_patron_activity_date,omitempty" xml:"last_patron_activity_date,omitempty"`
	LinkingID              string           `json:"linking_id,omitempty" xml:"linking_id,omitempty"`
	MiddleName             string           `json:"middle_name,omitempty" xml:"middle_name,omitempty"`
	Password               string           `json:"password,omitempty" xml:"password,omitempty"`
	PinNumber              string           `json:"pin_number,omitempty" xml:"pin_number,omitempty"`
	PreferredLanguage      *Code            `json:"preferred_language,omitempty" xml:"preferred_language,omitempty"`
	PrefFirstName          string           `json:"pref_first_name,omitempty" xml:"pref_first_name,omitempty"`
	PrefLastName           string           `json:"pref_last_name,omitempty" xml:"pref_last_name,omitempty"`
	PrefMiddleName         string           `json:"pref_middle_name,omitempty" xml:"pref_middle_name,omitempty"`
	PrefNameSuffix         string           `json:"pref_name_suffix,omitempty" xml:"pref_name_suffix,omitempty"`
	PrimaryID              string           `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	PurgeDate              string           `json:"purge_date,omitempty" xml:"purge_date,omitempty"`
	RsLibrary              []RsLibrary      `json:"rs_library,omitempty" xml:"rs_libraries>rs_library,omitempty"`
	RecordType             *Code            `json:"record_type,omitempty" xml:"record_type,omitempty"`
	SourceLinkID           string           `json:"source_link_id,omitempty" xml:"source_link_id,omitempty"`
	SourceInstitutionCode  string           `json:"source_institution_code,omitempty" xml:"source_institution_code,omitempty"`
	Status                 *Code            `json:"status,omitempty" xml:"status,omitempty"`
	StatusDate             string           `json:"status_date,omitempty" xml:"status_date,omitempty"`
	UserGroup              *Code            `json:"user_group,omitempty" xml:"user_group,omitempty"`
	UserIdentifier         []UserIdentifier `json:"user_identifier,omitempty" xml:"user_identifiers>user_identifier,omitempty"`
	UserNote               []UserNote       `json:"user_note,omitempty" xml:"user_note,omitempty"`
	UserTitle              *Code            `json:"user_title,omitempty" xml:"user_title,omitempty"`
	WebsiteURL             string           `json:"web_site_url,omitempty" xml:"web_site_url,omitempty"`
}

type UserIdentifier struct {
	IdType Code   `json:"id_type" xml:"id_type"`
	Value  string `json:"value" xml:"value"`
	Note   string `json:"note,omitempty" xml:"note,omitempty"`
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

type UserNote struct {
	CreatedBy    string `json:"created_by,omitempty" xml:"created_by,omitempty"`
	NoteText     string `json:"note_text" xml:"note_text"`
	NoteType     Code   `json:"note_type" xml:"note_type"`
	PopupNote    bool   `json:"popup_note" xml:"popup_note"`
	SegmentType  string `json:"segment_type,omitempty" xml:"segment_type,omitempty"`
	UserViewable bool   `json:"user_viewable" xml:"user_viewable"`
}

type GetUsersParams struct {
	Expand                string `url:"expand,omitempty"`
	Limit                 int    `url:"limit,omitempty"`
	Offset                int    `url:"offset,omitempty"`
	OrderBy               string `url:"order_by,omitempty"`
	Query                 string `url:"q,omitempty"`
	SourceInstitutionCode string `url:"source_institution_code,omitempty"`
	SourceUserID          string `url:"source_user_id,omitempty"`
}

func (c *Client) RawGetUsers(ctx context.Context, params GetUsersParams) ([]byte, error) {
	return c.rawRequest(ctx, "GET", "/users", params, nil)
}

func (c *Client) RawGetUser(ctx context.Context, id string) ([]byte, error) {
	return c.rawRequest(ctx, "GET", fmt.Sprintf("/users/%s", id), nil, nil)
}

func (c *Client) GetUser(ctx context.Context, id string) (*User, error) {
	resData := &User{}
	if err := c.request(ctx, "GET", fmt.Sprintf("/users/%s", id), nil, nil, resData); err != nil {
		return nil, err
	}
	return resData, nil
}

func (c *Client) RawUpdateUser(ctx context.Context, id string, body []byte) ([]byte, error) {
	return c.rawRequest(ctx, "PUT", fmt.Sprintf("/users/%s", id), nil, body)
}

func (c *Client) UpdateUser(ctx context.Context, id string, data *User) (*User, error) {
	resData := &User{}
	if err := c.request(ctx, "PUT", fmt.Sprintf("/users/%s", id), nil, data, resData); err != nil {
		return nil, err
	}
	return resData, nil
}

type DeleteUserParams struct {
	UserIDType string `url:"user_id_type,omitempty"`
}

func (c *Client) DeleteUser(ctx context.Context, id string, params DeleteUserParams) error {
	return c.request(ctx, "DELETE", fmt.Sprintf("/users/%s", id), params, nil, nil)
}
