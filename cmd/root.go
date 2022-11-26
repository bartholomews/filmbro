/*
Package cmd
Copyright © 2022 Federico Bartolomei filmbro@bartholomews.io
*/
package cmd

import (
	"github.com/bartholomews/filmbro/flags"
	"github.com/bartholomews/filmbro/mubi"
	"github.com/spf13/cobra"
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
		mubi.Go()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// TODO -> "mubi-to-letterboxd-lists" cmd (https://dev.to/divrhino/building-an-interactive-cli-app-with-go-cobra-promptui-346n)
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.filmbro.yaml)")

	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().IntVar(&flags.MubiUserId, "user", -1, "MUBI user id")
	listCmd.PersistentFlags().StringVar(&flags.MubiUserCountry, "country", "GB", "Country code")
	//err := listCmd.MarkFlagRequired("user")
	//if err != nil {
	//	fmt.Printf("Input error: [%s]", err)
	//	os.Exit(1)
	//}
}
