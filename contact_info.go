package alma

type ContactInfo struct {
	Address []Address `json:"address,omitempty"`
	Email   []Email   `json:"email,omitempty"`
	Phone   []Phone   `json:"phone,omitempty"`
}

type Address struct {
	AddressNote   string `json:"address_note,omitempty"`
	AddressType   []Code `json:"address_type,omitempty"`
	City          string `json:"city,omitempty"`
	Country       *Code  `json:"country,omitempty"`
	Line1         string `json:"line1,omitempty"`
	Line2         string `json:"line2,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Preferred     *bool  `json:"preferred,omitempty"`
	SegmentType   string `json:"segment_type,omitempty"`
	StartDate     string `json:"start_date,omitempty"`
	StateProvince string `json:"state_province,omitempty"`
}

type Email struct {
	EmailAddress string `json:"email_address,omitempty"`
	EmailType    []Code `json:"email_type,omitempty"`
	Preferred    *bool  `json:"preferred"`
	SegmentType  string `json:"segment_type,omitempty"`
}

type Phone struct {
	PhoneNumber  string `json:"phone_number,omitempty"`
	PhoneType    []Code `json:"phone_type,omitempty"`
	Preferred    *bool  `json:"preferred,omitempty"`
	PreferredSMS *bool  `json:"preferred_sms,omitempty"`
	SegmentType  string `json:"segment_type,omitempty"`
}
