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
	"resultsWatchlistMovies": {
		Status: http.StatusOK,
		Body: `{
  "page": 1,
  "results": [
    {
      "adult": false,
      "backdrop_path": "/2w4xG178RpB4MDAIfTkqAuSJzec.jpg",
      "genre_ids": [
        12,
        28,
        878
      ],
      "id": 11,
      "original_language": "en",
      "original_title": "Star Wars",
      "overview": "Princess Leia is captured and held hostage by the evil Imperial forces in their effort to take over the galactic Empire. Venturesome Luke Skywalker and dashing captain Han Solo team together with the loveable robot duo R2-D2 and C-3PO to rescue the beautiful princess and restore peace and justice in the Empire.",
      "popularity": 20.1525,
      "poster_path": "/6FfCtAuVAW8XJjZ7eWeLibRLWTw.jpg",
      "release_date": "1977-05-25",
      "title": "Star Wars",
      "video": false,
      "vote_average": 8.203,
      "vote_count": 21056
    },
    {
      "adult": false,
      "backdrop_path": "/hZkgoQYus5vegHoetLkCJzb17zJ.jpg",
      "genre_ids": [
        18
      ],
      "id": 550,
      "original_language": "en",
      "original_title": "Fight Club",
      "overview": "A ticking-time-bomb insomniac and a slippery soap salesman channel primal male aggression into a shocking new form of therapy. Their concept catches on, with underground \"fight clubs\" forming in every town, until an eccentric gets in the way and ignites an out-of-control spiral toward oblivion.",
      "popularity": 33.7606,
      "poster_path": "/pB8BM7pdSp6B6Ih7QZ4DrQ3PmJK.jpg",
      "release_date": "1999-10-15",
      "title": "Fight Club",
      "video": false,
      "vote_average": 8.438,
      "vote_count": 30142
    },
    {
      "adult": false,
      "backdrop_path": "/jqFjgNnxpXIXWuPsyfqmcLXRo9p.jpg",
      "genre_ids": [
        80,
        53
      ],
      "id": 500,
      "original_language": "en",
      "original_title": "Reservoir Dogs",
      "overview": "A botched robbery indicates a police informant, and the pressure mounts in the aftermath at a warehouse. Crime begets violence as the survivors -- veteran Mr. White, newcomer Mr. Orange, psychopathic parolee Mr. Blonde, bickering weasel Mr. Pink and Nice Guy Eddie -- unravel.",
      "popularity": 10.4402,
      "poster_path": "/xi8Iu6qyTfyZVDVy60raIOYJJmk.jpg",
      "release_date": "1992-09-02",
      "title": "Reservoir Dogs",
      "video": false,
      "vote_average": 8.122,
      "vote_count": 14565
    },
    {
      "adult": false,
      "backdrop_path": "/oyatchDPpS4I9jpIIezFJGrmXcR.jpg",
      "genre_ids": [
        18
      ],
      "id": 499,
      "original_language": "fr",
      "original_title": "Cléo de 5 à 7",
      "overview": "Agnès Varda eloquently captures Paris in the sixties with this real-time portrait of a singer set adrift in the city as she awaits test results of a biopsy. A chronicle of the minutes of one woman’s life, Cléo from 5 to 7 is a spirited mix of vivid vérité and melodrama, featuring a score by Michel Legrand and cameos by Jean-Luc Godard and Anna Karina.",
      "popularity": 2.4708,
      "poster_path": "/oelBStY4xpguaplRv15P3Za7Xsr.jpg",
      "release_date": "1962-04-11",
      "title": "Cléo from 5 to 7",
      "video": false,
      "vote_average": 7.7,
      "vote_count": 706
    }
  ],
  "total_pages": 1,
  "total_results": 4
}`,
	},
	"resultsWatchlistTv": {
		Status: http.StatusOK,
		Body: `{
  "page": 1,
  "results": [
    {
      "adult": false,
      "backdrop_path": "/oRdc2nn7jLOYy4fBdvmFKPsKzZE.jpg",
      "genre_ids": [
        80,
        18,
        9648
      ],
      "id": 2734,
      "origin_country": [
        "US"
      ],
      "original_language": "en",
      "original_name": "Law & Order: Special Victims Unit",
      "overview": "In the criminal justice system, sexually-based offenses are considered especially heinous. In New York City, the dedicated detectives who investigate these vicious felonies are members of an elite squad known as the Special Victims Unit. These are their stories.",
      "popularity": 341.293,
      "poster_path": "/abWOCrIo7bbAORxcQyOFNJdnnmR.jpg",
      "first_air_date": "1999-09-20",
      "name": "Law & Order: Special Victims Unit",
      "vote_average": 7.938,
      "vote_count": 3905
    },
    {
      "adult": false,
      "backdrop_path": "/dWl6Nez0OI4Sr2hImLSqfiODmwW.jpg",
      "genre_ids": [
        18,
        10768
      ],
      "id": 450,
      "origin_country": [
        "GB"
      ],
      "original_language": "en",
      "original_name": "Island at War",
      "overview": "Island at War is a British television series that tells the story of the German Occupation of the Channel Islands. It primarily focuses on three local families: the upper class Dorrs, the middle class Mahys and the working class Jonases, and four German officers. The fictional island of St. Gregory serves as a stand-in for the real-life islands Jersey and Guernsey, and the story is compiled from the events on both islands.\n\nProduced by Granada Television in Manchester, Island at War had an estimated budget of £9,000,000 and was filmed on location in the Isle of Man from August 2003 to October 2003. When the series was shown in the UK, it appeared in six 70-minute episodes.",
      "popularity": 1.6837,
      "poster_path": "/g47UV12d7sPUxkSF1ARrsYDJhta.jpg",
      "first_air_date": "2004-07-11",
      "name": "Island at War",
      "vote_average": 7.4,
      "vote_count": 9
    }
  ],
  "total_pages": 1,
  "total_results": 2
}`,
	},
	"resultsAddWatchlist": {
		Status: http.StatusCreated,
		Body: `{
  "success": true,
  "status_code": 1,
  "status_message": "Success."
}`,
	},
	"resultsGetRated": {
		Status: http.StatusOK,
		Body: `{
  "page": 1,
  "results": [
    {
      "adult": false,
      "backdrop_path": "/mQZJoIhTEkNhCYAqcHrQqhENLdu.jpg",
      "genre_ids": [
        16,
        878,
        10751
      ],
      "id": 1184918,
      "original_language": "en",
      "original_title": "The Wild Robot",
      "overview": "After a shipwreck, an intelligent robot called Roz is stranded on an uninhabited island. To survive the harsh environment, Roz bonds with the island's animals and cares for an orphaned baby goose.",
      "popularity": 64.8238,
      "poster_path": "/wTnV3PCVW5O92JMrFvvrRcV39RU.jpg",
      "release_date": "2024-09-12",
      "title": "The Wild Robot",
      "video": false,
      "vote_average": 8.327,
      "vote_count": 4716,
      "rating": 8
    },
    {
      "adult": false,
      "backdrop_path": "/vz14urfVNjvibJEWDAaJScxZDcZ.jpg",
      "genre_ids": [
        10751,
        35,
        12,
        14
      ],
      "id": 950387,
      "original_language": "en",
      "original_title": "A Minecraft Movie",
      "overview": "Four misfits find themselves struggling with ordinary problems when they are suddenly pulled through a mysterious portal into the Overworld: a bizarre, cubic wonderland that thrives on imagination. To get back home, they'll have to master this world while embarking on a magical quest with an unexpected, expert crafter, Steve.",
      "popularity": 695.7057,
      "poster_path": "/yFHHfHcUgGAxziP1C3lLt0q2T4s.jpg",
      "release_date": "2025-03-31",
      "title": "A Minecraft Movie",
      "video": false,
      "vote_average": 6.1,
      "vote_count": 499,
      "rating": 8
    }
  ],
  "total_pages": 1,
  "total_results": 2
}`,
	},
	"resultsGetRatedEpisodes": {
		Status: http.StatusOK,
		Body: `{
  "page": 1,
  "results": [
    {
      "air_date": "2019-04-28",
      "episode_number": 3,
      "episode_type": "standard",
      "id": 1551827,
      "name": "The Long Night",
      "overview": "The Night King and his army have arrived at Winterfell and the great battle begins. Arya looks to prove her worth as a fighter.",
      "production_code": "803",
      "runtime": 82,
      "season_number": 8,
      "show_id": 1399,
      "still_path": "/mFtHbZenI5rRPqC5OFafoVmjEjq.jpg",
      "vote_average": 6.873,
      "vote_count": 308,
      "rating": 4
    },
    {
      "air_date": "2019-05-19",
      "episode_number": 6,
      "episode_type": "finale",
      "id": 1551830,
      "name": "The Iron Throne",
      "overview": "In the aftermath of the devastating attack on King's Landing, Daenerys must face the survivors.",
      "production_code": "806",
      "runtime": 80,
      "season_number": 8,
      "show_id": 1399,
      "still_path": "/zBi2O5EJfgTS6Ae0HdAYLm9o2nf.jpg",
      "vote_average": 4.574,
      "vote_count": 343,
      "rating": 2
    }
  ],
  "total_pages": 1,
  "total_results": 2
}`,
	},
}

func mockServer(h http.HandlerFunc) (string, func()) {
	ts := httptest.NewServer(h)

	return ts.URL, func() {
		ts.Close()
	}
}
