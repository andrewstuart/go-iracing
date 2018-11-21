package iracing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type ScheduleRes struct {
	Contents       []Contents `json:"contents"`
	NumActiveRaces string     `json:"num_active_races"`
	Status         Status     `json:"status"`
}

type Contents struct {
	Bannerhideat Time   `json:"bannerhideat"`
	Bannershowat Time   `json:"bannershowat"`
	Bannertext   string `json:"bannertext"`
	EventAt      Time   `json:"eventat"`
	Now          Time   `json:"now"`
}

// Time handles converting string-wrappted unix time, sometime with
// decimal, and turning it into an embedded time.Time.
type Time struct {
	time.Time
}

func (i *Time) UnmarshalJSON(bs []byte) error {
	bs = bytes.Split(bytes.Trim(bs, "\""), []byte{'.'})[0]
	unix, err := strconv.Atoi(string(bs))
	if err != nil {
		return err
	}

	i.Time = time.Unix(int64(unix)/1000, 0)
	return nil
}

type Status struct {
	HTTPCode int `json:"http_code"`
}

// GetSchedule returns the current
func (c *Client) GetSchedule() (*ScheduleRes, error) {
	res, err := http.Get("https://www.iracing.com/live/schedule/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var s ScheduleRes
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
