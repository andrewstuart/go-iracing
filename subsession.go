package iracing

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func (c *Client) SubsessionResults(sess, cust int) (*SubsessionResults, error) {
	vals := url.Values{
		"subsessionID": []string{fmt.Sprint(sess)},
		"custid":       []string{fmt.Sprint(cust)},
	}.Encode()

	res, err := c.Post("http://members.iracing.com/membersite/member/GetSubsessionResults", "application/x-www-form-urlencoded", strings.NewReader(vals))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var sr SubsessionResults
	return &sr, json.NewDecoder(res.Body).Decode(&sr)
}

type SubsessionResults struct {
	Catid                 int     `json:"catid"`
	Cautiontype           int     `json:"cautiontype"`
	Cornersperlap         int     `json:"cornersperlap"`
	DriverChangeParam1    int     `json:"driver_change_param1"`
	DriverChangeParam2    int     `json:"driver_change_param2"`
	DriverChangeRule      int     `json:"driver_change_rule"`
	DriverChanges         int     `json:"driver_changes"`
	Eventavglap           int     `json:"eventavglap"`
	Eventlapscomplete     int     `json:"eventlapscomplete"`
	Eventstrengthoffield  int     `json:"eventstrengthoffield"`
	Evttype               int     `json:"evttype"`
	LeagueSeasonID        string  `json:"league_season_id"`
	Leagueid              string  `json:"leagueid"`
	Leavemarbles          int     `json:"leavemarbles"`
	MaxTeamDrivers        int     `json:"max_team_drivers"`
	Maxweeks              int     `json:"maxweeks"`
	MinTeamDrivers        int     `json:"min_team_drivers"`
	Ncautionlaps          int     `json:"ncautionlaps"`
	Ncautions             int     `json:"ncautions"`
	Nlapsforqualavg       int     `json:"nlapsforqualavg"`
	Nlapsforsoloavg       int     `json:"nlapsforsoloavg"`
	Nleadchanges          int     `json:"nleadchanges"`
	Pointstype            string  `json:"pointstype"`
	Privatesessionid      int     `json:"privatesessionid"`
	RaceWeekNum           int     `json:"race_week_num"`
	Rowss                 []Rows  `json:"rows"`
	RservStatus           string  `json:"rserv_status"`
	RubberlevelPractice   int     `json:"rubberlevel_practice"`
	RubberlevelQualify    int     `json:"rubberlevel_qualify"`
	RubberlevelRace       int     `json:"rubberlevel_race"`
	RubberlevelWarmup     int     `json:"rubberlevel_warmup"`
	SeasonName            string  `json:"season_name"`
	SeasonQuarter         int     `json:"season_quarter"`
	SeasonShortname       string  `json:"season_shortname"`
	SeasonYear            int     `json:"season_year"`
	Seasonid              int     `json:"seasonid"`
	SeriesName            string  `json:"series_name"`
	SeriesShortname       string  `json:"series_shortname"`
	Seriesid              int     `json:"seriesid"`
	Sessionid             int     `json:"sessionid"`
	Sessionname           string  `json:"sessionname"`
	Simsestype            int     `json:"simsestype"`
	StartTime             string  `json:"start_time"`
	Subsessionid          int     `json:"subsessionid"`
	Timeofday             int     `json:"timeofday"`
	TrackConfigName       string  `json:"track_config_name"`
	TrackName             string  `json:"track_name"`
	Trackid               int     `json:"trackid"`
	WeatherFogDensity     int     `json:"weather_fog_density"`
	WeatherRh             int     `json:"weather_rh"`
	WeatherSkies          int     `json:"weather_skies"`
	WeatherTempUnits      int     `json:"weather_temp_units"`
	WeatherTempValue      float64 `json:"weather_temp_value"`
	WeatherType           int     `json:"weather_type"`
	WeatherWindDir        int     `json:"weather_wind_dir"`
	WeatherWindSpeedUnits int     `json:"weather_wind_speed_units"`
	WeatherWindSpeedValue float64 `json:"weather_wind_speed_value"`
}

