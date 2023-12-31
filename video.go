package openrtb

import (
	"errors"

	"github.com/goccy/go-json"
)

// Validation errors
var (
	ErrInvalidVideoNoMIMEs     = errors.New("openrtb: video has no mimes")
	ErrInvalidVideoNoLinearity = errors.New("openrtb: video linearity missing")
	ErrInvalidVideoNoProtocols = errors.New("openrtb: video protocols missing")
)

// Video object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	MIMEs           []string            `json:"mimes,omitempty"`          // Content MIME types supported.
	Protocols       []Protocol          `json:"protocols,omitempty"`      // Video bid response protocols
	BlockedAttrs    []CreativeAttribute `json:"battr,omitempty"`          // Blocked creative attributes
	RqdDurs         []int64             `json:"rqddurs,omitempty"`        // Precise acceptable durations for video creatives inseconds.
	PlaybackMethods []VideoPlayback     `json:"playbackmethod,omitempty"` // List of allowed playback methods
	Delivery        []ContentDelivery   `json:"delivery,omitempty"`       // List of supported delivery methods
	CompanionAds    []Banner            `json:"companionad,omitempty"`    // -
	APIs            []APIFramework      `json:"api,omitempty"`            // List of supported API frameworks
	CompanionTypes  []CompanionType     `json:"companiontype,omitempty"`  // -
	Ext             json.RawMessage     `json:"ext,omitempty"`            // -
	PodID           string              `json:"podid,omitempty"`          // Pod id unique identifier for video ad pod
	MinBitrate      int                 `json:"minbitrate,omitempty"`     // Minimum bit rate in Kbps
	MaxBitrate      int                 `json:"maxbitrate,omitempty"`     // Maximum bit rate in Kbps
	BoxingAllowed   *int                `json:"boxingallowed,omitempty"`  // If exchange publisher has rules preventing letter boxing
	MinCPMPerSecond float32             `json:"mincpmpersec,omitempty"`   //   Minimum CPM per second. This is a price floor for the portion of a video ad pod
	MinDuration     int16               `json:"minduration,omitempty"`    // Minimum video ad duration in seconds
	MaxDuration     int16               `json:"maxduration,omitempty"`    // Maximum video ad duration in seconds
	Width           int16               `json:"w"`                        // Width of the player in pixels
	Height          int16               `json:"h"`                        // Height of the player in pixels
	SkipMin         int16               `json:"skipmin,omitempty"`        // Videos of total duration greater than this number of seconds can be skippable
	SkipAfter       int16               `json:"skipafter,omitempty"`      // Number of seconds a video must play before skipping is enabled
	Sequence        int16               `json:"sequence,omitempty"`       // Default: 1
	MaxExtended     int16               `json:"maxextended,omitempty"`    // Maximum extended video ad duration
	PodDuration     int16               `json:"poddur,omitempty"`         // Pod Duration total amount of time in seconds that advertisers may fill for a video ad pod
	StartDelay      StartDelay          `json:"startdelay,omitempty"`     // Indicates the start delay in seconds
	Skip            int8                `json:"skip,omitempty"`           // Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	Protocol        Protocol            `json:"protocol,omitempty"`       // Video bid response protocols DEPRECATED
	Linearity       VideoLinearity      `json:"linearity,omitempty"`      // Indicates whether the ad impression is linear or non-linear
	PodSequence     PodSequence         `json:"podseq,omitempty"`         // Pod Sequence position of the video ad pod
	SlotInPod       SlotPositionInPod   `json:"slotinpod,omitempty"`      // Slot Position in ad
	Position        AdPosition          `json:"pos,omitempty"`            // Ad Position
	Placement       VideoPlacement      `json:"placement,omitempty"`      // Video placement type, DEPRECATED
	Plcmt           VideoPlcmt          `json:"plcmt,omitempty"`          // Video Plcmt type ad defined in ADCOM1.0
}

type jsonVideo Video

// Validate the object
func (v *Video) Validate() error {
	if len(v.MIMEs) == 0 {
		return ErrInvalidVideoNoMIMEs
	} else if v.Linearity == 0 {
		return ErrInvalidVideoNoLinearity
	} else if v.Protocol == 0 && len(v.Protocols) == 0 {
		return ErrInvalidVideoNoProtocols
	}
	return nil
}

// GetBoxingAllowed returns the boxing-allowed indicator
func (v *Video) GetBoxingAllowed() int {
	if v.BoxingAllowed != nil {
		return *v.BoxingAllowed
	}
	return 1
}

// MarshalJSON custom marshalling with normalization
func (v *Video) MarshalJSON() ([]byte, error) {
	v.normalize()
	return json.Marshal((*jsonVideo)(v))
}

// UnmarshalJSON custom unmarshalling with normalization
func (v *Video) UnmarshalJSON(data []byte) error {
	var h jsonVideo
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*v = (Video)(h)
	v.normalize()
	return nil
}

func (v *Video) normalize() {
	if v.Sequence == 0 {
		v.Sequence = 1
	}
	if v.Linearity == 0 {
		v.Linearity = VideoLinearityLinear
	}
}
