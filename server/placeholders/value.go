package placeholders

import (
	"github.com/gophersiesta/gophersiesta/common"
)


// CreateValues transforms a map of string to Values struct
func CreateValues(m map[string]string) common.Values {
	values := make([]*common.Value, len(m))
	i := 0
	for k, v := range m {
		value := &common.Value{}
		value.Name = k
		value.Value = v
		values[i] = value
		i++
	}

	return common.Values{values}
}
