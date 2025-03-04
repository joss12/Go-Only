package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/go-api/dbtools"
	"github.com/go-api/restlayer"
)

type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func main() {

	file, err := os.Open("db/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	conf := new(Configuration)
	json.NewDecoder(file).Decode(conf)

	dbtools.DBInit(conf.DriverName, conf.DataSourceName)

	restlayer.RestStart("127.0.0.1:8080")
}
