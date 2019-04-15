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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/haad/worktracker/cmd"
	"github.com/haad/worktracker/sql"
)

// initConfig reads in config file
func initConfig() {

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(home + "/.worktracker")
	viper.SetConfigFile("worktracker")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed(), "Error: ", err.Error())
	}
}

// Customer model for keeping info about given customer
func main() {

	initConfig()
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sql.DBInit("sqlite3", home+"/.worktracker/"+"worktracker.db")
	sql.DBPreload()

	cmd.Execute()

	fmt.Println("Closing DB..")
	sql.DBc.Close()
}
