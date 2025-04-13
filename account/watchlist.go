package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AddWatchlistResponse struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

type watchlistMoviesResults struct {
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
}

type WatchlistMoviesResponse struct {
	Page         int                      `json:"page"`
	Results      []watchlistMoviesResults `json:"results"`
	TotalPages   int                      `json:"total_pages"`
	TotalResults int                      `json:"total_results"`
}

type watchlistTvResults struct {
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
}

type WatchlistTvResponse struct {
	Page         int                  `json:"page"`
	Results      []watchlistTvResults `json:"results"`
	TotalPages   int                  `json:"total_pages"`
	TotalResults int                  `json:"total_results"`
}

func AddWatchlist(url, mediaType string, mediaID int, watchlist bool) (*AddWatchlistResponse, error) {

	u := fmt.Sprintf("%s/watchlist", url)

	media := struct {
		MediaType string `json:"media_type"`
		MediaID   int    `json:"media_id"`
		Watchlist bool   `json:"watchlist"`
	}{
		MediaType: mediaType,
		MediaID:   mediaID,
		Watchlist: watchlist,
	}

	var body bytes.Buffer

	if err := json.NewEncoder(&body).Encode(media); err != nil {
		return nil, err
	}

	statusCode := http.StatusOK
	if watchlist {
		statusCode = http.StatusCreated
	}

	respByte, err := sendRequest(u, http.MethodPost, "application/json", statusCode, &body)
	if err != nil {
		return nil, err
	}

	var resp *AddWatchlistResponse
	if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func GetWatchlist(url, mediaType string) (any, error) {

	// TODO: handle query params
	u := fmt.Sprintf("%s/watchlist/%s?language=en-US&page=1&sort_by=created_at.asc", url, mediaType)

	respByte, err := sendRequest(u, http.MethodGet, "", http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	if mediaType == "movies" {
		var resp WatchlistMoviesResponse
		if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
			return nil, err
		}

		return resp, nil
	}

	var resp WatchlistTvResponse
	if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}
