package restlayer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-api/dbtools"
	"github.com/go-api/model"
	"github.com/gorilla/mux"
)

func SelectStudentBasedName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	name, ok := vars["name"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Student not found")
	}
	student := dbtools.SelectStudentBasedName(name)
	json.NewEncoder(w).Encode(student)
}

func SelectAllStudents(w http.ResponseWriter, r *http.Request) {
	students := dbtools.SelectAllStudent()

	json.NewEncoder(w).Encode(students)
}

func SaveStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not add new student by error:%v", err)
		return
	}

	dbtools.Save(student)

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not update student by error:%v", err)
		return
	}

	dbtools.Update(student)
}
