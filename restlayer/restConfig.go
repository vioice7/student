package restlayer

import (
	"net/http"

	"github.com/gorilla/mux"
)

func restConfig(router *mux.Router) {

	restRouter := router.PathPrefix("/rest/api").Subrouter()

	httpRouter := router.PathPrefix("").Subrouter()

	// localhost:8080/rest/api/students
	restRouter.Methods("GET").Path("/students").HandlerFunc(SelectAllStudents)

	//localhost:8080/rest/api/student/{name}
	restRouter.Methods("GET").Path("/student/{name}").HandlerFunc(SelectStudentBasedName)

	//localhost:8080/rest/api/student/id/{id}
	restRouter.Methods("GET").Path("/student/id/{id}").HandlerFunc(SelectStudentBasedId)

	//localhost:8080/rest/api/student/reg/{reg}
	restRouter.Methods("GET").Path("/student/reg/{reg}").HandlerFunc(SelectStudentBasedReg)

	//localhost:8080/rest/api/student/add
	restRouter.Methods("POST").Path("/student/add").HandlerFunc(SaveStudent)

	//localhost:8080/rest/api/student/edit
	restRouter.Methods("POST").Path("/student/edit").HandlerFunc(UpdateStudent)

	//localhost:8080/rest/api/student/deleteid/{id}
	restRouter.Methods("DELETE").Path("/student/deleteid/{id}").HandlerFunc(DeleteStudentId)

	//localhost:8080/rest/api/student/deleteall
	restRouter.Methods("DELETE").Path("/student/deleteall").HandlerFunc(DeleteAll)

	//localhost:8080/rest/api/student/addmultiple
	restRouter.Methods("POST").Path("/student/addmultiple").HandlerFunc(SaveMultipleStudents)

	//localhost:8080/index.html
	httpRouter.Methods("GET").Path("/").HandlerFunc(ShowHtmlFile)

	//localhost:8080/showall.html
	httpRouter.Methods("GET").Path("/showall.html").HandlerFunc(ShowHtmlFile)

	//localhost:8080/create.html
	httpRouter.Methods("GET").Path("/create.html").HandlerFunc(ShowHtmlFile)

	//localhost:8080/deleteall.html
	httpRouter.Methods("GET").Path("/deleteall.html").HandlerFunc(ShowHtmlFile)

	//localhost:8080/deletebyid.html
	httpRouter.Methods("GET").Path("/deleteid.html").HandlerFunc(ShowHtmlFile)

	//localhost:8080/edit.html
	httpRouter.Methods("GET").Path("/edit.html").HandlerFunc(ShowHtmlFile)

	//localhost:8080/showstudent.html
	httpRouter.Methods("GET").Path("/showstudent.html").HandlerFunc(ShowHtmlFile)

	//localhost:8080/showstudentid.html
	httpRouter.Methods("GET").Path("/showstudentid.html").HandlerFunc(ShowHtmlFile)

	//localhost:8080/showstudentreg.html
	httpRouter.Methods("GET").Path("/showstudentreg.html").HandlerFunc(ShowHtmlFile)

	//localhost:8080/configHtmlServer.json
	httpRouter.Methods("GET").Path("/configHtmlServer.json").HandlerFunc(ShowHtmlFile)
}

func RestStart(endpoint string) error {

	router := mux.NewRouter()

	restConfig(router)

	return http.ListenAndServe(endpoint, router)
}
