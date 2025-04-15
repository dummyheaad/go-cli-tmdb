package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestDetailsAction(t *testing.T) {
	testCases := []struct {
		name      string
		accountID string
		expError  error
		expOut    string
		resp      struct {
			Status int
			Body   string
		}
		closeServer bool
		isRaw       bool
	}{
		{
			name:      "Details",
			accountID: "null",
			expError:  nil,
			expOut:    "Account details for 21907685\nID: 21907685\nUsername: clairvoyance27\n",
			resp:      testResp["resultsDetails"],
			isRaw:     false,
		},
		{
			name:      "DetailsRaw",
			accountID: "null",
			expError:  nil,
			expOut:    "{\n   \"avatar\": {\n      \"gravatar\": {\n         \"hash\": \"5a33321a08977fbf047ab4d39105637a\"\n      },\n      \"tmdb\": {\n         \"avatar_path\": \"/yUaRo4KmeADP7lkAS0t9p7r36yQ.jpg\"\n      }\n   },\n   \"id\": 21907685,\n   \"iso_639_1\": \"en\",\n   \"iso_3166_1\": \"ID\",\n   \"name\": \"Uka\",\n   \"include_adult\": false,\n   \"username\": \"clairvoyance27\"\n}",
			resp:      testResp["resultsDetails"],
			isRaw:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.resp.Status)
					fmt.Fprintln(w, tc.resp.Body)
				})
			defer cleanup()

			if tc.closeServer {
				cleanup()
			}

			var out bytes.Buffer

			err := detailsAction(&out, url, tc.accountID, tc.isRaw)

			if tc.expError != nil {
				if err == nil {
					t.Fatalf("Expected error %q, got no error.", tc.expError)
				}

				if !errors.Is(err, tc.expError) {
					t.Errorf("Expected error %q, got %q.", tc.expError, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q.", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("Expected output %q, got %q.", tc.expOut, out.String())
			}
		})
	}
}

func TestListsAction(t *testing.T) {
	testCases := []struct {
		name     string
		page     []string
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		closeServer bool
		isRaw       bool
	}{
		{
			name:     "Lists",
			page:     []string{"1"},
			expError: nil,
			expOut:   "Lists:\n1. Name: my-list-2\nDescription: test my list 2\nList Type: movie\nTotal Items: 2\n\n2. Name: my-list\nDescription: test my list\nList Type: movie\nTotal Items: 2\n\n",
			resp:     testResp["resultsLists"],
			isRaw:    false,
		},
		{
			name:     "NoLists",
			page:     []string{"500"},
			expError: nil,
			expOut:   "Lists:\n",
			resp:     testResp["resultsNoLists"],
			isRaw:    false,
		},
		{
			name:     "ListsRaw",
			page:     []string{"1"},
			expError: nil,
			expOut:   "{\n   \"page\": 1,\n   \"results\": [\n      {\n         \"description\": \"test my list 2\",\n         \"favorite_count\": 0,\n         \"id\": 8525470,\n         \"item_count\": 2,\n         \"iso_639_1\": \"en\",\n         \"list_type\": \"movie\",\n         \"name\": \"my-list-2\",\n         \"poster_path\": null\n      },\n      {\n         \"description\": \"test my list\",\n         \"favorite_count\": 0,\n         \"id\": 8521773,\n         \"item_count\": 2,\n         \"iso_639_1\": \"en\",\n         \"list_type\": \"movie\",\n         \"name\": \"my-list\",\n         \"poster_path\": null\n      }\n   ],\n   \"total_pages\": 1,\n   \"total_results\": 2\n}",
			resp:     testResp["resultsLists"],
			isRaw:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.resp.Status)
					fmt.Fprintln(w, tc.resp.Body)
				})
			defer cleanup()

			if tc.closeServer {
				cleanup()
			}

			var out bytes.Buffer

			err := listsAction(&out, url, tc.page, tc.isRaw)

			if tc.expError != nil {
				if err == nil {
					t.Fatalf("Expected error %q, got no error.", tc.expError)
				}

				if !errors.Is(err, tc.expError) {
					t.Errorf("Expected error %q, got %q.", tc.expError, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q.", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("Expected output %q, got %q.", tc.expOut, out.String())
			}
		})
	}
}

