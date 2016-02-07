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

	"github.com/gophersiesta/gophersiesta/common"
)

func main() {

	//appname := "app1"
	//labels := []string{""}

	api := common.NewAPI("http://localhost:4747")
	api.Debug(true)

	fmt.Println("GET APPS ===")
	fmt.Println(api.GetApps())

	/*
		fmt.Println("GET TEMPLATE ===")
		fmt.Println(api.GetTemplate(appname))

		fmt.Println("GET PLACEHOLDERS ===")
		fmt.Println(api.GetPlaceholders(appname))

		fmt.Println("GET LABELS ===")
		fmt.Println(api.GetLabels(appname))

		fmt.Println("GET VALUES ===")
		fmt.Println(api.GetValues(appname, labels))

		fmt.Println("SET VALUES ===")
		v1 := common.Value{"name1", "value1"}
		v2 := common.Value{"name2", "value2"}
		v3 := common.Value{"name3", "value3"}
		values := common.Values{[]*common.Value{&v1, &v2, &v3}}
		fmt.Println(api.SetValues(appname, labels, values))

		fmt.Println("GET VALUES (AGAIN) ===")
		fmt.Println(api.GetValues(appname, labels))

		fmt.Println("RENDER ===")
		fmt.Println(api.Render(appname, labels, "json"))
	*/
}
