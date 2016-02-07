package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"os"

	"github.com/gophersiesta/gophersiesta/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/gophersiesta/gophersiesta/common"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set all the values to configuration manager. Needed {appname + label}",
	Long:  "Set all the values to be setup. From appname + label",
	Run: func(cmd *cobra.Command, args []string) {
		api := common.NewAPI(source)

		data := map[string]interface{}{}
		json.Unmarshal([]byte(properties), &data)

		var values common.Values

		if len(data) > 0 { //it's a JSON
			for k, v := range data {
				values.Values = append(values.Values, &common.Value{k, fmt.Sprint(v)})
			}
		} else if strings.Contains(properties, "=") {
			pairs := strings.Split(properties, ",")
			for _, v := range pairs {
				vv := strings.Split(v, "=")
				values.Values = append(values.Values, &common.Value{vv[0], strings.Join(vv[1:], "=")})
			}
		} else {
			fmt.Println("ERROR")
			os.Exit(0)
		}

		success, err := api.SetValues(appName, strings.Split(label, ","), values)

		if err != nil {
			fmt.Println("ERROR")
		}

		fmt.Println(success)

	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&appName, "appname", "a", "", "Application name")
	setCmd.Flags().StringVarP(&source, "source", "s", "https://gophersiesta.herokuapp.com/", "Source directory to read from")
	setCmd.Flags().StringVarP(&label, "label", "l", "", "Select label to be fetch")
	setCmd.Flags().StringVarP(&properties, "properties", "p", "", "json encoded properties")

}
