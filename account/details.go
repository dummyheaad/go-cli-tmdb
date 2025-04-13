package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var (
	ErrConnection      = errors.New("connection error")
	ErrNotFound        = errors.New("not found")
	ErrInvalidResponse = errors.New("invalid server response")
	ErrInvalid         = errors.New("invalid data")
	ErrNotNumber       = errors.New("not a number")
)

type gravatar struct {
	Hash string `json:"hash"`
}

type tmdb struct {
	AvatarPath string `json:"avatar_path"`
}

type avatar struct {
	Gravatar gravatar `json:"gravatar"`
	Tmdb     tmdb     `json:"tmdb"`
}

type DetailsResponse struct {
	Avatar       avatar `json:"avatar"`
	ID           int    `json:"id"`
	ISO_639_1    string `json:"iso_639_1"`
	ISO_3166_1   string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

func newClient() *http.Client {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	return c
}

func GetDetails(url string) (*DetailsResponse, error) {

	respByte, err := sendRequest(url, http.MethodGet, "", http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	var resp *DetailsResponse
	if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}
