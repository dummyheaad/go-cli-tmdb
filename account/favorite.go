package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type AddFavoriteResponse struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

type favMovieResults struct {
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

type FavoriteMoviesResponse struct {
	Page         int               `json:"page"`
	Results      []favMovieResults `json:"results"`
	TotalPages   int               `json:"total_pages"`
	TotalResults int               `json:"total_results"`
}

type favTvResults struct {
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

type FavoriteTvResponse struct {
	Page         int            `json:"page"`
	Results      []favTvResults `json:"results"`
	TotalPages   int            `json:"total_pages"`
	TotalResults int            `json:"total_results"`
}

func sendRequest(url, method, contentType string,
	expStatus int, body io.Reader) ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	authToken := os.Getenv("AUTH_TOKEN")

	req, err := http.NewRequest(method, url, body)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+authToken)
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	r, err := newClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != expStatus {
		msg, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("cannot read body: %w", err)
		}
		err = ErrInvalidResponse
		if r.StatusCode == http.StatusNotFound {
			err = ErrNotFound
		}

		return nil, fmt.Errorf("%w: %s", err, msg)
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %w", err)
	}

	return resp, nil
}

func AddFavorite(url, mediaType string, mediaID int, favorite bool) (*AddFavoriteResponse, error) {

	u := fmt.Sprintf("%s/favorite", url)

	media := struct {
		MediaType string `json:"media_type"`
		MediaID   int    `json:"media_id"`
		Favorite  bool   `json:"favorite"`
	}{
		MediaType: mediaType,
		MediaID:   mediaID,
		Favorite:  favorite,
	}

	var body bytes.Buffer

	if err := json.NewEncoder(&body).Encode(media); err != nil {
		return nil, err
	}

	statusCode := http.StatusOK
	if favorite {
		statusCode = http.StatusCreated
	}

	respByte, err := sendRequest(u, http.MethodPost, "application/json", statusCode, &body)
	if err != nil {
		return nil, err
	}

	var resp *AddFavoriteResponse
	if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func GetFavorite(url, mediaType string) (any, error) {

	// TODO: handle query params
	u := fmt.Sprintf("%s/favorite/%s?language=en-US&page=1&sort_by=created_at.asc", url, mediaType)

	respByte, err := sendRequest(u, http.MethodGet, "", http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	if mediaType == "movies" {
		var resp FavoriteMoviesResponse
		if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
			return nil, err
		}

		return resp, nil
	}

	var resp FavoriteTvResponse
	if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}
