package openrtb

import "github.com/goccy/go-json"

// Banner object must be included directly in the impression object if the impression offered
// for auction is display or rich media, or it may be optionally embedded in the video object to
// describe the companion banners available for the linear or non-linear video ad.  The banner
// object may include a unique identifier; this can be useful if these IDs can be leveraged in the
// VAST response to dictate placement of the companion creatives when multiple companion ad
// opportunities of the same size are available on a page.
type Banner struct {
	Formats      []Format            `json:"format,omitempty"`   // Array of format objects representing the banner sizes permitted.
	BlockedTypes []BannerType        `json:"btype,omitempty"`    // Blocked banner types
	BlockedAttrs []CreativeAttribute `json:"battr,omitempty"`    // Blocked creative attributes
	MIMEs        []string            `json:"mimes,omitempty"`    // Whitelist of content MIME types supported
	ExpDirs      []ExpDir            `json:"expdir,omitempty"`   // Specify properties for an expandable ad
	APIs         []APIFramework      `json:"api,omitempty"`      // List of supported API frameworks
	Ext          json.RawMessage     `json:"ext,omitempty"`      // -
	ID           string              `json:"id,omitempty"`       // A unique identifier
	Width        int16               `json:"w,omitempty"`        // Width
	Height       int16               `json:"h,omitempty"`        // Height
	WidthMax     int16               `json:"wmax,omitempty"`     // Width maximum DEPRECATED
	HeightMax    int16               `json:"hmax,omitempty"`     // Height maximum DEPRECATED
	WidthMin     int16               `json:"wmin,omitempty"`     // Width minimum DEPRECATED
	HeightMin    int16               `json:"hmin,omitempty"`     // Height minimum DEPRECATED
	Position     AdPosition          `json:"pos,omitempty"`      // Ad Position
	TopFrame     int8                `json:"topframe,omitempty"` // Default: 0 ("1": Delivered in top frame, "0": Elsewhere)
	VCM          int8                `json:"vcm,omitempty"`      // Represents the relationship with video. 0 = concurrent, 1 = end-card
}
