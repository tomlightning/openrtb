package openrtb

import (
	"encoding/json"
	"errors"
)

var (
	errValidationResId      = errors.New("openrtb response: missing ID")
	errValidationResSeatbid = errors.New("openrtb response: missing seatbids")
	errValidationSeatbidBid = errors.New("openrtb response: seatbid is missing bids")
	errValidationBidId      = errors.New("openrtb response: bid is missing ID")
	errValidationBidImpid   = errors.New("openrtb response: bid is missing impression ID")
	errValidationBidPrice   = errors.New("openrtb response: bid is missing price")
)

// ID and at least one “seatbid” object is required, which contains a bid on at least one impression.
// Other attributes are optional since an exchange may establish default values.
// No-Bids on all impressions should be indicated as a HTTP 204 response.
// For no-bids on specific impressions, the bidder should omit these from the bid response.
type Response struct {
	Id         *string    `json:"id"`                   // Reflection of the bid request ID for logging purposes
	Seatbid    []Seatbid  `json:"seatbid"`              // Array of seatbid objects
	Bidid      *string    `json:"bidid,omitempty"`      // Optional response tracking ID for bidders
	Cur        *string    `json:"cur,omitempty"`        // Bid currency
	Customdata *string    `json:"customdata,omitempty"` // Encoded user features
	Ext        Extensions `json:"ext,omitempty"`        // Custom specifications in Json
}

// Response object to_json
// @return [Json,Error]
func (res *Response) JSON() ([]byte, error) {
	return json.Marshal(res)
}

// Validate Response required attributes
// @return [Boolean,Error] true if response,seatbid,bid required attrs present
func (res *Response) Valid() (bool, error) {

	if res.Id == nil {
		return false, errValidationResId
	} else if res.Seatbid == nil || len(res.Seatbid) < 1 {
		return false, errValidationResSeatbid
	}

	for _, sb := range res.Seatbid {
		if ok, err := sb.Valid(); !ok {
			return ok, err
		}
	}

	return true, nil
}

// Seatbid methods

// Validate Seatbid required attributes
// @return [Boolean,Error]
func (sb *Seatbid) Valid() (bool, error) {

	if sb.Bid == nil || len(sb.Bid) < 1 {
		return false, errValidationSeatbidBid
	}

	for _, bid := range sb.Bid {
		if ok, err := bid.Valid(); !ok {
			return ok, err
		}
	}

	return true, nil
}

// Bid methods

// Validate Bid required attributes
// @return [Boolean,Error]
func (bid *Bid) Valid() (bool, error) {

	if bid.Id == nil {
		return false, errValidationBidId
	} else if bid.Impid == nil {
		return false, errValidationBidImpid
	} else if bid.Price == nil {
		return false, errValidationBidPrice
	}

	return true, nil
}

// Set the ID
func (bid *Bid) SetID(id string) *Bid {
	if bid.Id == nil {
		bid.Id = new(string)
	}
	*bid.Id = id
	return bid
}

// Set the ImpID
func (bid *Bid) SetImpID(id string) *Bid {
	if bid.Impid == nil {
		bid.Impid = new(string)
	}
	*bid.Impid = id
	return bid
}

// Set the Price
func (bid *Bid) SetPrice(p float32) *Bid {
	if bid.Price == nil {
		bid.Price = new(float32)
	}
	*bid.Price = p
	return bid
}