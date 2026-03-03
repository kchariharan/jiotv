package epg

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
)

// Channel XML tag structure for the EPG
type Channel struct {
	XMLName xml.Name `xml:"channel"`      // XML tag name
	ID      int      `xml:"id,attr"`      // ID is attribute of channel tag
	Display string   `xml:"display-name"` // Display name of the channel
}

// Icon XML tag for Programme XML tag in EPG
type Icon struct {
	XMLName xml.Name `xml:"icon"`     // XML tag name
	Src     string   `xml:"src,attr"` // Src is attribute of the icon tag
}

// Title XML tag for Programme XML tag in EPG
// Title is the name of the programme or show being aired on the channel
type Title struct {
	XMLName xml.Name `xml:"title"`
	Value   string   `xml:",chardata"` // Title of the programme
	Lang    string   `xml:"lang,attr"` // Language of the title
}

// Category XML tag for Programme XML tag in EPG
// Category is the type of the programme or show being aired on the channel
type Category struct {
	XMLName xml.Name `xml:"category"`
	Value   string   `xml:",chardata"` // Category of the programme
	Lang    string   `xml:"lang,attr"` // Language of the category
}

// Desc represents Description XML tag for Programme XML tag in EPG
type Desc struct {
	XMLName xml.Name `xml:"desc"`
	Value   string   `xml:",chardata"` // Description of the programme
	Lang    string   `xml:"lang,attr"` // Language of the description
}

// Programme XML tag structure for EPG
// Each programme tag represents a show being aired on a channel
type Programme struct {
	XMLName  xml.Name `xml:"programme"`    // XML tag name
	Channel  string   `xml:"channel,attr"` // Channel is attribute of programme tag
	Start    string   `xml:"start,attr"`   // Start time of the programme
	Stop     string   `xml:"stop,attr"`    // Stop time of the programme
	Title    Title    `xml:"title"`        // Title of the programme
	Desc     Desc     `xml:"desc"`         // Description of the programme
	Category Category `xml:"category"`     // Category of the programme
	Icon     Icon     `xml:"icon"`         // Icon of the programme
}

// EPG XML tag structure
type EPG struct {
	XMLName     xml.Name    `xml:"tv"`            // XML tag name
	XMLVersion  string      `xml:"version,attr"`  // XML version
	XMLEncoding string      `xml:"encoding,attr"` // XML encoding
	Channel     []Channel   `xml:"channel"`       // Channel tags
	Programme   []Programme `xml:"programme"`     // Programme tags
}

// ChannelObject represents Individual channel detail from JioTV API response
type ChannelObject struct {
	ChannelID   JSONInt    `json:"channel_id"`   // Channel ID
	ChannelName JSONString `json:"channel_name"` // Channel name
	LogoURL     JSONString `json:"logoUrl"`      // Channel logo URL
}

// ChannelsResponse represents Channel details from JioTV API response
type ChannelsResponse struct {
	Channels []ChannelObject `json:"result"`  // Channels
	Code     int             `json:"code"`    // Response code
	Message  string          `json:"message"` // Response message
}

// EPGObject represents Individual EPG detail from JioTV EPG API response
type EPGObject struct {
	StartEpoch   JSONInt64  `json:"startEpoch"`       // Start time of the programme
	EndEpoch     JSONInt64  `json:"endEpoch"`         // End time of the programme
	ChannelID    JSONInt    `json:"channel_id"`       // Channel ID
	ChannelName  JSONString `json:"channel_name"`     // Channel name
	ShowCategory JSONString `json:"showCategory"`     // Category of the show
	Description  JSONString `json:"description"`      // Description of the show
	Title        JSONString `json:"showname"`         // Title of the show
	Thumbnail    JSONString `json:"episodeThumbnail"` // Thumbnail of the show
	Poster       JSONString `json:"episodePoster"`    // Poster of the show
}

// EPGResponse represents EPG details from JioTV EPG API response
type EPGResponse struct {
	EPG []EPGObject `json:"epg"` // EPG details for a channel
}

// JSONInt64 is a custom type for unmarshaling int64 from both strings and integers
type JSONInt64 int64

// UnmarshalJSON unmarshals both strings and integers into int64
func (i *JSONInt64) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*i = JSONInt64(val)
		return nil
	}
	var val int64
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*i = JSONInt64(val)
	return nil
}

// JSONInt is a custom type for unmarshaling int from both strings and integers
type JSONInt int

// UnmarshalJSON unmarshals both strings and integers into int
func (i *JSONInt) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		*i = JSONInt(val)
		return nil
	}
	var val int
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*i = JSONInt(val)
	return nil
}

// JSONString is a custom type for unmarshaling string from both string and array of strings
type JSONString string

// UnmarshalJSON unmarshals both string and array of strings into string
func (s *JSONString) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '[' {
		var arr []string
		if err := json.Unmarshal(data, &arr); err != nil {
			return err
		}
		if len(arr) > 0 {
			*s = JSONString(arr[0])
		} else {
			*s = ""
		}
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = JSONString(str)
	return nil
}

func (s JSONString) String() string {
	return string(s)
}
