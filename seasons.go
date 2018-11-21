package iracing

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const seasonURL = "https://members.iracing.com/membersite/member/GetSeasons"

// DefaultSeasonFields used by the client
var DefaultSeasonFields = []string{"year", "quarter", "seriesshortname", "seriesid", "active", "catid", "carclasses", "tracks", "seasonid"}

// GetSeasonOpts tunes the GetSeasons query
type GetSeasonOpts struct {
	OnlyActive bool
	Fields     []string
}

// Encode turns GetSeasonOpts into their query string
func (o GetSeasonOpts) Encode() string {
	u := url.Values{
		"onlyActive": []string{"0"},
	}

	if o.OnlyActive {
		u.Set("onlyActive", "1")
	}
	if len(o.Fields) > 0 {
		u.Set("fields", strings.Join(o.Fields, ","))
	} else {
		u.Set("fields", strings.Join(DefaultSeasonFields, ","))
	}

	return u.Encode()
}

func (c *Client) GetSeasons(o *GetSeasonOpts) ([]Season, error) {
	if o == nil {
		o = &GetSeasonOpts{}
	}

	u, err := url.Parse(seasonURL)
	if err != nil {
		return nil, err
	}

	u.RawQuery = o.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var s []Season
	return s, json.NewDecoder(res.Body).Decode(&s)
}

type Season struct {
	Active          bool       `json:"active"`
	CarClasses      []CarClass `json:"carclasses"`
	CatID           int        `json:"catid"`
	Quarter         int        `json:"quarter"`
	SeasonID        int        `json:"seasonid"`
	SeriesID        int        `json:"seriesid"`
	SeriesShortName String     `json:"seriesshortname"`
	Tracks          []Track    `json:"tracks"`
	Year            int        `json:"year"`
}

type CarClass struct {
	CarsInClass []CarsInClass `json:"carsinclass"`
	CustID      int           `json:"custid"`
	ID          int           `json:"id"`
	LowerName   String        `json:"lowername"`
	Name        String        `json:"name"`
	RelSpeed    int           `json:"relspeed"`
	ShortName   String        `json:"shortname"`
}

type Track struct {
	Config    String `json:"config"`
	ID        int    `json:"id"`
	Lowername String `json:"lowername"`
	Name      String `json:"name"`
	PkgID     int    `json:"pkgid"`
	Priority  int    `json:"priority"`
	RaceWeek  int    `json:"raceweek"`
	TimeOfDay int    `json:"timeOfDay"`
}

type CarsInClass struct {
	ID   int    `json:"id"`
	Name String `json:"name"`
}
