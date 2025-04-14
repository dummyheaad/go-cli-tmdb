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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"example.com/dummyheaad/tmdbCLI/account"
)

// detailsCmd represents the details command
var detailsCmd = &cobra.Command{
	Use:          "details",
	Short:        "Get the public details of an account on TMDB",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		accountID, err := cmd.Flags().GetString("account-id")
		if err != nil {
			return err
		}

		isRaw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}

		return detailsAction(os.Stdout, apiRoot, accountID, isRaw)
	},
}

func detailsAction(out io.Writer, apiRoot, accountID string, isRaw bool) error {
	url := fmt.Sprintf("%s/account/%s", apiRoot, accountID)

	resp, err := account.GetDetails(url)
	if err != nil {
		return err
	}

	if isRaw {
		return printResp(out, resp)
	}

	return printDetails(out, resp)
}

func printDetails(out io.Writer, resp *account.DetailsResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	fmt.Fprintf(w, "Account details for %d\n", resp.ID)
	fmt.Fprintf(w, "ID: %d\n", resp.ID)
	fmt.Fprintf(w, "Username: %s\n", resp.Username)
	return w.Flush()
}

func init() {
	accountCmd.AddCommand(detailsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detailsCmd.PersistentFlags().String("foo", "", "A help for foo")

	detailsCmd.Flags().String("account-id", "null", "Specify the account id")
	detailsCmd.Flags().BoolP("raw", "r", false, "Print raw json output")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
