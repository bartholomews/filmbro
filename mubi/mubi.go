package mubi

import (
	"encoding/json"
	"fmt"
	"github.com/bartholomews/filmbro/flags"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
)

func getRequest(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Client-Country", flags.MubiUserCountry)
	req.Header.Add("Client", "web")
	cobra.CheckErr(err)
	res, err := http.DefaultClient.Do(req)
	cobra.CheckErr(err)
	return res
}

func UserLists() []List {
	var userListsResponse = getUserLists()
	return userListsResponse.Lists
}

func FilmsInList(list List) []Film {
	var films []Film
	for _, film := range list.FilmIds {
		films = append(films, getFilm(film))
	}
	return films
}

func GetAllRatingsForUser() RatingsLookup {
	ratingsLookup := RatingsLookup{}

	getRatings := func(endpoint string) UserRatings {
		fmt.Printf("Calling [%s]\n", endpoint)
		res := getRequest(endpoint)
		var userRatingsResponse UserRatings

		err := json.NewDecoder(res.Body).Decode(&userRatingsResponse)
		fmt.Println(userRatingsResponse)
		if err != nil {
			fmt.Printf("Something went wrong while getting ratings: [%s]", err)
			os.Exit(1)
		}

		for _, rating := range userRatingsResponse.Ratings {
			ratingsLookup[rating.FilmId] = rating.Stars
		}
		fmt.Printf("Got [%d] results\n", len(userRatingsResponse.Ratings))
		fmt.Printf("RatingsLookup is now [%d]\n", len(ratingsLookup))
		return userRatingsResponse
	}

	var cursorParam = ""
	for {
		var endpoint = fmt.Sprintf("%s/ratings?per_page=25%s", userEndpoint(), cursorParam)
		rating := getRatings(endpoint)
		// FIXME[FB] Find out why the second call is always empty
		//  The cursor from first call is suspiciously large (e.g. 824636555976 vs from ui is like 24836661)
		if rating.Meta.NextCursor == nil {
			break
		}
		fmt.Printf("Next cursor is [%d]\n", rating.Meta.NextCursor)
		cursorParam = fmt.Sprintf("&before=%d", rating.Meta.NextCursor)
	}

	return ratingsLookup
}

func userEndpoint() string {
	return "https://api.mubi.com/v3/users/" + strconv.Itoa(flags.MubiUserId)
}

func GetUser() User {
	res := getRequest(userEndpoint())
	var user User
	cobra.CheckErr(json.NewDecoder(res.Body).Decode(&user))
	return user
}

func getFilm(id int) Film {
	res := getRequest("https://api.mubi.com/v3/films/" + strconv.Itoa(id))
	var filmResponse Film
	cobra.CheckErr(json.NewDecoder(res.Body).Decode(&filmResponse))
	return filmResponse
}

func getUserLists() UserListsResponse {
	var endpoint = fmt.Sprintf("%s/lists?per_page=100", userEndpoint())
	res := getRequest(endpoint)
	var userListsResponse UserListsResponse
	cobra.CheckErr(json.NewDecoder(res.Body).Decode(&userListsResponse))
	return userListsResponse
}
