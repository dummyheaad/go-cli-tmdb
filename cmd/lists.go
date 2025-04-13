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

	"example.com/dummyheaad/tmdbCLI/account"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listsCmd represents the lists command
var listsCmd = &cobra.Command{
	Use:          "lists <page>",
	Short:        "Get a users list of custom lists",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")

		return listsAction(os.Stdout, apiRoot, args)
	},
}

func listsAction(out io.Writer, apiRoot string, args []string) error {
	var page int
	if len(args) == 0 {
		page = 1
	}
	page, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/account/null", apiRoot)

	resp, err := account.GetLists(url, page)
	if err != nil {
		return err
	}

	return printResp(out, resp)
}

func init() {
	accountCmd.AddCommand(listsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
