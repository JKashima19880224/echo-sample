package response

type MusicianList struct {
	TotalCount int        `json:"total_count"`
	Items      []Musician `json:"items"`
}

type Musician struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
