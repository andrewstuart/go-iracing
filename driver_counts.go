package iracing

import "encoding/json"

func (c *Client) JSON(u string, out interface{}) error {
	res, err := c.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&out)
}

func (c *Client) GetDriverCounts() (*Counts, error) {
	var ct Counts
	return &ct, c.JSON("http://members.iracing.com/membersite/member/GetDriverCounts?invokedby=racepanel", &ct)
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
