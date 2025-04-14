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
	"strconv"
	"text/tabwriter"

	"example.com/dummyheaad/tmdbCLI/account"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listsCmd represents the lists command
var listsCmd = &cobra.Command{
	Use:          "lists <page>",
	Short:        "Get a users list of custom lists",
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		isRaw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}

		return listsAction(os.Stdout, apiRoot, args, isRaw)
	},
}

func listsAction(out io.Writer, apiRoot string, args []string, isRaw bool) error {
	var (
		page = 1
		err  error
	)
	if len(args) > 0 {
		page, err = strconv.Atoi(args[0])
		if err != nil {
			return err
		}
	}

	url := fmt.Sprintf("%s/account/null", apiRoot)

	resp, err := account.GetLists(url, page)
	if err != nil {
		return err
	}

	if isRaw {
		return printResp(out, resp)
	}

	return printLists(out, resp)
}

func printLists(out io.Writer, resp *account.ListsResponse) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	results := resp.Results
	fmt.Fprint(w, "Lists:\n")
	for i, r := range results {
		fmt.Fprintf(w, "%d. ", i+1)
		fmt.Fprintf(w, "Name: %s\n", r.Name)
		fmt.Fprintf(w, "Description: %s\n", r.Description)
		fmt.Fprintf(w, "List Type: %s\n", r.ListType)
		fmt.Fprintf(w, "Total Items: %d\n\n", r.ItemCount)
	}
	return w.Flush()
}

func init() {
	accountCmd.AddCommand(listsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listsCmd.PersistentFlags().String("foo", "", "A help for foo")

	listsCmd.Flags().BoolP("raw", "r", false, "Print raw json output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
