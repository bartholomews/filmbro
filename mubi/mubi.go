package mubi

import (
	"encoding/json"
	"fmt"
	"github.com/bartholomews/filmbro/flags"
	"github.com/spf13/cobra"
	"net/http"
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

func Lists() []List {
	var user = getUser()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	var userListsResponse = getUserLists()
	fmt.Printf("I found %d Mubi lists for %s:\n", userListsResponse.Meta.TotalCount, user.Name)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//for i, list := range userListsResponse.Lists {
	//	fmt.Printf("%02d. %s\n", i + 1, list.Title)
	//}
	return userListsResponse.Lists
}

func FilmsInList(list List) []Film {
	var films []Film
	for _, film := range list.FilmIds {
		films = append(films, getFilm(film))
	}
	return films
}

func userEndpoint() string {
	return "https://api.mubi.com/v3/users/" + strconv.Itoa(flags.MubiUserId)
}

func getUser() User {
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

func getUserLists() UserLists {
	var endpoint = fmt.Sprintf("%s/lists?per_page=100", userEndpoint())
	res := getRequest(endpoint)
	var userListsResponse UserLists
	cobra.CheckErr(json.NewDecoder(res.Body).Decode(&userListsResponse))
	return userListsResponse
}
