package common

import (
	"crypto/tls"
	"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"fmt"
	"net/http/httputil"
)

type API struct {
	endPoint    string
	version     string
	client   	*http.Client
	debug 		bool
}

func NewAPI(endPoint string) *API {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	if len(endPoint)>0 && endPoint[len(endPoint)-1:] != "/" {
		endPoint += "/"
	}

	// for debug purposes only
	if endPoint=="" {
		endPoint = "https://gophersiesta.herokuapp.com/"
	}

	return &API{endPoint, "v1", client, false}
}

func (this *API) call(action, httpMethod string, params map[string]string) ([]byte, error) {

	var err error
	var res *http.Response

	// version is not used for the moment
	url := this.endPoint+/*this.version+"/"+*/action

	if httpMethod=="POST" {

		jsonString, jsonError := json.Marshal(params)
		if jsonError!=nil {
			return nil, jsonError
		}

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonString)))
		res, err = this.client.Do(req)
		if this.debug {
			fmt.Println("[POST] =>", url, string(jsonString))
			fmt.Println("[POST] <=")
			dump, derr := httputil.DumpResponse(res, true)
			fmt.Println(string(dump), derr)
		}

	} else {
		valuesStr := ""
		for key, val := range params {
			valuesStr += "&"+key+"="+val
		}
		res, err = this.client.Get(url+"?"+valuesStr)
		if this.debug {
			fmt.Println("[GET] =>", url, valuesStr)
			fmt.Println("[GET] <=")
			dump, derr := httputil.DumpResponse(res, true)
			fmt.Println(string(dump), derr)
		}
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return body, err
}

func (this *API) Debug(value bool) {
	this.debug = value
}

func (this *API) SetVersion(version string) {
	this.version = version
}

func (this *API) SetEndPoint(endPoint string) {
	this.endPoint = endPoint
}

// http://gophersiesta.herokuapp.com/conf/:appname
func (this *API) GetTemplate(appname string) (string, error) {
	dataStream, err := this.call("conf/"+appname, "GET", nil)
	var data string
	json.Unmarshal(dataStream, &data)
	return data, err
}

// http://gophersiesta.herokuapp.com/conf/:appname/placeholders
func (this *API) GetPlaceholders(appname string) (Placeholders, error) {
	dataStream, err := this.call("conf/" + appname + "/placeholders", "GET", nil)
	var data Placeholders
	json.Unmarshal(dataStream, &data)
	return data, err
}

// http://gophersiesta.herokuapp.com/conf/:appname/labels
func (this *API) GetLabels(appname string) ([]string, error) {
	dataStream, err := this.call("conf/" + appname + "/labels", "GET", nil)
	var data []string
	json.Unmarshal(dataStream, &data)
	return data, err
}

// http://gophersiesta.herokuapp.com/conf/:appname/values?labels=:label1,:label2
func (this *API) GetValues(appname string, labels []string) (Values, error) {
	var params = map[string]string{
		"labels": strings.Join(labels, ","),
	}

	dataStream, err := this.call("conf/" + appname + "/values", "GET", params)
	var data Values
	json.Unmarshal(dataStream, &data)
	return data, err
}

// http://gophersiesta.herokuapp.com/conf/:appname/values?labels=:label1,:label2
func (this *API) SetValues(appname string, labels []string, values Values) (string, error) {
	urlValues := "labels=" + strings.Join(labels, ",")

	params, _ := values.toMapString()

	dataStream, err := this.call("conf/" + appname + "/values?" + urlValues, "POST", params)
	var data string
	json.Unmarshal(dataStream, &data)
	return data, err
}

// http://gophersiesta.herokuapp.com/conf/:appname/values?labels=:label1,:label2
func (this *API) Render(appname string, labels []string, format string) (string, error) {
	var params = map[string]string{
		"labels": strings.Join(labels, ","),
	}

	dataStream, err := this.call("conf/" + appname + "/render/" + format, "GET", params)
	var data string
	json.Unmarshal(dataStream, &data)
	return data, err
}


