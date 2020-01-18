package tmdb

type ExternalFindResults struct {
	TvResults []struct {
		OriginalName     string   `json:"original_name"`
		ID               int      `json:"id"`
		Name             string   `json:"name"`
		VoteCount        int      `json:"vote_count"`
		VoteAverage      float64  `json:"vote_average"`
		FirstAirDate     string   `json:"first_air_date"`
		PosterPath       string   `json:"poster_path"`
		GenreIds         []int    `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
		Popularity       float64  `json:"popularity"`
	} `json:"tv_results"`
	MovieResults []struct {
		ID               int     `json:"id"`
		Video            bool    `json:"video"`
		VoteCount        int     `json:"vote_count"`
		VoteAverage      float64 `json:"vote_average"`
		Title            string  `json:"title"`
		ReleaseDate      string  `json:"release_date"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		GenreIds         []int   `json:"genre_ids"`
		BackdropPath     string  `json:"backdrop_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		Popularity       float64 `json:"popularity"`
	} `json:"movie_results"`
}

type ShowInfo struct {
	BackdropPath string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int    `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"created_by"`
	EpisodeRunTime []int  `json:"episode_run_time"`
	FirstAirDate   string `json:"first_air_date"`
	Genres         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int      `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int     `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int     `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int     `json:"vote_count"`
	} `json:"last_episode_to_air"`
	Name             string      `json:"name"`
	NextEpisodeToAir interface{} `json:"next_episode_to_air"`
	Networks         []struct {
		Name          string `json:"name"`
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string   `json:"overview"`
	Popularity          float64  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	Seasons []struct {
		AirDate      string      `json:"air_date"`
		EpisodeCount int         `json:"episode_count"`
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		Overview     string      `json:"overview"`
		PosterPath   interface{} `json:"poster_path"`
		SeasonNumber int         `json:"season_number"`
	} `json:"seasons"`
	Status      string  `json:"status"`
	Type        string  `json:"type"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}
type episodeInfoContainer struct {
	ID           string        `json:"_id"`
	AirDate      string        `json:"air_date"`
	Name         string        `json:"name"`
	Overview     string        `json:"overview"`
	PosterPath   string        `json:"poster_path"`
	SeasonNumber int           `json:"season_number"`
	Episodes     []EpisodeInfo `json:"episodes"`
}

type EpisodeInfo struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	Crew           []struct {
		ID          int    `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Department  string `json:"department"`
		Job         string `json:"job"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
	GuestStars []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		CreditID    string `json:"credit_id"`
		Character   string `json:"character"`
		Order       int    `json:"order"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"guest_stars"`
}
