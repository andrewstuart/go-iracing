package iracing

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Category is the category of race
type Category int

//go:generate stringer -type Category -type LicenseClass
const (
	Oval Category = iota + 1
	Road
	Mud
	RX
)

var (
	// AllCats is a list of all categories
	AllCats = []Category{Oval, Road, Mud, RX}

	catByName = map[string]Category{}
)

func init() {
	for _, c := range AllCats {
		catByName[strings.ToLower(c.String())] = c
	}
}

// LookupCategory returns the Category by its string name
func LookupCategory(s string) Category {
	return catByName[strings.ToLower(s)]
}

// A LicenseClass is the license class for a race
type LicenseClass int

// Known LicenseClasses
const (
	Rookie LicenseClass = iota + 1
	D
	C
	B
	A
	Pro
)

var plusReplacer = strings.NewReplacer("+", " ")

// A String is a URL-encoded string, common in the iRacing API
type String string

// UnmarshalJSON implements json unmarshaler
func (i *String) UnmarshalJSON(bs []byte) error {
	bs = bytes.Trim(bs, "\"")
	st, err := url.QueryUnescape(string(bs))
	if err != nil {
		st = plusReplacer.Replace(string(bs))
	}
	*i = String(st)
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
	CatID           Category         `json:"catID"`
	Eligible        bool             `json:"eligible"`
	Image           string           `json:"image"`
	Mpr             int              `json:"mpr"`
	SeasonSchedules []SeasonSchedule `json:"seasonSchedules"`
	SeriesID        int              `json:"seriesID"`
	SeriesName      String           `json:"seriesName"`
}

func (s Series) CurrentSchedule() *SeasonSchedule {
	for _, sched := range s.SeasonSchedules {
		if time.Since(sched.SeasonStartDate.Time) < 12*7*24*time.Hour {
			return &sched
		}
	}
	return nil
}

func (s Series) NextSchedule() *SeasonSchedule {
	var latest *SeasonSchedule
	for i, sched := range s.SeasonSchedules {
		d := time.Until(sched.SeasonStartDate.Time)
		if d > 0 && (latest == nil || d < time.Until(latest.SeasonStartDate.Time)) {
			latest = &s.SeasonSchedules[i]
		}
	}
	return latest
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
	CarClasses           []int        `json:"carClasses"`
	FixedSetup           bool         `json:"fixedSetup"`
	LicenseGroup         LicenseClass `json:"licenseGroup"`
	MultiClass           bool         `json:"multiClass"`
	OpenPracticeDrivers  int          `json:"openPracticeDrivers"`
	OpenPracticeSessions int          `json:"openPracticeSessions"`
	Races                []Race       `json:"races"`
	SeasonID             int          `json:"seasonID"`
	SeasonStartDate      UnixTime     `json:"seasonStartDate"`
}

func (ss SeasonSchedule) NextRace() *Race {
	if len(ss.Races) < 1 {
		return nil
	}
	if len(ss.Races) < 2 {
		return &ss.Races[0]
	}

	var next *Race
	for i, race := range ss.Races {
		d := time.Until(race.StartTime.Time)
		if d > 0 && (next == nil || race.StartTime.Before(next.StartTime.Time)) {
			next = &ss.Races[i]
		}
	}
	return next
}

type Race struct {
	EndTime                 UnixTime       `json:"endTime"`
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
	TrackConfigName         String         `json:"trackConfigName"`
	TrackID                 int            `json:"trackID"`
	TrackName               String         `json:"trackName"`
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
