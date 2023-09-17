package people

type MainType struct {
	Error                string   `json:"error"`
	Limit                int      `json:"limit"`
	Offset               int      `json:"offset"`
	NumberOfPageResults  int      `json:"number_of_page_results"`
	NumberOfTotalResults int      `json:"number_of_total_results"`
	StatusCode           int      `json:"status_code"`
	Results              []Result `json:"results"`
	Version              string   `json:"version"`
}

type Result struct {
	Aliases                 interface{} `json:"aliases"`
	APIDetailURL            string      `json:"api_detail_url"`
	Birth                   interface{} `json:"birth"`
	CountOfIssueAppearances interface{} `json:"count_of_isssue_appearances"`
	Country                 interface{} `json:"country"`
	DateAdded               string      `json:"date_added"`
	DateLastUpdated         string      `json:"date_last_updated"`
	Death                   interface{} `json:"death"`
	Deck                    interface{} `json:"deck"`
	Description             interface{} `json:"description"`
	Email                   interface{} `json:"email"`
	Gender                  int         `json:"gender"`
	Hometown                interface{} `json:"hometown"`
	ID                      int         `json:"id"`
	Image                   Image       `json:"image"`
	Name                    string      `json:"name"`
	SiteDetailURL           string      `json:"site_detail_url"`
	Website                 interface{} `json:"website"`
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
