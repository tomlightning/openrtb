package openrtb

import "github.com/goccy/go-json"

// Content object describes the content in which the impression will appear, which may be syndicated or nonsyndicated
// content. This object may be useful when syndicated content contains impressions and does
// not necessarily match the publisher's general content. The exchange might or might not have
// knowledge of the page where the content is running, as a result of the syndication method. For
// example might be a video impression embedded in an iframe on an unknown web property or device.
type Content struct {
	Categories         []ContentCategory `json:"cat,omitempty"`                // Array of IAB content categories that describe the content.
	Data               []Data            `json:"data,omitempty"`               // Additional content data.
	KwArray            []string          `json:"kwarray,omitempty"`            // Array of keywords about the site. Only one of ‘keywords’ or‘kwarray’ may be present.
	Ext                json.RawMessage   `json:"ext,omitempty"`                // -
	ID                 string            `json:"id,omitempty"`                 // ID uniquely identifying the content.
	Title              string            `json:"title,omitempty"`              // Content title.
	Series             string            `json:"series,omitempty"`             // Content series.
	Season             string            `json:"season,omitempty"`             // Content season.
	Artist             string            `json:"artist,omitempty"`             // Artist credited with the content.
	Genre              string            `json:"genre,omitempty"`              // Genre that best describes the content
	Album              string            `json:"album,omitempty"`              // Album to which the content belongs; typically for audio.
	ISRC               string            `json:"isrc,omitempty"`               // International Standard Recording Code conforming to ISO - 3901.
	URL                string            `json:"url,omitempty"`                // URL of the content, for buy-side contextualization or review.
	ContentRating      string            `json:"contentrating,omitempty"`      // Content rating (e.g., MPAA).
	UserRating         string            `json:"userrating,omitempty"`         // User rating of the content (e.g., number of stars, likes, etc.).
	Keywords           string            `json:"keywords,omitempty"`           // Comma separated list of keywords describing the content.
	Language           string            `json:"language,omitempty"`           // Content language using ISO-639-1-alpha-2.
	LanguageB          string            `json:"langb,omitempty"`              // Content language using IETF BCP 47. Only one of language or langb should be present.
	Producer           *Producer         `json:"producer,omitempty"`           // The producer.
	LiveStream         int               `json:"livestream,omitempty"`         // 0 = not live, 1 = content is live (e.g., stream, live blog).
	SourceRelationship int               `json:"sourcerelationship,omitempty"` // 0 = indirect, 1 = direct.
	Length             int               `json:"len,omitempty"`                // Length of content in seconds; appropriate for video or audio.
	Network            *ChannelEntity    `json:"network,omitempty"`            // Details about the network the content is on.
	Channel            *ChannelEntity    `json:"channel,omitempty"`            // Details about the channel the content is on.
	Episode            int8              `json:"episode,omitempty"`            // Episode number (typically applies to video content).
	CategoryTaxonomy   CategoryTaxonomy  `json:"cattax,omitempty"`             // Defines the taxonomy in use.
	ProductionQuality  ProductionQuality `json:"prodq,omitempty"`              // Production quality per IAB's classification.
	VideoQuality       ProductionQuality `json:"videoquality,omitempty"`       // DEPRECATED. Video quality per IAB's classification.
	Context            ContentContext    `json:"context,omitempty"`            // Type of content (game, video, text, etc.).
	MediaRating        IQGRating         `json:"qagmediarating,omitempty"`     // Media rating per QAG guidelines.
	Embeddable         int8              `json:"embeddable,omitempty"`         // Indicator of whether or not the content is embeddable (e.g., an embeddable video player), where 0 = no, 1 = yes.
}
