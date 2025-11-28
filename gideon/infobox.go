package gideon

type InfoboxType string

const (
	InfoboxPortable InfoboxType = "portable"
	InfoboxLegacy   InfoboxType = "legacy"
	InfoboxCustom   InfoboxType = "custom"
)

type InfoboxStructure struct {
	InfoboxType InfoboxType `json:"type"`
	Heading     string      `json:"heading"`
	CompatXML   *string     `json:"xml"`
}

type InfoboxSection struct {
	Heading  string `json:"heading"`
	Contents []struct {
		Key   string `json:"key"`
		Value any    `json:"value"`
	} `json:"contents"`
}

// Base infobox parser
func (f *PageResponse[any]) Infobox() (*InfoboxStructure, error) {
	return &InfoboxStructure{}, nil
}

type InfoboxWithConfig struct {
	Selector string
}

func (f *PageResponse[any]) InfoboxConfig(cfg InfoboxWithConfig) (*InfoboxStructure, error) {
	return &InfoboxStructure{}, nil
}

// Returns the infobox contents as-is if the wiki uses the modern infobox
//
// In most cases, most infoboxes on wikis uses a simple key-value structure, but
// it won't account for any edge cases where an infobox has custom formatting on
// its XML template. For granular control on how infobox content is handled,
// use `Infobox.Parse()` instead.
func (f *InfoboxStructure) Basic() InfoboxStructure {}

func (f *InfoboxStructure) Parse(callback func(section []InfoboxSection)) (InfoboxStructure, error) {
	return InfoboxStructure{}, nil
}
