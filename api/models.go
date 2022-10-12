package api

type Result struct {
	Title   string `json:"title"`
	Alt     string `json:"alt"`
	Price   string `json:"price"`
	Region  string `json:"region"`
	Preview string `json:"preview"`
	Link    string `json:"link"`
}
