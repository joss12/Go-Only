package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-mysql/database/dbtools"
	// "github.com/go-mysql/database/model"
)

type configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func main() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	//creating an object
	conf := new(configuration)
	json.NewDecoder(file).Decode(conf)

	dbtools.DBInitializer(conf.DriverName, conf.DataSourceName)

	// student := dbtools.SelectStudentBasedID(1)

	// fmt.Println("ID:",student.ID, student.Name,"\nAge",student.Age)

	// student := model.Student{
	// 	ID: 2,
	// 	Name: "Grace",
	// 	Age: 10,
	// }
	// lastInsertID := dbtools.Save(student)
	// fmt.Println("Last Inserted ID:", lastInsertID)

	// student := model.Student{
	// 	ID:   2,
	// 	Name: "Esther",
	// 	Age:  11,
	// }
	// rowsAffected := dbtools.Update(student)
	// fmt.Println("RowsAffected:", rowsAffected)

	rowsAffected := dbtools.Delete(1);
	fmt.Println("Rows Affected", rowsAffected)

}
