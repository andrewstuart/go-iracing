package iracing

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func (c Client) Search(term string) (*SearchResponse, error) {
	addr := "https://members.iracing.com/membersite/member/GetDriverStatus?searchTerms=" + uglifyString(term)
	res, err := c.Get(addr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var sRes SearchResponse
	err = json.NewDecoder(res.Body).Decode(&sRes)
	if err != nil {
		return nil, err
	}
	return &sRes, nil
}

type SearchResponse struct {
	Blacklisted bool `json:"blacklisted"`
	Friends     bool `json:"friends"`
	// FsRacers    []Member `json:"fsRacers"`
	Search  int     `json:"search"`
	Racers  []Racer `json:"searchRacers"`
	Studied bool    `json:"studied"`
}

type Racer struct {
	// Broadcast        `json:"broadcast"`
	CatID         int    `json:"catId"`
	CustID        int    `json:"custid"`
	DriverChanges bool   `json:"driverChanges"`
	EventTypeID   int    `json:"eventTypeId"`
	Helmet        Helmet `json:"helmet"`
	LastLogin     Time   `json:"lastLogin"`
	LastSeen      int    `json:"lastSeen"`
	MaxUsers      int    `json:"maxUsers"`
	Name          String `json:"name"`
	// PrivateSession   `json:"privateSession"`
	PrivateSessionID int    `json:"privateSessionId"`
	RegOpen          bool   `json:"regOpen"`
	SeasonID         int    `json:"seasonId"`
	SeriesID         int    `json:"seriesId"`
	SessionStatus    string `json:"sessionStatus"`
	SessionTypeID    int    `json:"sessionTypeId"`
	SpotterAccess    int    `json:"spotterAccess"`
	StartTime        int    `json:"startTime"`
	SubSessionStatus string `json:"subSessionStatus"`
	TrackID          int    `json:"trackId"`
	UserRole         int    `json:"userRole"`
}

func uglifyString(s string) string {
	bs := []byte(s)
	ss := make([]string, len(bs))

	for i := range bs {
		ss[i] = fmt.Sprintf("&#%d;", int(bs[i]))
	}

	return url.QueryEscape(strings.Join(ss, ""))
}
