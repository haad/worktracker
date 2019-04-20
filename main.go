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
	"io"
	"log"
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

func backupDB(sourceDB string, backupDB string) {
	from, err := os.Open(sourceDB)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(backupDB, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

}

// Customer model for keeping info about given customer
func main() {

	// initConfig()
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbfile := home + "/.worktracker/" + "worktracker.db"
	backupfile := dbfile + ".backup"

	// Backup db on first run to make sure we have a copy
	backupDB(dbfile, backupfile)

	sql.DBInit("sqlite3", dbfile)
	sql.DBPreload()

	cmd.Execute()

	fmt.Println("Closing DB..")
	sql.DBc.Close()
}
