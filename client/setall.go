package client

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/gophersiesta/gophersiesta/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/gophersiesta/gophersiesta/common"
)

// setCmd represents the set command
var setAllCmd = &cobra.Command{
	Use:   "setall",
	Short: "Set the values to configuration manager. Needed {appname + label}",
	Long:  "Set the values to be setup. From appname + label",
	Run: func(cmd *cobra.Command, args []string) {
		api := common.NewAPI(source)
		pls, err := api.GetPlaceholders(appName)

		if err != nil {
			log.Fatal("Could not get Placehodlers")
		}

		values, err := api.GetValues(appName, strings.Split(label, ","))
		if err != nil {
			log.Fatal("Could not get Values")
		}

		fmt.Println(values.String())
		fmt.Println(pls.String())

		fmt.Printf("\nThere are %v common. Listing: \n", len(pls.Placeholders))

		for _, p := range pls.Placeholders {
			fmt.Printf("%s [$%s] (current value: \"%s\")\n", p.PropertyName, p.PlaceHolder, p.PropertyValue)
		}

		fmt.Printf("\n\n")
		fmt.Printf("Type value for each placeholder and press ENTER, or ENTER to skip or left as before: \n")
		fmt.Printf("	explanation: property.path [$PLACE_HOLDER] --> curentvalue \n")

		mValues, err := values.ToMapString()
		sValues := common.Values{}

		for _, p := range pls.Placeholders {

			pv := p.PropertyValue
			if _, ok := mValues[p.PropertyName]; ok {
				pv = mValues[p.PropertyName]
			}

			value := setPropertyHolderValue(p, pv)
			sValue := common.Value{p.PropertyName, value}
			sValues.Values = append(sValues.Values, &sValue)
		}

		api.SetValues(appName, strings.Split(label, ","), sValues)
	},
}

func setPropertyHolderValue(p *common.Placeholder, currentVal string) string {
	var value string
	fmt.Printf("%s [$%s] --> %s:", p.PropertyName, p.PlaceHolder, currentVal)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		value = scanner.Text()
		break
	}

	if value != "" {
		return value
	}
	return currentVal
}

func init() {
	rootCmd.AddCommand(setAllCmd)

	setAllCmd.Flags().StringVarP(&appName, "appname", "a", "", "Application name")
	setAllCmd.Flags().StringVarP(&source, "source", "s", "https://gophersiesta.herokuapp.com/", "Source directory to read from")
	setAllCmd.Flags().StringVarP(&label, "label", "l", "", "Select label to be fetch")

}
