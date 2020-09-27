package model

const (
	linksLimit = 20
)

//LinksRequest contains list of URLs.
type LinksRequest []string

//LinkResponse contains details with URL response.
type LinkResponse struct {
	URL    string `json:"url"`
	Status int    `json:"status"`
}

//LinksResponse collection of URLs responses.
type LinksResponse []LinkResponse

//ExceedsLimit checks if collection exceeds defined limit.
func (lnks LinksRequest) ExceedsLimit() bool {
	return len(lnks) > linksLimit
}

//IsEmpty checks if collection is empty.
func (lnks LinksRequest) IsEmpty() bool {
	return len(lnks) == 0
}
