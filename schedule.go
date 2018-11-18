package iracing

import (
	"encoding/json"
	"net/http"
)

type ScheduleRes struct {
	Contents       []Contents `json:"contents"`
	NumActiveRaces string     `json:"num_active_races"`
	Status         Status     `json:"status"`
}

type Contents struct {
	Bannerhideat string `json:"bannerhideat"`
	Bannershowat string `json:"bannershowat"`
	Bannertext   string `json:"bannertext"`
	Eventat      string `json:"eventat"`
	Now          string `json:"now"`
}

type Status struct {
	HTTPCode int `json:"http_code"`
}

func GetSchedule() (*ScheduleRes, error) {
	res, err := http.Get("https://www.iracing.com/live/schedule/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var s ScheduleRes
	json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

}
