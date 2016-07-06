package gocensys

//Search endpoint allows searches against the current data in the IPv4, Top Million Websites, and Certificates indexes
//using the same search syntax as the primary site. The endpoint returns
//a paginated result set of hosts (or websites or certificates) that match the search.
func (c CensysAPI) Search(index string, query map[string]interface{}) (SearchResults, error) {
	var searchResults SearchResults

	err := c.apiPost("/search/"+index, query, &searchResults)

	return searchResults, err
}

//SearchResults represents the response from a search request.
type SearchResults struct {
	Status   string        `json:"status"`
	Results  []interface{} `json:"results"`
	Metadata struct {
		Count       int    `json:"count"`
		Query       string `json:"query"`
		BackendTime int    `json:"backend_time"`
		Page        int    `json:"page"`
		Pages       int    `json:"pages"`
	} `json:"metadata"`
}
