/*
Package cmd
Copyright © 2022 Federico Bartolomei filmbro@bartholomews.io
*/
package cmd

import (
	"fmt"
	"github.com/bartholomews/filmbro/flags"
	"github.com/bartholomews/filmbro/mubi"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
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

var listCmd = &cobra.Command{
	Use:   "lists",
	Short: "Show mubi lists",
	Run: func(cmd *cobra.Command, args []string) {
		flags.MubiUserId = promptInt(promptContent{
			"Invalid entry: you can find it in the browser console after you login.",
			"Please provide your Mubi user id",
		})
		lists := mubi.Lists()
		var listStr []string
		for _, e := range lists {
			listStr = append(listStr, e.Title)
		}

		pr := promptui.Select{
			Label: "Please select a list of the form 'yyyy/mm' in order to create a Letterboxd importer for Diary entries:",
			Items: listStr,
		}

		index, _, err := pr.Run()
		cobra.CheckErr(err)

		selectedList := lists[index]
		matchRegex, _ := regexp.MatchString(`[1-9]\d{3}/\d{2}`, selectedList.Title)
		if !matchRegex {
			fmt.Printf("Expected a list with title matching 'yyyy/mm', got: [%s]\n", selectedList.Title)
			os.Exit(1)
		}

		//filmsForList := mubi.FilmsInList(selectedList)
		//for _, film := range filmsForList {
		//	fmt.Println(film.Title)
		//}
		//
		//watchedDate := selectedList.Title[0:4] + "-" + selectedList.Title[5:7] + "-01"
		//letterboxd.CreateCsvImport(filmsForList, watchedDate)
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

	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringVar(&flags.MubiUserCountry, "country", "GB", "Country code")
}
