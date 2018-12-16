package iracing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"
)

const memberURL = "https://members.iracing.com/membersite/member/GetMember"

// GetSelf returns data about the currently-logged-in user.
func (c Client) GetSelf() (*Member, error) {
	res, err := c.Get(memberURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var m Member
	return &m, json.NewDecoder(res.Body).Decode(&m)
}

var (
	memberRE  = regexp.MustCompile(`(?m)^buf = '(.*?)';$`)
	profileRE = regexp.MustCompile(`(?m)^var buf = '({.*?})';$`)
)

// GetMember returns user data about the requested customer ID.
func (c Client) GetMember(custID int) (*Member, error) {
	res, err := c.Get(fmt.Sprintf("https://members.iracing.com/membersite/member/CareerStats.do?custid=%d", custID))
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	matches := memberRE.FindAllSubmatch(bs, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("could not find member json in HTML body")
	}

	profMatches := profileRE.FindAllSubmatch(bs, -1)
	if len(profMatches) == 0 {
		return nil, fmt.Errorf("could not find member profile json in HTML body")
	}

	var m Member
	err = json.Unmarshal(matches[0][1], &m)
	if err != nil {
		return nil, err
	}

	var p struct{ Profile Profile }
	err = json.Unmarshal(profMatches[0][1], &p)
	if err != nil {
		return nil, err
	}

	m.Profile = p.Profile

	return &m, nil
}

type DMY struct {
	time.Time
}

func (m *DMY) UnmarshalJSON(bs []byte) error {
	bs = bytes.Trim(bs, "\"")

	var err error
	m.Time, err = time.Parse("02-01-2006", string(bs))
	return err
}

// Member holds data about the current user.
type Member struct {
	AllowedSeasons []interface{} `json:"allowedSeasons"`
	ClubID         int           `json:"clubId"`
	CustID         int           `json:"custId"`
	DisplayName    String        `json:"displayName"`
	FavCar         int           `json:"favCar"`
	FavTrack       int           `json:"favTrack"`
	Friend         bool          `json:"friend"`
	HasReadPP      bool          `json:"hasReadPP"`
	HasReadTC      bool          `json:"hasReadTC"`
	Helmet         Helmet        `json:"helmet"`
	Licenses       []License     `json:"licenses"`
	MemberSince    DMY           `json:"memberSince"`
	Profile        Profile
	ReadPP         String `json:"readPP"`
	ReadTC         String `json:"readTC"`
	Watched        bool   `json:"watched"`
}

// Helmet holds data about the helmet color/etc.
type Helmet struct {
	C1 string `json:"c1"`
	C2 string `json:"c2"`
	C3 string `json:"c3"`
	Hp int    `json:"hp"`
	Ll int    `json:"ll"`
}

type License struct {
	CatID               Category  `json:"catId"`
	IRating             ScrewyInt `json:"iRating"`
	LicColor            string    `json:"licColor"`
	LicGroup            int       `json:"licGroup"`
	LicGroupDisplayName String    `json:"licGroupDisplayName"`
	LicLevel            int       `json:"licLevel"`
	LicLevelDisplayName String    `json:"licLevelDisplayName"`
	MprNumRaces         int       `json:"mprNumRaces"`
	MprNumTTs           int       `json:"mprNumTTs"`
	SrPrime             int       `json:"srPrime,string"`
	SrSub               int       `json:"srSub,string"`
	SubLevel            int       `json:"subLevel"`
	TtRating            ScrewyInt `json:"ttRating"`
}

type ScrewyInt int

func (s *ScrewyInt) UnmarshalJSON(bs []byte) error {
	if bytes.Equal(bs, []byte(`"---"`)) {
		return nil // leave as zero
	}

	var i int
	err := json.Unmarshal(bs, &i)
	if err != nil {
		return err
	}
	*s = ScrewyInt(i)
	return nil
}

// Profile is a member's profile data
type Profile struct {
	AIM                string `json:"AIM"`
	Bio                String `json:"BIO"`
	Birthdate          String `json:"BIRTHDATE"`
	ClubID             int    `json:"CLUBID,string"`
	ClubRegionImg      String `json:"CLUB_REGION_IMG"`
	Email              String `json:"EMAIL"`
	FavBooks           String `json:"FAV_BOOKS"`
	FavGames           String `json:"FAV_GAMES"`
	FavHobbiES         String `json:"FAV_HOBBIES"`
	FavIRacingCars     String `json:"FAV_IRacing_CARS"`
	FavIRacingSeries   String `json:"FAV_IRacing_SERIES"`
	FavIRacingTracks   String `json:"FAV_IRacing_TRACKS"`
	FavMovies          String `json:"FAV_MOVIES"`
	FavMusic           String `json:"FAV_MUSIC"`
	FavQuotation       String `json:"FAV_QUOTATION"`
	FavRealCars        String `json:"FAV_REAL_CARS"`
	FavRealTrackS      String `json:"FAV_REAL_TRACKS"`
	FavSports          String `json:"FAV_SPORTS"`
	FavTvshows         String `json:"FAV_TV_SHOWS"`
	FavWebsite         String `json:"FAV_WEBSITE"`
	Hometown           String `json:"HOMETOWN"`
	LastLogin          Time   `json:"LASTLOGIN"`
	Location           String `json:"LOCATION"`
	MemberSince        Time   `json:"MemberSINCE"`
	MemberPhotoHeight  String `json:"Member_PHOTO_HEIGHT"`
	MemberPhotoOffsetX String `json:"Member_PHOTO_OFFSET_X"`
	MemberPhotoOffsetY String `json:"Member_PHOTO_OFFSET_Y"`
	MemberPhotoWidth   String `json:"Member_PHOTO_WIDTH"`
	Name               String `json:"NAME"`
	Nickname           String `json:"NICKNAME"`
	Web1               String `json:"WEB1"`
	Web2               String `json:"WEB2"`
	Web4               String `json:"WEB4"`
}
