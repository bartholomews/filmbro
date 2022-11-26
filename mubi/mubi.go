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

func Go() {
	var user = getUser()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Hello %s\n", user.Name)
	var userLists = getUserLists()
	fmt.Printf("I found %d Mubi lists:\n", len(userLists))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for _, list := range userLists {
		fmt.Printf("[%d] %s\n", list.Id+1, list.Title)
	}
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

func getUserLists() []List {
	var endpoint = fmt.Sprintf("%s/lists", userEndpoint())
	res := getRequest(endpoint)
	var userListsResponse UserLists
	cobra.CheckErr(json.NewDecoder(res.Body).Decode(&userListsResponse))
	return userListsResponse.Lists
}
