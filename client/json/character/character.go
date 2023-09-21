package character

type MainType struct {
	Error                string  `json:"error"`
	Limit                int     `json:"limit"`
	Offset               int     `json:"offset"`
	NumberOfPageResults  int     `json:"number_of_page_results"`
	NumberOfTotalResults int     `json:"number_of_total_results"`
	StatusCode           int     `json:"status_code"`
	Results              Results `json:"results"`
	Version              string  `json:"version"`
}

type Results struct {
	Aliases                 string               `json:"aliases"`
	APIDetailURL            string               `json:"api_detail_url"`
	Birth                   any                  `json:"birth"`
	CharacterEnemies        []CharacterEnemy     `json:"character_enemies"`
	CharacterFriends        []CharacterFriend    `json:"character_friends"`
	CountOfIssueAppearances int                  `json:"count_of_issue_appearances"`
	Creators                []Creator            `json:"creators"`
	DateAdded               string               `json:"date_added"`
	DateLastUpdated         string               `json:"date_last_updated"`
	Deck                    string               `json:"deck"`
	Description             string               `json:"description"`
	FirstAppearedInIssue    FirstAppearedInIssue `json:"first_appeared_in_issue"`
	Gender                  int                  `json:"gender"`
	ID                      int                  `json:"id"`
	Image                   Image                `json:"image"`
	IssueCredits            []IssueCredit        `json:"issue_credits"`
	IssuesDiedIn            []any                `json:"issues_died_in"`
	Movies                  []Movie              `json:"movies"`
	Name                    string               `json:"name"`
	Origin                  Origin               `json:"origin"`
	Powers                  []Power              `json:"powers"`
	Publisher               Publisher            `json:"publisher"`
	RealName                string               `json:"real_name"`
	SiteDetailURL           string               `json:"site_detail_url"`
	StoryArcCredits         []any                `json:"story_arc_credits"`
	TeamEnemies             []TeamEnemy          `json:"team_enemies"`
	TeamFriends             []TeamFriend         `json:"team_friends"`
	Teams                   []Team               `json:"teams"`
	VolumeCredits           []VolumeCredit       `json:"volume_credits"`
}

type CharacterEnemy struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type CharacterFriend struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Creator struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type FirstAppearedInIssue struct {
	APIDetailURL string `json:"api_detail_url"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IssueNumber  string `json:"issue_number"`
}

type Image struct {
	IconURL        string `json:"icon_url"`
	MediumURL      string `json:"medium_url"`
	ScreenURL      string `json:"screen_url"`
	ScreenLargeURL string `json:"screen_large_url"`
	SmallURL       string `json:"small_url"`
	SuperURL       string `json:"super_url"`
	ThumbURL       string `json:"thumb_url"`
	TinyURL        string `json:"tiny_url"`
	OriginalURL    string `json:"original_url"`
	ImageTags      string `json:"image_tags"`
}

type IssueCredit struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Movie struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Origin struct {
	APIDetailURL string `json:"api_detail_url"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
}

type Power struct {
	APIDetailURL string `json:"api_detail_url"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
}

type Publisher struct {
	APIDetailURL string `json:"api_detail_url"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
}

type TeamEnemy struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type TeamFriend struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Team struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type VolumeCredit struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}
