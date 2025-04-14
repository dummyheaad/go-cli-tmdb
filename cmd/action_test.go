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

}
