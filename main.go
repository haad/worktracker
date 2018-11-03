// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	//	"fmt"
	//	"log"
	//	"os"

	"github.com/haad/worktracker/cmd"
	"github.com/haad/worktracker/sql"
)

// Customer model for keeping info about given customer
func main() {

	sql.DBInit("sqlite3", "test.db")
	sql.DBPreload()

	cmd.Execute()

	fmt.Println("Closing DB..")
	sql.DBc.Close()
}
