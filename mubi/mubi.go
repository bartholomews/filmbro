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
		if err != nil {
			fmt.Printf("Something went wrong while getting ratings: [%s]", err)
			os.Exit(1)
		}

		for _, rating := range userRatingsResponse.Ratings {
			ratingsLookup[rating.FilmId] = rating.Stars
		}
		return userRatingsResponse
	}

	var cursorParam = ""
	for {
		var endpoint = fmt.Sprintf("%s/ratings?per_page=100%s", userEndpoint(), cursorParam)
		rating := getRatings(endpoint)
		fmt.Printf("Total ratings accumulated so far: [%d]\n", len(ratingsLookup))
		if rating.Meta.NextCursor == 0 {
			break
		}
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
