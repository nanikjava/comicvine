package character

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

type Origin struct {
	APIDetailURL string `json:"api_detail_url"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
}

type Publisher struct {
	APIDetailURL string `json:"api_detail_url"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
}

type Results struct {
	Aliases                 string               `json:"aliases"`
	APIDetailURL            string               `json:"api_detail_url"`
	Birth                   any                  `json:"birth"`
	CountOfIssueAppearances int                  `json:"count_of_issue_appearances"`
	DateAdded               string               `json:"date_added"`
	DateLastUpdated         string               `json:"date_last_updated"`
	Deck                    string               `json:"deck"`
	Description             string               `json:"description"`
	FirstAppearedInIssue    FirstAppearedInIssue `json:"first_appeared_in_issue"`
	Gender                  int                  `json:"gender"`
	ID                      int                  `json:"id"`
	Image                   Image                `json:"image"`
	Name                    string               `json:"name"`
	Origin                  Origin               `json:"origin"`
	Publisher               Publisher            `json:"publisher"`
	RealName                string               `json:"real_name"`
	SiteDetailURL           string               `json:"site_detail_url"`
}

type MainType struct {
	Error                string    `json:"error"`
	Limit                int       `json:"limit"`
	Offset               int       `json:"offset"`
	NumberOfPageResults  int       `json:"number_of_page_results"`
	NumberOfTotalResults int       `json:"number_of_total_results"`
	StatusCode           int       `json:"status_code"`
	Results              []Results `json:"results"`
	Version              string    `json:"version"`
}
