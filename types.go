package paginator

type links struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}
type Meta struct {
	PerPage     int    `json:"per_page"`
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	From        int    `json:"from"`
	To          int    `json:"to"`
	Path        string `json:"path"`
	Total       int    `json:"total"`
}