func TestGetFavoriteAction(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		isRaw    bool
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		closeServer bool
	}{
		{
			name:     "FavMovies",
			args:     []string{"movies"},
			isRaw:    false,
			expError: nil,
			expOut:   "Favorite Movies:\n1. Title: Cosmic Chaos\nRelease Date: 2023-08-03\nPopularity: 160.04\nVote Count: 46\nVote Average: 6.00\n\n2. Title: Absolut\nRelease Date: 2005-04-20\nPopularity: 0.29\nVote Count: 29\nVote Average: 7.80\n\n",
			resp:     testResp["resultsFavMovies"],
		},
		{
			name:     "FavTv",
			args:     []string{"tv"},
			isRaw:    false,
			expError: nil,
			expOut:   "Favorite TV Shows:\n1. Name: Till Death Us Do Part\nFirst Air Date: 1966-06-06\nPopularity: 12.82\nVote Count: 24\nVote Average: 7.40\n\n2. Name: Game of Thrones\nFirst Air Date: 2011-04-17\nPopularity: 192.15\nVote Count: 24838\nVote Average: 8.46\n\n",
			resp:     testResp["resultsFavTv"],
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.resp.Status)
					fmt.Fprintln(w, tc.resp.Body)
				})
			defer cleanup()

			if tc.closeServer {
				cleanup()
			}

			var out bytes.Buffer

			err := getAction(&out, url, tc.args, tc.isRaw)

			if tc.expError != nil {
				if err == nil {
					t.Fatalf("Expected error %q, got no error.", tc.expError)
				}

				if !errors.Is(err, tc.expError) {
					t.Errorf("Expected error %q, got %q.", tc.expError, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q.", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("Expected output %q, got %q.", tc.expOut, out.String())
			}
		})
	}
}

func TestAddFavoriteAction(t *testing.T) {
	expURLPath := "/account/null/favorite"
	expMethod := http.MethodPost
	expBody := "{\"media_type\":\"movie\",\"media_id\":650,\"favorite\":true}\n"
	expContentType := "application/json"
	expOut := "{\n   \"success\": true,\n   \"status_code\": 1,\n   \"status_message\": \"Success.\"\n}"
	args := []string{"movie", "650", "yes"}

	url, cleanup := mockServer(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != expURLPath {
				t.Errorf("Expected path %q, got %q", expURLPath, r.URL.Path)
			}

			if r.Method != expMethod {
				t.Errorf("Expected method %q, got %q", expMethod, r.Method)
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			r.Body.Close()

			if string(body) != expBody {
				t.Errorf("Expected body %q, got %q", expBody, string(body))
			}

			contentType := r.Header.Get("Content-Type")
			if contentType != expContentType {
				t.Errorf("Expected Content-Type %q, got %q", expContentType, contentType)
			}

			w.WriteHeader(testResp["resultsAddFav"].Status)
			fmt.Fprintln(w, testResp["resultsAddFav"].Body)
		})
	defer cleanup()

	var out bytes.Buffer

	if err := addAction(&out, url, args); err != nil {
		t.Fatalf("Expected no error, got %q", err)
	}

	if expOut != out.String() {
		t.Errorf("Expected output %q, got %q", expOut, out.String())
	}
}

func TestGetWatchlistAction(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		isRaw       bool
		closeServer bool
	}{
		{
			name:        "WatchlistMovies",
			args:        []string{"movies"},
			expError:    nil,
			expOut:      "Watchlist Movies:\n1. Title: Star Wars\nRelease Date: 1977-05-25\nPopularity: 20.152500\nVote Count: 21056\nVote Average: 8.203000\n\n2. Title: Fight Club\nRelease Date: 1999-10-15\nPopularity: 33.760600\nVote Count: 30142\nVote Average: 8.438000\n\n3. Title: Reservoir Dogs\nRelease Date: 1992-09-02\nPopularity: 10.440200\nVote Count: 14565\nVote Average: 8.122000\n\n4. Title: Cléo from 5 to 7\nRelease Date: 1962-04-11\nPopularity: 2.470800\nVote Count: 706\nVote Average: 7.700000\n\n",
			resp:        testResp["resultsWatchlistMovies"],
			isRaw:       false,
			closeServer: false,
		},
		{
			name:        "WatchlistTv",
			args:        []string{"tv"},
			expError:    nil,
			expOut:      "Watchlist TV Shows:\n1. Name: Law & Order: Special Victims Unit\nFirst Air Date: 1999-09-20\nPopularity: 341.293000\nVote Count: 3905\nVote Average: 7.938000\n\n2. Name: Island at War\nFirst Air Date: 2004-07-11\nPopularity: 1.683700\nVote Count: 9\nVote Average: 7.400000\n\n",
			resp:        testResp["resultsWatchlistTv"],
			isRaw:       false,
			closeServer: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.resp.Status)
					fmt.Fprintf(w, tc.resp.Body)
				})
			defer cleanup()

			if tc.closeServer {
				cleanup()
			}

			var out bytes.Buffer

			err := getWatchlistAction(&out, url, tc.args, tc.isRaw)

			if tc.expError != nil {
				if err == nil {
					t.Fatalf("Expected error %q, got no error.", tc.expError)
				}

				if !errors.Is(err, tc.expError) {
					t.Errorf("Expected error %q, got %q.", tc.expError, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q.", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("Expected output %q, got %q.", tc.expOut, out.String())
			}
		})
	}
}

