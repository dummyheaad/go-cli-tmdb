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

		return getRatedAction(os.Stdout, apiRoot, args)
	},
}

func getRatedAction(out io.Writer, apiRoot string, args []string) error {
	url := fmt.Sprintf("%s/account/null", apiRoot)

	mediaType := args[0]

	resp, err := account.GetRatedShow(url, mediaType)
	if err != nil {
		return err
	}

	switch r := resp.(type) {
	case *account.RatedMoviesResponse:
		return printResp(out, r)
	case *account.RatedTvResponse:
		return printResp(out, r)
	default:
		return printResp(out, r)
	}
}

var getRatedEpisodesCmd = &cobra.Command{
	Use:          "get-eps",
	Short:        "Get a users list of rated TV episodes",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		return getRatedEpisodesAction(os.Stdout, apiRoot)
	},
}

func getRatedEpisodesAction(out io.Writer, apiRoot string) error {
	url := fmt.Sprintf("%s/account/null", apiRoot)

	resp, err := account.GetRatedEpisodes(url)
	if err != nil {
		return err
	}

	return printResp(out, resp)
}

func init() {
	accountCmd.AddCommand(ratedCmd)

	ratedCmd.AddCommand(getRatedCmd)
	ratedCmd.AddCommand(getRatedEpisodesCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ratedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ratedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
