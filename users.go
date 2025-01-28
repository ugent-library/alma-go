package alma

import (
	"context"
	"fmt"
)

type User struct {
	AccountType         *Code        `json:"account_type,omitempty"`
	BirthDate           string       `json:"birth_date,omitempty"`
	CampusCode          *Code        `json:"campus_code,omitempty"`
	CatalogerLevel      *Code        `json:"cataloger_level,omitempty"`
	ContactInfo         *ContactInfo `json:"contact_info,omitempty"`
	CreatedBy           string       `json:"created_by,omitempty"`
	CreatedDate         string       `json:"created_date,omitempty"`
	ExpiryDate          string       `json:"expiry_date,omitempty"`
	ExternalID          string       `json:"external_id,omitempty"`
	FirstName           string       `json:"first_name,omitempty"`
	ForcePasswordChange string       `json:"force_password_change,omitempty"`
	FullName            string       `json:"full_name,omitempty"`
	Gender              *Code        `json:"gender,omitempty"`
	JobCategory         *Code        `json:"job_category,omitempty"`
	JobDescription      string       `json:"job_description,omitempty"`
	LastModifiedBy      string       `json:"last_modified_by,omitempty"`
	LastModifiedDate    string       `json:"last_modified_date,omitempty"`
	LastName            string       `json:"last_name,omitempty"`
	MiddleName          string       `json:"middle_name,omitempty"`
	Password            string       `json:"password,omitempty"`
	PinNumber           string       `json:"pin_number,omitempty"`
	PreferredLanguage   *Code        `json:"preferred_language,omitempty"`
	PrefFirstName       string       `json:"pref_first_name,omitempty"`
	PrefLastName        string       `json:"pref_last_name,omitempty"`
	PrefMiddleName      string       `json:"pref_middle_name,omitempty"`
	PrefNameSuffix      string       `json:"pref_name_suffix,omitempty"`
	PrimaryID           string       `json:"primary_id,omitempty"`
	RecordType          *Code        `json:"record_type,omitempty"`
	Status              *Code        `json:"status,omitempty"`
	StatusDate          string       `json:"status_date,omitempty"`
	UserGroup           *Code        `json:"user_group,omitempty"`
	UserTitle           *Code        `json:"user_title,omitempty"`
	WebsiteURL          string       `json:"web_site_url,omitempty"`
}

func (c *Client) GetRawUser(ctx context.Context, id string) ([]byte, error) {
	return c.getRaw(ctx, fmt.Sprintf("/users/%s", id))
}

func (c *Client) GetUser(ctx context.Context, id string) (*User, error) {
	b, err := c.GetRawUser(ctx, id)
	return decodeRaw[User](b, err)
}
