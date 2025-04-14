package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ratedMoviesResults struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	Rating           float64 `json:"rating"`
}

type RatedMoviesResponse struct {
	Page         int                  `json:"page"`
	Results      []ratedMoviesResults `json:"results"`
	TotalPages   int                  `json:"total_pages"`
	TotalResults int                  `json:"total_results"`
}

type ratedTvResults struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIds         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	FirstAirDate     string   `json:"first_air_date"`
	Name             string   `json:"name"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Rating           float64  `json:"rating"`
}

type RatedTvResponse struct {
	Page         int              `json:"page"`
	Results      []ratedTvResults `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

type ratedTvEpisodeResults struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	EpisodeType    string  `json:"episode_type"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	Rating         float64 `json:"rating"`
}

type RatedTvEpisodeResponse struct {
	Page         int                     `json:"page"`
	Results      []ratedTvEpisodeResults `json:"results"`
	TotalPages   int                     `json:"total_pages"`
	TotalResults int                     `json:"total_results"`
}

func GetRatedEpisodes(url string) (*RatedTvEpisodeResponse, error) {

	// TODO: handle query params
	u := fmt.Sprintf("%s/rated/tv/episodes?language=en-US&page=1&sort_by=created_at.asc", url)

	respByte, err := sendRequest(u, http.MethodGet, "", http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	var resp *RatedTvEpisodeResponse
	if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func GetRatedShow[T *RatedMoviesResponse | *RatedTvResponse](url, mediaType string) (T, error) {
	var resp T

	// TODO: handle query params
	u := fmt.Sprintf("%s/rated/%s?language=en-US&page=1&sort_by=created_at.asc", url, mediaType)

	respByte, err := sendRequest(u, http.MethodGet, "", http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}
