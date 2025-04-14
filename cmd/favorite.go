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

// favoriteCmd represents the favorite command
var favoriteCmd = &cobra.Command{
	Use:          "favorite",
	Short:        "Manage favorite movies/tv shows",
	SilenceUsage: true,
}

var addCmd = &cobra.Command{
	Use:          "add <media_type> <media_id> <is_favourite>",
	Short:        "Mark a movie or TV show as a favourite\n\n<media_type>: movies or tv\n<media_id>: valid media id (integer)\n<is_favourite>: yes or no\n",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")
		return addAction(os.Stdout, apiRoot, args)
	},
}

func addAction(out io.Writer, apiRoot string, args []string) error {
	mediaType := args[0]
	if mediaType != "movie" && mediaType != "tv" {
		return errors.New("invalid <media_type> value")
	}

	mediaID, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	var isFavourite bool
	switch args[2] {
	case "yes":
		isFavourite = true
	case "no":
		isFavourite = false
	default:
		return errors.New("invalid <is_favourite> value")
	}

	url := fmt.Sprintf("%s/account/null", apiRoot)

	resp, err := account.AddFavorite(url, mediaType, mediaID, isFavourite)
	if err != nil {
		return err
	}

	return printResp(out, resp)
}

var getCmd = &cobra.Command{
	Use:          "get <media_type>",
	Short:        "Get a users list of favourite movies/tv shows\n<media_type>: movies or tv",
	SilenceUsage: true,
	Args:         cobra.OnlyValidArgs,
	ValidArgs:    []string{"movies", "tv"},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		isRaw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}

		return getAction(os.Stdout, apiRoot, args, isRaw)
	},
}

func getAction(out io.Writer, apiRoot string, args []string, isRaw bool) error {

	url := fmt.Sprintf("%s/account/null", apiRoot)

	mediaType := args[0]

	if mediaType == "movies" {
		resp, err := account.GetFavorite[*account.FavoriteMoviesResponse](url, mediaType)
		if err != nil {
			return err
		}

		if isRaw {
			return printResp(out, resp)
		}

		return printFavMovies(out, resp)
	}
	resp, err := account.GetFavorite[*account.FavoriteTvResponse](url, mediaType)
	if err != nil {
		return err
	}

	if isRaw {
		return printResp(out, resp)
	}

	return printFavTv(out, resp)
}

func printFavMovies(out io.Writer, resp *account.FavoriteMoviesResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	results := resp.Results
	fmt.Fprint(w, "Favorite Movies:\n")
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

func printFavTv(out io.Writer, resp *account.FavoriteTvResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	results := resp.Results
	fmt.Fprint(w, "Favorite TV Shows:\n")
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
	accountCmd.AddCommand(favoriteCmd)

	favoriteCmd.AddCommand(addCmd)
	favoriteCmd.AddCommand(getCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// favoriteCmd.PersistentFlags().String("foo", "", "A help for foo")

	getCmd.Flags().BoolP("raw", "r", false, "Print raw json output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// favoriteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
