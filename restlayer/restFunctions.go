package restlayer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"rest/database/dbtools"
	"rest/database/model"
	"strconv"

	"github.com/gorilla/mux"
)

func SelectStudentBasedName(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	name, ok := vars["name"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Student not Found")
	}

	student, err := dbtools.SelectStudentBasedName(name)

	if err != nil {
		json.NewEncoder(response).Encode("Student not found!")
	} else {
		json.NewEncoder(response).Encode(student)
	}

}

func SelectStudentBasedId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Student not Found")
	}

	student, err := dbtools.SelectStudentBasedId(id)

	if err != nil {
		json.NewEncoder(response).Encode("Student not found!")
	} else {
		json.NewEncoder(response).Encode(student)
	}

}

func SelectStudentBasedReg(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	reg, ok := vars["reg"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Student not Found")
	}

	student, err := dbtools.SelectStudentBasedReg(reg)

	if err != nil {
		json.NewEncoder(response).Encode("Student not found!")
	} else {
		json.NewEncoder(response).Encode(student)
	}

}

func SelectAllStudents(response http.ResponseWriter, request *http.Request) {

	students := dbtools.SelectAllStudents()

	json.NewEncoder(response).Encode(students)
}

func SaveStudent(response http.ResponseWriter, request *http.Request) {

	var student model.Student

	err := json.NewDecoder(request.Body).Decode(&student)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not add new Student by error:%v", err)
		return
	}

	dbtools.Save(student)

}

func UpdateStudent(response http.ResponseWriter, request *http.Request) {

	var student model.Student

	err := json.NewDecoder(request.Body).Decode(&student)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update student by error:%v", err)
		return
	}

	dbtools.Update(student)

}

func DeleteStudentId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Id student not Found")
	}

	// convert string to int
	idStudent, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Cannot convert string to int")
	}

	student := dbtools.DeleteId(idStudent)

	json.NewEncoder(response).Encode(student)

}

func DeleteAll(response http.ResponseWriter, request *http.Request) {

	student := dbtools.DeleteAll()

	json.NewEncoder(response).Encode(student)

}

func SaveMultipleStudents(response http.ResponseWriter, request *http.Request) {

	var students []model.Student

	err := json.NewDecoder(request.Body).Decode(&students)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update student by error:%v", err)
		return
	}

	dbtools.SaveMultipleStudents(students)

}

// --- These functions allow to show the data in HTML format on the same domain (CORS policy avoid) ---

func ShowHtmlFile(response http.ResponseWriter, request *http.Request) {

	http.ServeFile(response, request, request.URL.Path[1:])

}

func indexTemplateHandling(response http.ResponseWriter, request *http.Request) {

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"templates/base_template.html",
		"templates/index_template.html",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func createTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/create_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func deleteAllTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/deleteall_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func deleteIdTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/deleteid_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func editTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/edit_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func showAllTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/showall_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func showStudentTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/showstudent_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func showStudentIdTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/showstudentid_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func showStudentRegTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/showstudentreg_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}
