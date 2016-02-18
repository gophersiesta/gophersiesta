// Copyright © 2016 GOPHERSIESTA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"os"

	"strings"

	"github.com/gophersiesta/gophersiesta/common"
)

func main() {

	host := os.Getenv("SIESTA_HOST")
	appname := os.Getenv("SIESTA_APP")
	labels := strings.Split(os.Getenv("SIESTA_LABELS"), ",")

	if appname == "" {
		fmt.Println("ERROR, SIESTA_APP not defined")
		os.Exit(1)
	}

	if host == "" {
		host = "http://gophersiesta.herokuapp.com/"
	}

	api := common.NewAPI(host)

	values, err := api.GetValues(appname, labels)
	if err != nil {
		fmt.Println("ERROR GETTING THE VALUES")
		os.Exit(2)
	}

	for _, v := range values.Values {
		os.Setenv(v.Name, v.Value)
	}

	listEnv()

}

func listEnv() {
	// list out the
	var env []string
	env = os.Environ()

	fmt.Println("List of Environtment variables : \n")

	for index, value := range env {
		name := strings.Split(value, "=") // split by = sign

		fmt.Printf("[%d] %s : %v\n", index, name[0], name[1])
	}

}
