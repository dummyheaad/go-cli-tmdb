/*
Copyright © 2025 Aysuka Ansari, LLC
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"example.com/dummyheaad/tmdbCLI/account"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ratedCmd represents the rated command
var ratedCmd = &cobra.Command{
	Use:          "rated",
	Short:        "Manage rated movies/tv show",
	SilenceUsage: true,
}

var getRatedCmd = &cobra.Command{
	Use:          "get <media_type>",
	Short:        "Get a users list of rated movies/TV shows",
	SilenceUsage: true,
	Args:         cobra.OnlyValidArgs,
	ValidArgs:    []string{"movies", "tv"},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		isRaw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}

		return getRatedAction(os.Stdout, apiRoot, args, isRaw)
	},
}

func getRatedAction(out io.Writer, apiRoot string, args []string, isRaw bool) error {
	url := fmt.Sprintf("%s/account/null", apiRoot)

	mediaType := args[0]

	if mediaType == "movies" {
		resp, err := account.GetRatedShow[*account.RatedMoviesResponse](url, mediaType)
		if err != nil {
			return err
		}

		if isRaw {
			return printResp(out, resp)
		}

		return printRatedMovies(out, resp)
	}

	resp, err := account.GetRatedShow[*account.RatedTvResponse](url, mediaType)
	if err != nil {
		return err
	}

	if isRaw {
		return printResp(out, resp)
	}

	return printRatedTv(out, resp)
}

func printRatedMovies(out io.Writer, resp *account.RatedMoviesResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	results := resp.Results
	for i, r := range results {
		fmt.Fprintf(w, "%d. ", i+1)
		fmt.Fprintf(w, "Title: %s\n", r.Title)
		fmt.Fprintf(w, "Release Date: %s\n", r.ReleaseDate)
		fmt.Fprintf(w, "Popularity: %.2f\n", r.Popularity)
		fmt.Fprintf(w, "Vote Count: %d\n", r.VoteCount)
		fmt.Fprintf(w, "Vote Average: %.2f\n\n", r.VoteAverage)
	}
	return w.Flush()
}

func printRatedTv(out io.Writer, resp *account.RatedTvResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	results := resp.Results
	for i, r := range results {
		fmt.Fprintf(w, "%d. ", i+1)
		fmt.Fprintf(w, "Name: %s\n", r.Name)
		fmt.Fprintf(w, "First Air Date: %s\n", r.FirstAirDate)
		fmt.Fprintf(w, "Popularity: %.2f\n", r.Popularity)
		fmt.Fprintf(w, "Vote Count: %d\n", r.VoteCount)
		fmt.Fprintf(w, "Vote Average: %.2f\n\n", r.VoteAverage)
	}
	return w.Flush()
}

var getRatedEpisodesCmd = &cobra.Command{
	Use:          "get-eps",
	Short:        "Get a users list of rated TV episodes",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		isRaw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}

		return getRatedEpisodesAction(os.Stdout, apiRoot, isRaw)
	},
}

func getRatedEpisodesAction(out io.Writer, apiRoot string, isRaw bool) error {
	url := fmt.Sprintf("%s/account/null", apiRoot)

	resp, err := account.GetRatedEpisodes(url)
	if err != nil {
		return err
	}

	if isRaw {
		return printResp(out, resp)
	}

	return printRatedTvEps(out, resp)
}

func printRatedTvEps(out io.Writer, resp *account.RatedTvEpisodeResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	results := resp.Results
	for i, r := range results {
		fmt.Fprintf(w, "%d. ", i+1)
		fmt.Fprintf(w, "Name: %s\n", r.Name)
		fmt.Fprintf(w, "Eps Number: %d\n", r.EpisodeNumber)
		fmt.Fprintf(w, "Air Date: %s\n", r.AirDate)
		fmt.Fprintf(w, "Vote Count: %d\n", r.VoteCount)
		fmt.Fprintf(w, "Vote Average: %.2f\n\n", r.VoteAverage)
	}
	return w.Flush()
}

func init() {
	accountCmd.AddCommand(ratedCmd)

	ratedCmd.AddCommand(getRatedCmd)
	ratedCmd.AddCommand(getRatedEpisodesCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ratedCmd.PersistentFlags().String("foo", "", "A help for foo")

	getRatedCmd.Flags().BoolP("raw", "r", false, "Print raw json output")
	getRatedEpisodesCmd.Flags().BoolP("raw", "r", false, "Print raw json output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ratedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
