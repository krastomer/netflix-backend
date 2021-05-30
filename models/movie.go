package models

type MovieDetail struct {
	Name         string   `json:"name"`
	Actors       []string `json:"actors"`
	Directors    []string `json:"directors"`
	Year         int      `json:"year"`
	Rate         int      `json:"rate"`
	Genres       []string `json:"genres"`
	Description  string   `json:"description"`
	NumberSeason int      `json:"no_ss"`
	IsSeries     bool     `json:"is_series"`
}

func (m *MovieDetail) AddActor(s string) {
	m.Actors = append(m.Actors, s)
}

func (m *MovieDetail) AddDirector(s string) {
	m.Directors = append(m.Directors, s)
}

func (m *MovieDetail) AddGenres(s string) {
	m.Genres = append(m.Genres, s)
}
