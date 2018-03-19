package livebarn

// Sport represents a sport captured by LiveBarn
type Sport struct {
	Name string `json:"name"`
}

// Surface represents an ice arena's specific ice surface
type Surface struct {
	Name        string      `json:"name,omitempty"`
	UUID        string      `json:"uuid,omitempty"`
	OrderIndex  int         `json:"orderIndex,omitempty"`
	Sport       Sport       `json:"sport,omitempty"`
	ClosedFrom  interface{} `json:"closedFrom,omitempty"`
	ClosedUntil interface{} `json:"closedUntil,omitempty"`
	ComingSoon  bool        `json:"comingSoon,omitempty"`
}

// Venue represents an ice arena
type Venue struct {
	Address1            string  `json:"address1"`
	Address2            string  `json:"address2"`
	City                string  `json:"city"`
	Name                string  `json:"name"`
	UUID                string  `json:"uuid"`
	RecordingHoursLocal string  `json:"recordingHoursLocal"`
	PostalCode          string  `json:"postalCode"`
	FreePromoCodes      int     `json:"freePromoCodes"`
	AllSheetsCount      int     `json:"allSheetsCount"`
	Longitude           float64 `json:"longitude"`
	Latitude            float64 `json:"latitude"`
}
