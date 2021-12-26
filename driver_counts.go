package iracing

func (c *Client) GetDriverCounts() (*Counts, error) {
	var ct Counts
	return &ct, c.JSON("http://members.iracing.com./membersite/member/GetDriverCounts?invokedby=racepanel", &ct)
}

type Counts struct {
	DriverCounts       `json:"driverCounts"`
	FriendRequests     bool          `json:"friendRequests"`
	Newnotifications   []interface{} `json:"newnotifications"`
	PmFull             bool          `json:"pmFull"`
	UnreadPMCount      int           `json:"unreadPMCount"`
	UnviewedAwardCount int           `json:"unviewedAwardCount"`
}

type DriverCounts struct {
	LapCount String `json:"lapCount"`
	Myracers int    `json:"myracers"`
	Total    int    `json:"total"`
}
