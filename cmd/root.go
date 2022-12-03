/*
Package cmd
Copyright © 2022 Federico Bartolomei filmbro@bartholomews.io
*/
package cmd

import (
	"fmt"
	"github.com/bartholomews/filmbro/flags"
	"github.com/bartholomews/filmbro/letterboxd"
	"github.com/bartholomews/filmbro/mubi"
	"github.com/spf13/cobra"
	"regexp"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	// https://github.com/spf13/cobra-cli
	Use:   "filmbro",
	Short: "Mubi / Letterboxd utils",
	Long: `
    ___ _ _        _
   / __|_) |      | |
 _| |__ _| | ____ | |__   ____ ___
(_   __) | ||    \|  _ \ / ___) _ \
  | |  | | || | | | |_) ) |  | |_| |
  |_|  |_|\_)_|_|_|____/|_|   \___/
-------------------------------------------------------------------
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("Available commands: <list>")
	//},
}

var mubiDiaryToLetterboxdCmd = &cobra.Command{
	Use:   "mubi-diary-to-letterboxd",
	Short: "Create a Letterboxd csv importer file with Diary entries from Mubi lists with titles matching 'yyyy/mm'",
	Run: func(cmd *cobra.Command, args []string) {
		flags.MubiUserId = promptInt(promptContent{
			"Invalid entry: you can find it in the browser console after you login.",
			"Please provide your Mubi user id",
		})

		user := mubi.GetUser()
		lists := mubi.UserLists()

		var diaryLists []mubi.List
		for _, maybeDiaryList := range lists {
			isDiaryList, _ := regexp.MatchString(`[1-9]\d{3}/\d{2}`, maybeDiaryList.Title)
			if isDiaryList {
				diaryLists = append(diaryLists, maybeDiaryList)
			}
		}

		var diaryEntries []mubi.DiaryEntry

		fmt.Printf("Found %d Mubi lists for user [%s]\n", len(lists), user.Name)
		fmt.Printf("Found %d 'Diary lists'\n", len(diaryLists))

		if len(diaryLists) > 0 {

			fmt.Println("Retrieving all user ratings...")
			ratingsLookup := mubi.GetAllRatingsForUser()
			fmt.Printf("Retrieved %d user ratings\n", len(ratingsLookup))

			for _, diaryList := range diaryLists {
				watchedDate := diaryList.Title[0:4] + "-" + diaryList.Title[5:7] + "-01"
				filmsForList := mubi.FilmsInList(diaryList)
				for _, film := range filmsForList {
					var rating *int
					var maybeRating, hasRating = ratingsLookup[film.Id]
					if hasRating {
						rating = &maybeRating
					}
					diaryEntries = append(diaryEntries, mubi.DiaryEntry{
						Film: film, WatchedDate: watchedDate, Rating: rating,
					})
				}
			}

			fmt.Println(diaryEntries)
			letterboxd.CreateCsvImport(diaryEntries)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.filmbro.yaml)")

	rootCmd.AddCommand(mubiDiaryToLetterboxdCmd)
	mubiDiaryToLetterboxdCmd.PersistentFlags().StringVar(&flags.MubiUserCountry, "country", "GB", "Country code")
}
