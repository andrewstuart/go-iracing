package iracing

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var plusReplacer = strings.NewReplacer("+", " ")

type IRString string

func (i *IRString) UnmarshalJSON(bs []byte) error {
	bs = bytes.Trim(bs, "\"")
	st, err := url.QueryUnescape(string(bs))
	if err != nil {
		st = plusReplacer.Replace(string(bs))
	}
	*i = IRString(st)
	return nil
}

func (c Client) GetRaceGuide() (*RaceGuideRes, error) {
	res, err := c.Get("http://members.iracing.com/membersite/member/GetRaceGuide")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r RaceGuideRes
	return &r, json.NewDecoder(res.Body).Decode(&r)
}

type RaceGuideRes struct {
	Date   UnixTime `json:"date"`
	Series []Series `json:"series"`
}

type Series struct {
	CatID           int              `json:"catID"`
	Eligible        bool             `json:"eligible"`
	Image           string           `json:"image"`
	Mpr             int              `json:"mpr"`
	SeasonSchedules []SeasonSchedule `json:"seasonSchedules"`
	SeriesID        int              `json:"seriesID"`
	SeriesName      string           `json:"seriesName"`
}

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(bs []byte) error {
	i, err := strconv.Atoi(string(bs))
	if err != nil {
		return err
	}

	u.Time = time.Unix(int64(i/1000), 0)
	return nil
}

type SeasonSchedule struct {
	CarClasses           []int    `json:"carClasses"`
	FixedSetup           bool     `json:"fixedSetup"`
	LicenseGroup         int      `json:"licenseGroup"`
	MultiClass           bool     `json:"multiClass"`
	OpenPracticeDrivers  int      `json:"openPracticeDrivers"`
	OpenPracticeSessions int      `json:"openPracticeSessions"`
	Races                []Race   `json:"races"`
	SeasonID             int      `json:"seasonID"`
	SeasonStartDate      UnixTime `json:"seasonStartDate"`
}

type Race struct {
	EndTime                 int            `json:"endTime"`
	PreRegCount             int            `json:"preRegCount"`
	RaceLapLimit            int            `json:"raceLapLimit"`
	RaceTimeLimitMinutes    int            `json:"raceTimeLimitMinutes"`
	RaceWeekCars            RaceWeekCars   `json:"raceWeekCars"`
	RaceWeekNum             int            `json:"raceWeekNum"`
	RegCount                int            `json:"regCount"`
	RubberSettings          RubberSettings `json:"rubberSettings"`
	SessionID               int            `json:"sessionID"`
	SessionTypeID           int            `json:"sessionTypeID"`
	StandingStart           bool           `json:"standingStart"`
	StartTime               UnixTime       `json:"startTime"`
	TimeOfDay               UnixTime       `json:"timeOfDay"`
	TrackConfigName         string         `json:"trackConfigName"`
	TrackID                 int            `json:"trackID"`
	TrackName               string         `json:"trackName"`
	TrackRaceGuideImg       string         `json:"trackRaceGuideImg"`
	WeatherFogDensity       int            `json:"weatherFogDensity"`
	WeatherRelativeHumidity int            `json:"weatherRelativeHumidity"`
	WeatherSkies            int            `json:"weatherSkies"`
	WeatherTempUnits        int            `json:"weatherTempUnits"`
	WeatherTempValue        int            `json:"weatherTempValue"`
	WeatherType             int            `json:"weatherType"`
	WeatherWindDir          int            `json:"weatherWindDir"`
	WeatherWindSpeedUnits   int            `json:"weatherWindSpeedUnits"`
	WeatherWindSpeedValue   int            `json:"weatherWindSpeedValue"`
}

type RaceWeekCars map[int]map[int]FuelAndWeight

type RubberSettings struct {
	LeaveMarbles        bool   `json:"leaveMarbles"`
	RubberLevelPractice string `json:"rubberLevelPractice"`
	RubberLevelQualify  string `json:"rubberLevelQualify"`
	RubberLevelRace     string `json:"rubberLevelRace"`
	RubberLevelWarmUp   string `json:"rubberLevelWarmUp"`
}

type FuelAndWeight struct {
	MaxPctFuelFill  int `json:"maxPctFuelFill"`
	WeightPenaltyKG int `json:"weightPenaltyKG"`
}

func (f *FuelAndWeight) UnmarshalJSON(bs []byte) error {
	s := map[string]interface{}{}
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	if val, ok := s["maxPctFuelFill"]; ok {
		if ival, ok := val.(int); ok {
			f.MaxPctFuelFill = ival
		}
	}
	if val, ok := s["weightPenaltyKG"]; ok {
		if ival, ok := val.(int); ok {
			f.MaxPctFuelFill = ival
		}
	}
	return nil
}
