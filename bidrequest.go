package openrtb

import (
	"errors"

	"github.com/goccy/go-json"
)

// Validation errors
var (
	ErrInvalidReqNoID     = errors.New("openrtb: request ID missing")
	ErrInvalidReqNoImps   = errors.New("openrtb: request has no impressions")
	ErrInvalidReqMultiInv = errors.New("openrtb: request has multiple inventory sources") // has site and app
)

// BidRequest is the top-level bid request object contains a globally unique bid request or auction ID.  This "id"
// attribute is required as is at least one "imp" (i.e., impression) object.  Other attributes are
// optional since an exchange may establish default values.
type BidRequest struct {
	Impressions       []Impression      `json:"imp,omitempty"`     // -
	Seats             []string          `json:"wseat"`             // Array of buyer seats allowed to bid on this auction
	BlockedSeats      []string          `json:"bseat"`             // Array of buyer seats blocked to bid on this auction
	Languages         []string          `json:"wlang,omitempty"`   // Allowed list of languages for creatives using ISO-639-1-alpha-2. Omission implies no specific restrictions, but buyers would be advised to consider language attribute in the Device and/or Content objects if available. Only one of wlang or wlangb should be present.
	LanguagesB        []string          `json:"wlangb,omitempty"`  // Allowed list of languages for creatives using IETF BCP 47I. Omission implies no specific restrictions, but buyers would be advised to consider language attribute in the Device and/or Content objects if available. Only one of wlang or wlangb should be present.
	Currencies        []string          `json:"cur,omitempty"`     // Array of allowed currencies
	BlockedCategories []ContentCategory `json:"bcat,omitempty"`    // Blocked Advertiser Categories.
	BlockedAdvDomains []string          `json:"badv"`              // Array of strings of blocked toplevel domains of advertisers
	BlockedApps       []string          `json:"bapp"`              // Block list of applications by their platform-specific exchange-independent application identifiers. On Android, these should be bundle or package names (e.g., com.foo.mygame).  On iOS, these are numeric IDs.
	Ext               json.RawMessage   `json:"ext"`               // -
	ID                string            `json:"id"`                // Unique ID of the bid request
	Site              *Site             `json:"site,omitempty"`    // -
	App               *App              `json:"app,omitempty"`     // -
	Device            *Device           `json:"device,omitempty"`  // -
	User              *User             `json:"user,omitempty"`    // -
	Source            *Source           `json:"source,omitempty"`  // A Source object that provides data about the inventory source and which entity makes the final decision
	Regulations       *Regulations      `json:"regs,omitempty"`    // -
	TimeMax           int16             `json:"tmax,omitempty"`    // Maximum amount of time in milliseconds to submit a bid
	Test              int8              `json:"test"`              // Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode
	AuctionType       int8              `json:"at"`                // Auction type, where 1 = First Price, 2 = Second Price Plus. Exchange-specific auction types can be defined using values greater than 500.
	AllImpressions    int8              `json:"allimps,omitempty"` // Flag to indicate whether exchange can verify that all impressions offered represent all of the impressions available in context, Default: 0
}

// Validate the request
func (req *BidRequest) Validate() error {
	if req.ID == "" {
		return ErrInvalidReqNoID
	} else if len(req.Impressions) == 0 {
		return ErrInvalidReqNoImps
	} else if req.Site != nil && req.App != nil {
		return ErrInvalidReqMultiInv
	}

	for i := range req.Impressions {
		imp := req.Impressions[i]
		if err := (&imp).Validate(); err != nil {
			return err
		}
	}

	return nil
}
