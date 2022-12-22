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

	//localhost:8080/configHtmlServer.json
	httpRouter.Methods("GET").Path("/configHtmlServer.json").HandlerFunc(ShowHtmlFile)

	//localhost:8080/index.html
	httpRouter.HandleFunc("/", indexTemplateHandling)

	//localhost:8080/create.html
	httpRouter.HandleFunc("/create.html", createTemplateHandling)

	//localhost:8080/deleteall.html
	httpRouter.HandleFunc("/deleteall.html", deleteAllTemplateHandling)

	//localhost:8080/deleteid.html
	httpRouter.HandleFunc("/deleteid.html", deleteIdTemplateHandling)

	//localhost:8080/edit.html
	httpRouter.HandleFunc("/edit.html", editTemplateHandling)

	//localhost:8080/showall.html
	httpRouter.HandleFunc("/showall.html", showAllTemplateHandling)

	//localhost:8080/showstudent.html
	httpRouter.HandleFunc("/showstudent.html", showStudentTemplateHandling)

	//localhost:8080/showstudentid.html
	httpRouter.HandleFunc("/showstudentid.html", showStudentIdTemplateHandling)

	//localhost:8080/showstudentreg.html
	httpRouter.HandleFunc("/showstudentreg.html", showStudentRegTemplateHandling)
}

func RestStart(endpoint string) error {

	router := mux.NewRouter()

	restConfig(router)

	return http.ListenAndServe(endpoint, router)
}
