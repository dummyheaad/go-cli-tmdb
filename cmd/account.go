/*
Copyright © 2025 Aysuka Ansari, LLC
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"example.com/dummyheaad/tmdbCLI/account"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:          "account",
	Short:        "TMDB API for account",
	SilenceUsage: true,
}

func printDetails(out io.Writer, resp *account.DetailsResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	fmt.Fprintf(w, "Account details for %d\n", resp.ID)
	fmt.Fprintf(w, "ID: %d\n", resp.ID)
	fmt.Fprintf(w, "Username: %s\n", resp.Username)
	fmt.Fprintf(w, "Include Adult: %t\n", resp.IncludeAdult)
	fmt.Fprintf(w, "ISO_639_1: %s\n", resp.ISO_639_1)
	fmt.Fprintf(w, "ISO_3166_1: %s\n", resp.ISO_3166_1)
	fmt.Fprintf(w, "Avatar\n\tGravatar\n\t\tHash: %s\n", resp.Avatar.Gravatar.Hash)
	fmt.Fprintf(w, "Avatar\n\tTMDB\n\t\tAvatar Path: %s\n", resp.Avatar.Tmdb.AvatarPath)
	return w.Flush()
}

// TODO: implement printFavMovies and printFavTvShow
// TODO: implement printWatchlistMovies and printWatchlistTvShow

func printResp(out io.Writer, resp any) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	// Print using type conversion interface{} into struct
	s, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s", string(s))
	return w.Flush()
}

func init() {
	rootCmd.AddCommand(accountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
