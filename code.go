package alma

type Code struct {
	Value string `json:"value" xml:",chardata"`
	Desc  string `json:"desc" xml:"-"`
}
