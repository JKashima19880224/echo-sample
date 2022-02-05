package request

type MusicianList struct {
	PerPage int `query:"per_page"`
	Page    int `query:"page"`
}
