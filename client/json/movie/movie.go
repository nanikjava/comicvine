package movie

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

type Producer struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Studio struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Writer struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Result struct {
	APIDetailURL     string      `json:"api_detail_url"`
	BoxOfficeRevenue string      `json:"box_office_revenue"`
	Budget           string      `json:"budget"`
	DateAdded        string      `json:"date_added"`
	DateLastUpdated  string      `json:"date_last_updated"`
	Deck             string      `json:"deck"`
	Description      string      `json:"description"`
	Distributor      interface{} `json:"distributor"`
	HasStaffReview   interface{} `json:"has_staff_review"`
	ID               int         `json:"id"`
	Image            Image       `json:"image"`
	Name             string      `json:"name"`
	Producers        []Producer  `json:"producers"`
	Rating           string      `json:"rating"`
	ReleaseDate      string      `json:"release_date"`
	Runtime          string      `json:"runtime"`
	SiteDetailURL    string      `json:"site_detail_url"`
	Studios          []Studio    `json:"studios"`
	TotalRevenue     string      `json:"total_revenue"`
	Writers          []Writer    `json:"writers"`
}

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
