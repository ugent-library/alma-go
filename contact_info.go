package alma

type ContactInfo struct {
	Address []Address `json:"address,omitempty" xml:"addresses>address,omitempty"`
	Email   []Email   `json:"email,omitempty" xml:"emails>email,omitempty"`
	Phone   []Phone   `json:"phone,omitempty" xml:"phones>phone,omitempty"`
}

type Address struct {
	AddressNote   string `json:"address_note,omitempty" xml:"address_note,omitempty"`
	AddressType   []Code `json:"address_type" xml:"address_types>address_type"`
	City          string `json:"city" xml:"city"`
	Country       *Code  `json:"country,omitempty" xml:"country,omitempty"`
	EndDate       string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Line1         string `json:"line1" xml:"line1"`
	Line2         string `json:"line2,omitempty" xml:"line2,omitempty"`
	Line3         string `json:"line3,omitempty" xml:"line3,omitempty"`
	Line4         string `json:"line4,omitempty" xml:"line4,omitempty"`
	Line5         string `json:"line5,omitempty" xml:"line5,omitempty"`
	PostalCode    string `json:"postal_code,omitempty" xml:"postal_code,omitempty"`
	Preferred     *bool  `json:"preferred,omitempty" xml:"preferred,attr,omitempty"`
	SegmentType   string `json:"segment_type,omitempty" xml:"segment_type,attr,omitempty"`
	StartDate     string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	StateProvince string `json:"state_province,omitempty" xml:"state_province,omitempty"`
}

type Email struct {
	Description  string `json:"description,omitempty" xml:"description,omitempty"`
	EmailAddress string `json:"email_address,omitempty" xml:"email_address,omitempty"`
	EmailType    []Code `json:"email_type,omitempty" xml:"email_types>email_type,omitempty"`
	Preferred    *bool  `json:"preferred,omitempty" xml:"preferred,attr,omitempty"`
	SegmentType  string `json:"segment_type,omitempty" xml:"segment_type,attr,omitempty"`
}

type Phone struct {
	PhoneNumber  string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	PhoneType    []Code `json:"phone_type,omitempty" xml:"phone_types>phone_type,omitempty"`
	Preferred    *bool  `json:"preferred,omitempty" xml:"preferred,attr,omitempty"`
	PreferredSMS *bool  `json:"preferred_sms,omitempty" xml:"preferred_sms,attr,omitempty"`
	SegmentType  string `json:"segment_type,omitempty" xml:"segment_type,attr,omitempty"`
}
