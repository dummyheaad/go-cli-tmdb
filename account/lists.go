package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type listsResults struct {
	Description   string      `json:"description"`
	FavoriteCount int         `json:"favorite_count"`
	ID            int         `json:"id"`
	ItemCount     int         `json:"item_count"`
	Iso6391       string      `json:"iso_639_1"`
	ListType      string      `json:"list_type"`
	Name          string      `json:"name"`
	PosterPath    interface{} `json:"poster_path"`
}

type ListsResponse struct {
	Page         int            `json:"page"`
	Results      []listsResults `json:"results"`
	TotalPages   int            `json:"total_pages"`
	TotalResults int            `json:"total_results"`
}

func GetLists(url string, page int) (*ListsResponse, error) {

	u := fmt.Sprintf("%s/lists?page=%d", url, page)

	respByte, err := sendRequest(u, http.MethodGet, "", http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	var resp *ListsResponse
	if err := json.NewDecoder(bytes.NewReader(respByte)).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}
