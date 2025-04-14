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
	"text/tabwriter"

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

		isRaw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}

		return getWatchlistAction(os.Stdout, apiRoot, args, isRaw)
	},
}

func getWatchlistAction(out io.Writer, apiRoot string, args []string, isRaw bool) error {

	url := fmt.Sprintf("%s/account/null", apiRoot)

	mediaType := args[0]

	if mediaType == "movies" {
		resp, err := account.GetWatchlist[*account.WatchlistMoviesResponse](url, mediaType)
		if err != nil {
			return err
		}

		if isRaw {
			return printResp(out, resp)
		}

		return printWatchlistMovies(out, resp)
	}

	resp, err := account.GetWatchlist[*account.WatchlistTvResponse](url, mediaType)
	if err != nil {
		return err
	}

	if isRaw {
		return printResp(out, resp)
	}

	return printWatchlistTv(out, resp)
}

func printWatchlistMovies(out io.Writer, resp *account.WatchlistMoviesResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	results := resp.Results
	fmt.Fprint(w, "Watchlist Movies:\n")
	for i, r := range results {
		fmt.Fprintf(w, "%d. ", i+1)
		fmt.Fprintf(w, "Title: %s\n", r.Title)
		fmt.Fprintf(w, "Release Date: %s\n", r.ReleaseDate)
		fmt.Fprintf(w, "Popularity: %f\n", r.Popularity)
		fmt.Fprintf(w, "Vote Count: %d\n", r.VoteCount)
		fmt.Fprintf(w, "Vote Average: %f\n\n", r.VoteAverage)
	}
	return w.Flush()
}

func printWatchlistTv(out io.Writer, resp *account.WatchlistTvResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	results := resp.Results
	fmt.Fprint(w, "Watchlist TV Shows:\n")
	for i, r := range results {
		fmt.Fprintf(w, "%d. ", i+1)
		fmt.Fprintf(w, "Name: %s\n", r.Name)
		fmt.Fprintf(w, "First Air Date: %s\n", r.FirstAirDate)
		fmt.Fprintf(w, "Popularity: %f\n", r.Popularity)
		fmt.Fprintf(w, "Vote Count: %d\n", r.VoteCount)
		fmt.Fprintf(w, "Vote Average: %f\n\n", r.VoteAverage)
	}
	return w.Flush()
}

func init() {
	accountCmd.AddCommand(watchlistCmd)

	watchlistCmd.AddCommand(addWatchlistCmd)
	watchlistCmd.AddCommand(getWatchlistCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	getWatchlistCmd.Flags().BoolP("raw", "r", false, "Print raw json output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
