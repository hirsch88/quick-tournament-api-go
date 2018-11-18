package dtos

// API struct is the construction for the public
// information about the API.
type API struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
