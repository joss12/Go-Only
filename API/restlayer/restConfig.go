package restlayer

import (
	"net/http"

	"github.com/gorilla/mux"
)

func restConfig(router *mux.Router) {
	restRouter := router.PathPrefix("/rest/api").Subrouter()


	// localhost:8080/rest/api/students
	restRouter.Methods("GET").Path("/students").HandlerFunc(SelectAllStudents)

	// localhost:8080/rest/api/student/{name}
	restRouter.Methods("GET").Path("/student/{name}").HandlerFunc(SelectStudentBasedName)

	// localhost:8080/rest/api/student/add
	restRouter.Methods("POST").Path("/student/add").HandlerFunc(SaveStudent)

	// localhost:8080/rest/api/student/edit
	restRouter.Methods("POST").Path("/student/edit").HandlerFunc(UpdateStudent)

}

func RestStart(endpoint string)error{

	router:=mux.NewRouter()
	restConfig(router)

	return http.ListenAndServe(endpoint,router)
}
