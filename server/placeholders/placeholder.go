package placeholders

import (
	"fmt"
	"strings"

	"github.com/gophersiesta/gophersiesta/common"
	"github.com/spf13/viper"
)

// GetPlaceHolders uses the provided viper configuration to extract properties that have placeholders in is values
func GetPlaceHolders(conf *viper.Viper) common.Placeholders {
	list := parseMap(conf.AllSettings())

	properties := CreateProperties(list)

	return common.Placeholders{properties}
}

// CreateProperties transform the propsMap into a Property struct slice
func CreateProperties(propsMap map[string]string) []*common.Placeholder {
	count := len(propsMap)

	ps := make([]*common.Placeholder, count)
	i := 0
	for k, v := range propsMap {
		p, d, err := extractPlaceholder(v)
		if err == nil {
			p := &common.Placeholder{k, d, p}
			ps[i] = p
		}

		i++
	}

	return ps
}

func extractPlaceholder(s string) (string, string, error) {
	if s[:2] != "${" {
		return "", "", fmt.Errorf("%s does not contain any placeholder with format ${PLACEHOLER_VARIABLE[:defaultvalue]}", s)
	}

	if s[len(s)-1:len(s)] != "}" {
		return "", "", fmt.Errorf("%s does not contain any placeholder with format ${PLACEHOLER_VARIABLE[:defaultvalue]}", s)
	}

	s = s[2:]
	s = s[0 : len(s)-1]

	defaultValue := ""
	if strings.Contains(s, ":") {
		dv := strings.Split(s, ":")
		defaultValue = strings.Join(dv[1:], ":")
	}

	return strings.Split(s, ":")[0], defaultValue, nil
}

func parseMap(props map[string]interface{}) map[string]string {
	list := make(map[string]string)
	for key, value := range props {
		switch v := value.(type) {
		case map[interface{}]interface{}:
			l := parseMapInterface(v, key, list)
			for pkey, pvalue := range l {
				list[pkey] = pvalue
			}
		case string:
			if v[:2] == "${" {
				list[key] = v
			}
		default:
		}
	}
	return list
}

func parseMapInterface(props map[interface{}]interface{}, key string, list map[string]string) map[string]string {
	for k, value := range props {
		actKey := key + "." + fmt.Sprint(k)

		switch v := value.(type) {
		case map[interface{}]interface{}:
			list = parseMapInterface(v, actKey, list)
		case string:
			if v[:2] == "${" {
				keystr := fmt.Sprint(actKey) // <-- HACK
				list[keystr] = v
			}
		default:
		}
	}
	return list
}
