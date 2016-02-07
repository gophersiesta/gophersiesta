package client

import (
	"fmt"

	"strings"

	"github.com/gophersiesta/gophersiesta/common"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get --appname=[appname] --source=[source_url]",
	Short: "Get the values to be setup. From appname + label",
	Long:  "Get the values to be setup. From appname + label",
	Run: func(cmd *cobra.Command, args []string) {
		api := common.NewAPI(source)
		values, err := api.GetValues(appName, strings.Split(label, ","))

		if err != nil {
			fmt.Println("{}")
		}

		fmt.Println(values.String())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&appName, "appname", "a", "", "Application name")
	getCmd.Flags().StringVarP(&source, "source", "s", "https://gophersiesta.herokuapp.com/", "Source directory to read from")
	getCmd.Flags().StringVarP(&label, "label", "l", "", "Select label to be fetch")
}