type Rows struct {
	Aggchamppoints    int     `json:"aggchamppoints"`
	Avglap            int     `json:"avglap"`
	Bestlapnum        int     `json:"bestlapnum"`
	Bestlaptime       int     `json:"bestlaptime"`
	Bestnlapsnum      int     `json:"bestnlapsnum"`
	Bestnlapstime     int     `json:"bestnlapstime"`
	Bestquallapat     int     `json:"bestquallapat"`
	Bestquallapnum    int     `json:"bestquallapnum"`
	Bestquallaptime   int     `json:"bestquallaptime"`
	CarColor1         string  `json:"car_color1"`
	CarColor2         string  `json:"car_color2"`
	CarColor3         string  `json:"car_color3"`
	CarNumberColor1   string  `json:"car_number_color1"`
	CarNumberColor2   string  `json:"car_number_color2"`
	CarNumberColor3   string  `json:"car_number_color3"`
	CarPattern        int     `json:"car_pattern"`
	Carclassid        int     `json:"carclassid"`
	Carid             int     `json:"carid"`
	Carnum            string  `json:"carnum"`
	Carnumberfont     int     `json:"carnumberfont"`
	Carnumberslant    int     `json:"carnumberslant"`
	Carsponsor1       int     `json:"carsponsor1"`
	Carsponsor2       int     `json:"carsponsor2"`
	CcName            string  `json:"ccName"`
	CcNameShort       string  `json:"ccNameShort"`
	Champpoints       int     `json:"champpoints"`
	Classinterval     int     `json:"classinterval"`
	Clubid            int     `json:"clubid"`
	Clubname          string  `json:"clubname"`
	Clubpoints        int     `json:"clubpoints"`
	Clubshortname     string  `json:"clubshortname"`
	Custid            int     `json:"custid"`
	Displayname       string  `json:"displayname"`
	Division          int     `json:"division"`
	Divisionname      string  `json:"divisionname"`
	Dropracepoints    int     `json:"dropracepoints"`
	Evttypename       string  `json:"evttypename"`
	Finishpos         int     `json:"finishpos"`
	Finishposinclass  int     `json:"finishposinclass"`
	Groupid           int     `json:"groupid"`
	Heatinfoid        int     `json:"heatinfoid"`
	HelmColor1        string  `json:"helm_color1"`
	HelmColor2        string  `json:"helm_color2"`
	HelmColor3        string  `json:"helm_color3"`
	HelmPattern       int     `json:"helm_pattern"`
	Hostid            string  `json:"hostid"`
	Incidents         int     `json:"incidents"`
	Interval          int     `json:"interval"`
	Lapscomplete      int     `json:"lapscomplete"`
	Lapslead          int     `json:"lapslead"`
	LeaguePoints      string  `json:"league_points"`
	LicenseChangeOval int     `json:"license_change_oval"`
	LicenseChangeRoad int     `json:"license_change_road"`
	Licensecategory   string  `json:"licensecategory"`
	Licensegroup      int     `json:"licensegroup"`
	MaxPctFuelFill    int     `json:"max_pct_fuel_fill"`
	Multiplier        int     `json:"multiplier"`
	Newcpi            float64 `json:"newcpi"`
	Newirating        int     `json:"newirating"`
	Newlicenselevel   int     `json:"newlicenselevel"`
	Newsublevel       int     `json:"newsublevel"`
	Newttrating       int     `json:"newttrating"`
	Officialsession   int     `json:"officialsession"`
	Oldcpi            float64 `json:"oldcpi"`
	Oldirating        int     `json:"oldirating"`
	Oldlicenselevel   int     `json:"oldlicenselevel"`
	Oldsublevel       int     `json:"oldsublevel"`
	Oldttrating       int     `json:"oldttrating"`
	Optlapscomplete   int     `json:"optlapscomplete"`
	Pos               int     `json:"pos"`
	Quallaptime       int     `json:"quallaptime"`
	Reasonout         string  `json:"reasonout"`
	Reasonoutid       int     `json:"reasonoutid"`
	Restrictresults   string  `json:"restrictresults"`
	Sessionstarttime  int     `json:"sessionstarttime"`
	Simsesname        string  `json:"simsesname"`
	Simsesnum         int     `json:"simsesnum"`
	Simsestypename    string  `json:"simsestypename"`
	Startpos          int     `json:"startpos"`
	SuitColor1        string  `json:"suit_color1"`
	SuitColor2        string  `json:"suit_color2"`
	SuitColor3        string  `json:"suit_color3"`
	SuitPattern       int     `json:"suit_pattern"`
	TrackCategory     string  `json:"track_category"`
	TrackCatid        int     `json:"track_catid"`
	Vehiclekeyid      int     `json:"vehiclekeyid"`
	WeightPenaltyKg   int     `json:"weight_penalty_kg"`
	WheelChrome       int     `json:"wheel_chrome"`
	WheelColor        string  `json:"wheel_color"`
}
