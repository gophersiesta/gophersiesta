package client

import (
	"fmt"

	"strings"

	"github.com/gophersiesta/gophersiesta/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/gophersiesta/gophersiesta/common"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get --appname=[appname] --source=[source_url]",
	Short: "Get the values to be setup. From appname + label",
	Long:  "Get the values to be setup. From appname + label",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GetValues())
	},
}

// GetValues return the raw response from the server calling http://url/conf/:appname/values
func GetValues() string {

	api := common.NewAPI(source)
	values, err := api.GetValues(appName, strings.Split(label, ","))

	if err != nil {
		return "{}"
	}

	return values.String()
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&appName, "appname", "a", "", "Application name")
	getCmd.Flags().StringVarP(&source, "source", "s", "https://gophersiesta.herokuapp.com/", "Source directory to read from")
	getCmd.Flags().StringVarP(&label, "label", "l", "", "Select label to be fetch")

}
