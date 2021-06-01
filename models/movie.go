package models

type MovieDetail struct {
	ID           int      `json:"id_movie"`
	Name         string   `json:"name"`
	Actors       []People `json:"actors"`
	Directors    []People `json:"directors"`
	Year         int      `json:"year"`
	Rate         int      `json:"rate"`
	Genres       []People `json:"genres"`
	Description  string   `json:"description"`
	NumberSeason int      `json:"no_ss"`
	IsSeries     bool     `json:"is_series"`
	MyList       bool     `json:"my_list"`
}

// use for directorList actorList generesList browseList
type MovieList struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Movie []People `json:"movie"`
}

type People struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PeoplePoster struct {
	ID        int    `json:"id_movie"`
	Name      string `json:"name"`
	PosterURL string `json:"poster_url"`
}

func (m *MovieDetail) AddActor(s People) {
	m.Actors = append(m.Actors, s)
}

func (m *MovieDetail) AddDirector(s People) {
	m.Directors = append(m.Directors, s)
}

func (m *MovieDetail) AddGenres(s People) {
	m.Genres = append(m.Genres, s)
}

func (m *MovieList) AddMovie(s People) {
	m.Movie = append(m.Movie, s)
}
