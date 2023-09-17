package power

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
	Aliases         interface{} `json:"aliases"`
	APIDetailURL    string      `json:"api_detail_url"`
	DateAdded       string      `json:"date_added"`
	DateLastUpdated string      `json:"date_last_updated"`
	Description     string      `json:"description"`
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	SiteDetailURL   string      `json:"site_detail_url"`
}
