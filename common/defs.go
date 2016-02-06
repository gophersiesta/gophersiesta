package common

import "encoding/json"

// Placeholder groups the basic information to work with placeholders
type Placeholder struct { // ${DATASOURCE_URL:jdbc:mysql://localhost:3306/shcema?profileSQL=true}
	PropertyName  string `json:"property_name"`  // the full path to the property datasource.url
	PropertyValue string `json:"property_value"` // jdbc:mysql://localhost:3306/shcema?profileSQL=true
	PlaceHolder   string `json:"placeholder"`    // DATASOURCE_URL
}

// Placeholders is a collection of Placeholder
type Placeholders struct {
	Placeholders []*Placeholder `json:"placeholders"`
}

// Value is a pair of name and his value
type Value struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Values is a collection of Value
type Values struct {
	Values []*Value `json:"values"`
}

func (values *Values) String() string {

	params, err := values.ToMapString()
	if err != nil {
		return ""
	}

	jsonString, jsonError := json.Marshal(params)
	if jsonError != nil {
		return ""
	}

	return string(jsonString)
}

func (values *Values) ToMapString() (map[string]string, error) {
	vMap := make(map[string]string)

	for _, v := range values.Values {
		vMap[v.Name] = v.Value
	}

	return vMap, nil
}

func (pls *Placeholders) String() string {

	var params []map[string]string
	for _, p := range pls.Placeholders {
		params = append(params, p.ToMapString())
	}

	jsonString, jsonError := json.Marshal(params)
	if jsonError != nil {
		return ""
	}

	return string(jsonString)
}

func (placeholder *Placeholder) ToMapString() map[string]string {
	pMap := make(map[string]string)

	pMap["name"] = placeholder.PropertyName
	pMap["value"] = placeholder.PropertyValue
	pMap["placeholder"] = placeholder.PlaceHolder

	return pMap
}
