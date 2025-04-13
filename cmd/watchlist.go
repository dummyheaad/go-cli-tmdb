/*
Copyright © 2025 Aysuka Ansari, LLC
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"example.com/dummyheaad/tmdbCLI/account"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// watchlistCmd represents the watchlist command
var watchlistCmd = &cobra.Command{
	Use:          "watchlist",
	Short:        "Manage watchlist movies/tv shows",
	SilenceUsage: true,
}

var addWatchlistCmd = &cobra.Command{
	Use:          "add <media_type> <media_id> <is_watchlist>",
	Short:        "Add a movie or TV show to your watchlist",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		return addWatchlistAction(os.Stdout, apiRoot, args)
	},
}

func addWatchlistAction(out io.Writer, apiRoot string, args []string) error {
	mediaType := args[0]
	if mediaType != "movie" && mediaType != "tv" {
		return errors.New("invalid <media_type> value")
	}

	mediaID, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	var isWatchlist bool
	switch args[2] {
	case "yes":
		isWatchlist = true
	case "no":
		isWatchlist = false
	default:
		return errors.New("invalid <is_watchlist> value")
	}

	url := fmt.Sprintf("%s/account/null", apiRoot)

	resp, err := account.AddWatchlist(url, mediaType, mediaID, isWatchlist)
	if err != nil {
		return err
	}

	return printResp(out, resp)
}

var getWatchlistCmd = &cobra.Command{
	Use:          "get <media_type>",
	Short:        "Get a list of movies/tv show added to a users watchlist",
	SilenceUsage: true,
	Args:         cobra.OnlyValidArgs,
	ValidArgs:    []string{"movies", "tv"},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		return getWatchlistAction(os.Stdout, apiRoot, args)
	},
}

func getWatchlistAction(out io.Writer, apiRoot string, args []string) error {

	url := fmt.Sprintf("%s/account/null", apiRoot)

	mediaType := args[0]

	resp, err := account.GetWatchlist(url, mediaType)
	if err != nil {
		return err
	}

	switch r := resp.(type) {
	case *account.WatchlistMoviesResponse:
		return printResp(out, r)
	case *account.WatchlistTvResponse:
		return printResp(out, r)
	default:
		return printResp(out, r)
	}
}

func init() {
	accountCmd.AddCommand(watchlistCmd)

	watchlistCmd.AddCommand(addWatchlistCmd)
	watchlistCmd.AddCommand(getWatchlistCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
