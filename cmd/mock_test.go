package cmd

import (
	"net/http"
	"net/http/httptest"
)

var testResp = map[string]struct {
	Status int
	Body   string
}{
	"resultsDetails": {
		Status: http.StatusOK,
		Body: `{
  "avatar": {
    "gravatar": {
      "hash": "5a33321a08977fbf047ab4d39105637a"
    },
    "tmdb": {
      "avatar_path": "/yUaRo4KmeADP7lkAS0t9p7r36yQ.jpg"
    }
  },
  "id": 21907685,
  "iso_639_1": "en",
  "iso_3166_1": "ID",
  "name": "Uka",
  "include_adult": false,
  "username": "clairvoyance27"
}`,
	},
	"resultsLists": {
		Status: http.StatusOK,
		Body: `{
  "page": 1,
  "results": [
    {
      "description": "test my list 2",
      "favorite_count": 0,
      "id": 8525470,
      "item_count": 2,
      "iso_639_1": "en",
      "list_type": "movie",
      "name": "my-list-2",
      "poster_path": null
    },
    {
      "description": "test my list",
      "favorite_count": 0,
      "id": 8521773,
      "item_count": 2,
      "iso_639_1": "en",
      "list_type": "movie",
      "name": "my-list",
      "poster_path": null
    }
  ],
  "total_pages": 1,
  "total_results": 2
}`,
	},
	"resultsNoLists": {
		Status: http.StatusOK,
		Body: `{
  "page": 500,
  "results": [],
  "total_pages": 1,
  "total_results": 2
}`,
	},
	"resultsFavMovies": {
		Status: http.StatusOK,
		Body: `{
  "page": 1,
  "results": [
    {
      "adult": false,
      "backdrop_path": "/m2mzlsJjE3UAqeUB5fLUkpWg4Iq.jpg",
      "genre_ids": [
        53,
        878
      ],
      "id": 1165067,
      "original_language": "en",
      "original_title": "Cosmic Chaos",
      "overview": "Battles in virtual reality, survival in a post-apocalyptic wasteland, a Soviet spaceship giving a distress signal - Fantastic stories created with advanced special effects and passion.",
      "popularity": 160.0442,
      "poster_path": "/mClzWv7gBqgXfjZXp49Enyoex1v.jpg",
      "release_date": "2023-08-03",
      "title": "Cosmic Chaos",
      "video": false,
      "vote_average": 6,
      "vote_count": 46
    },
    {
      "adult": false,
      "backdrop_path": "/qmeFQVVx7qBBW8NE1ZNVh5PSwLS.jpg",
      "genre_ids": [
        53
      ],
      "id": 555,
      "original_language": "en",
      "original_title": "Absolut",
      "overview": "Two guys against globalization want to plant a virus in the network of a finance corporation. On the day of the attack Alex has an accident and cannot remember anything. Visions and reality are thrown together in a confusing maze. Alex tries to escape from this muddle but what he discovers turns out to be rather frightening…",
      "popularity": 0.2883,
      "poster_path": "/17tI2vsEoMZFnzfkg5RCrtcG59s.jpg",
      "release_date": "2005-04-20",
      "title": "Absolut",
      "video": false,
      "vote_average": 7.8,
      "vote_count": 29
    }
  ],
  "total_pages": 1,
  "total_results": 2
}`,
	},
	"resultsFavTv": {
		Status: http.StatusOK,
		Body: `{
  "page": 1,
  "results": [
    {
      "adult": false,
      "backdrop_path": "/jeP3It0ZPY3SKW3632qwLkkIZv3.jpg",
      "genre_ids": [
        35
      ],
      "id": 550,
      "origin_country": [
        "GB"
      ],
      "original_language": "en",
      "original_name": "Till Death Us Do Part",
      "overview": "Following the chronicles of the East End working-class Garnett family, headed by patriarch Alf Garnett, a reactionary working-class man who holds racist and anti-socialist views.",
      "popularity": 12.8235,
      "poster_path": "/5r8enLaWs3SnVoInZYsOLZgboki.jpg",
      "first_air_date": "1966-06-06",
      "name": "Till Death Us Do Part",
      "vote_average": 7.4,
      "vote_count": 24
    },
    {
      "adult": false,
      "backdrop_path": "/zZqpAXxVSBtxV9qPBcscfXBcL2w.jpg",
      "genre_ids": [
        10765,
        18,
        10759
      ],
      "id": 1399,
      "origin_country": [
        "US"
      ],
      "original_language": "en",
      "original_name": "Game of Thrones",
      "overview": "Seven noble families fight for control of the mythical land of Westeros. Friction between the houses leads to full-scale war. All while a very ancient evil awakens in the farthest north. Amidst the war, a neglected military order of misfits, the Night's Watch, is all that stands between the realms of men and icy horrors beyond.",
      "popularity": 192.1505,
      "poster_path": "/1XS1oqL89opfnbLl8WnZY1O1uJx.jpg",
      "first_air_date": "2011-04-17",
      "name": "Game of Thrones",
      "vote_average": 8.456,
      "vote_count": 24838
    }
  ],
  "total_pages": 1,
  "total_results": 3
}`,
	},
	"resultsAddFav": {
		Status: http.StatusCreated,
		Body: `{
  "success": true,
  "status_code": 1,
  "status_message": "Success."
}`,
	},
}

func mockServer(h http.HandlerFunc) (string, func()) {
	ts := httptest.NewServer(h)

	return ts.URL, func() {
		ts.Close()
	}
}
