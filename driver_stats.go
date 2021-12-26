package iracing

import "fmt"

// GetStats returns statistics for a given driver
func (c Client) GetStats(custID int) (*Stats, error) {
	var s Stats
	err := c.JSON(fmt.Sprintf("http://members.iracing.com./memberstats/member/GetCareerStats?custid=%d", custID), &s.CareerStats)
	if err != nil {
		return nil, err
	}

	err = c.JSON(fmt.Sprintf("http://members.iracing.com./memberstats/member/GetYearlyStats?custid=%d", custID), &s.YearlyStats)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Stats is stats, but for drivers.
type Stats struct {
	CareerStats []DriverStats
	YearlyStats []DriverStats
}

// DriverStats encapsulates yearly/career stats. Year is empty for career
// stats.
type DriverStats struct {
	AvgFinish       int     `json:"avgFinish"`
	AvgIncPerRace   float64 `json:"avgIncPerRace"`
	AvgPtsPerRace   int     `json:"avgPtsPerRace"`
	AvgStart        int     `json:"avgStart"`
	Category        String  `json:"category"`
	Clubpoints      int     `json:"clubpoints"`
	LapsLed         int     `json:"lapsLed"`
	LapsLedPerc     float64 `json:"lapsLedPerc"`
	Poles           int     `json:"poles"`
	Starts          int     `json:"starts"`
	Top5            int     `json:"top5"`
	Top5Perc        float64 `json:"top5Perc"`
	TotalLaps       int     `json:"totalLaps"`
	TotalClubpoints int     `json:"totalclubpoints"`
	WinPerc         float64 `json:"winPerc"`
	Wins            int     `json:"wins"`
	Year            int     `json:"year,string"`
}
