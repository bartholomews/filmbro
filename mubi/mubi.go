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
	res := getRequest(lists())
	var userListsResponse UserLists

	cobra.CheckErr(json.NewDecoder(res.Body).Decode(&userListsResponse))

	var userLists = userListsResponse.Lists
	fmt.Printf("I found %d Mubi lists:\n", len(userLists))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for _, list := range userLists {
		fmt.Printf("[%d] %s\n", list.Id+1, list.Title)
	}
}

var baseUri = "https://api.mubi.com/v3"

func lists() string {
	var req = baseUri + "/users/" + strconv.Itoa(flags.MubiUserId) + "/lists"
	fmt.Printf("GET request to Mubi: [%s]\n", req)
	return req
}
