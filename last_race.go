package iracing

import "fmt"

// GetLastRaces returns the last races for a Member.
func (c Client) GetLastRaces(custID int) ([]LastRaceStats, error) {
	var l []LastRaceStats
	return l, c.JSON(fmt.Sprintf("https://members.iracing.com/memberstats/member/GetLastRacesStats?custid=%d", custID), &l)
}

type LastRaceStats struct {
	CarClassID           int    `json:"carClassID"`
	CarColor1            string `json:"carColor1"`
	CarColor2            string `json:"carColor2"`
	CarColor3            string `json:"carColor3"`
	CarID                int    `json:"carID"`
	CarPattern           int    `json:"carPattern"`
	ChampPoints          int    `json:"champPoints"`
	ClubPoints           int    `json:"clubPoints"`
	Date                 string `json:"date"`
	FinishPos            int    `json:"finishPos"`
	Incidents            int    `json:"incidents"`
	Laps                 int    `json:"laps"`
	LapsLed              int    `json:"lapsLed"`
	LicenseLevel         int    `json:"licenseLevel"`
	QualifyTime          int    `json:"qualifyTime"`
	QualifyTimeFormatted String `json:"qualifyTimeFormatted"`
	SeasonID             int    `json:"seasonID"`
	SeriesID             int    `json:"seriesID"`
	Sof                  int    `json:"sof"`
	StartPos             int    `json:"startPos"`
	SubsessionID         int    `json:"subsessionID"`
	Time                 Time   `json:"time"`
	TrackID              int    `json:"trackID"`
	TrackName            String `json:"trackName"`
	WinnerHC1            string `json:"winnerHC1"`
	WinnerHC2            string `json:"winnerHC2"`
	WinnerHC3            string `json:"winnerHC3"`
	WinnerHPattern       int    `json:"winnerHPattern"`
	WinnerID             int    `json:"winnerID"`
	WinnerLL             int    `json:"winnerLL"`
	WinnerName           String `json:"winnerName"`
}