func TestAddWatchlistAction(t *testing.T) {
	expURLPath := "/account/null/watchlist"
	expMethod := http.MethodPost
	expBody := "{\"media_type\":\"movie\",\"media_id\":799,\"watchlist\":true}\n"
	expContentType := "application/json"
	expOut := "{\n   \"success\": true,\n   \"status_code\": 1,\n   \"status_message\": \"Success.\"\n}"
	args := []string{"movie", "799", "yes"}

	url, cleanup := mockServer(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != expURLPath {
				t.Errorf("Expected path %q, got %q", expURLPath, r.URL.Path)
			}

			if r.Method != expMethod {
				t.Errorf("Expected method %q, got %q", expMethod, r.Method)
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			r.Body.Close()

			if string(body) != expBody {
				t.Errorf("Expected body: %q, got %q", expBody, string(body))
			}

			contentType := r.Header.Get("Content-Type")
			if contentType != expContentType {
				t.Errorf("Expected Content-Type %q, got %q", expContentType, contentType)
			}

			w.WriteHeader(testResp["resultsAddWatchlist"].Status)
			fmt.Fprintln(w, testResp["resultsAddWatchlist"].Body)
		})
	defer cleanup()

	var out bytes.Buffer

	if err := addWatchlistAction(&out, url, args); err != nil {
		t.Fatalf("Expected no error, got %q", err)
	}

	if expOut != out.String() {
		t.Errorf("Expected output %q, got %q", expOut, out.String())
	}
}

func TestGetRatedAction(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		isRaw       bool
		closeServer bool
	}{
		{
			name:     "RatedMovies",
			args:     []string{"movies"},
			expError: nil,
			expOut:   "1. Title: The Wild Robot\nRelease Date: 2024-09-12\nPopularity: 64.82\nVote Count: 4716\nVote Average: 8.33\n\n2. Title: A Minecraft Movie\nRelease Date: 2025-03-31\nPopularity: 695.71\nVote Count: 499\nVote Average: 6.10\n\n",
			resp:     testResp["resultsGetRated"],
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.resp.Status)
					fmt.Fprintln(w, tc.resp.Body)
				})
			defer cleanup()

			if tc.closeServer {
				cleanup()
			}

			var out bytes.Buffer

			err := getRatedAction(&out, url, tc.args, tc.isRaw)

			if tc.expError != nil {
				if err == nil {
					t.Fatalf("Expected error %q, got no error.", tc.expError)
				}

				if !errors.Is(err, tc.expError) {
					t.Errorf("Expected error %q, got %q.", tc.expError, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q.", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("Expected output %q, got %q.", tc.expOut, out.String())
			}
		})
	}
}

func TestGetRatedEpisodesAction(t *testing.T) {
	testCases := []struct {
		name     string
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		isRaw       bool
		closeServer bool
	}{
		{
			name:     "RatedEpisodes",
			expError: nil,
			expOut:   "1. Name: The Long Night\nEps Number: 3\nAir Date: 2019-04-28\nVote Count: 308\nVote Average: 6.87\n\n2. Name: The Iron Throne\nEps Number: 6\nAir Date: 2019-05-19\nVote Count: 343\nVote Average: 4.57\n\n",
			resp:     testResp["resultsGetRatedEpisodes"],
			isRaw:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.resp.Status)
					fmt.Fprintln(w, tc.resp.Body)
				})
			defer cleanup()

			if tc.closeServer {
				cleanup()
			}

			var out bytes.Buffer
			err := getRatedEpisodesAction(&out, url, tc.isRaw)

			if tc.expError != nil {
				if err == nil {
					t.Fatalf("Expected error %q, got no error.", tc.expError)
				}

				if !errors.Is(err, tc.expError) {
					t.Errorf("Expected error %q, got %q.", tc.expError, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("Expected output %q, got %q.", tc.expOut, out.String())
			}
		})
	}
}
