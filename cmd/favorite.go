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

// favoriteCmd represents the favorite command
var favoriteCmd = &cobra.Command{
	Use:          "favorite",
	Short:        "Manage favorite movies/tv shows",
	SilenceUsage: true,
}

var addCmd = &cobra.Command{
	Use:          "add <media_type> <media_id> <is_favourite>",
	Short:        "Mark a movie or TV show as a favourite\n<media_type>: movies or tv\n<media_id>: valid media id (integer)\n<is_favourite>: yes or no",
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

		mediaType := args[0]

		return getAction(os.Stdout, apiRoot, mediaType)
	},
}

func getAction(out io.Writer, apiRoot, mediaType string) error {

	url := fmt.Sprintf("%s/account/null", apiRoot)

	resp, err := account.GetFavorite(url, mediaType)
	if err != nil {
		return err
	}

	switch r := resp.(type) {
	case *account.FavoriteMoviesResponse:
		return printResp(out, r)
	case *account.FavoriteTvResponse:
		return printResp(out, r)
	default:
		return printResp(out, r)
	}
}

func init() {
	accountCmd.AddCommand(favoriteCmd)

	favoriteCmd.AddCommand(addCmd)
	favoriteCmd.AddCommand(getCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// favoriteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// favoriteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
