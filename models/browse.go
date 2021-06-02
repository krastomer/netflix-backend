package models

type MovieHistory struct {
	IDHistory int    `json:"id_history"`
	IDEpisode int    `json:"id_episode"`
	StopTime  string `json:"stop_time"`
	Name      string `json:"name"`
	IDMovie   int    `json:"id_movie"`
	Rate      int    `json:"rate"`
	Year      int    `json:"year"`
	IsSeries  bool   `json:"is_series"`
	PosterURL string `json:"poster_url"`
}

type MovieSeq struct {
	Seq       int    `json:"no"`
	IDMovie   int    `json:"id_movie"`
	Name      string `json:"name"`
	PosterURL string `json:"poster_url"`
	NViews    int    `json:"n_views"`
}
