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

// ---
// Rest Layer API Students
// ---

func SelectStudentBasedName(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	name, ok := vars["name"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Student not found.")
	}

	student, err := dbtools.SelectStudentBasedName(name)

	if err != nil {
		json.NewEncoder(response).Encode("Student not found.")
	} else {
		json.NewEncoder(response).Encode(student)
	}

}

func SelectStudentBasedId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Student not found.")
	}

	student, err := dbtools.SelectStudentBasedId(id)

	if err != nil {
		json.NewEncoder(response).Encode("Student not found.")
	} else {
		json.NewEncoder(response).Encode(student)
	}
}

func SelectStudentBasedReg(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	reg, ok := vars["reg"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Student not found.")
	}

	student, err := dbtools.SelectStudentBasedReg(reg)

	if err != nil {
		json.NewEncoder(response).Encode("Student not found.")
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
		fmt.Fprintf(response, "Could not add new Student by error: %v.", err)
		return
	}

	dbtools.SaveStudent(student)

}

func UpdateStudent(response http.ResponseWriter, request *http.Request) {

	var student model.Student

	err := json.NewDecoder(request.Body).Decode(&student)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update student by error: %v.", err)
		return
	}

	dbtools.UpdateStudent(student)

}

func DeleteStudentId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Id student not found.")
	}

	// convert string to int
	idStudent, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	student := dbtools.DeleteStudentId(idStudent)

	json.NewEncoder(response).Encode(student)

}

func DeleteAllStudents(response http.ResponseWriter, request *http.Request) {

	student := dbtools.DeleteAllStudents()

	json.NewEncoder(response).Encode(student)

}

func SaveMultipleStudents(response http.ResponseWriter, request *http.Request) {

	var students []model.Student

	err := json.NewDecoder(request.Body).Decode(&students)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update student by error: %v.", err)
		return
	}

	dbtools.SaveMultipleStudents(students)

}

// ---
// Rest Layer API Teachers
// ---

func SelectAllTeachers(response http.ResponseWriter, request *http.Request) {

	teachers := dbtools.SelectAllTeachers()

	json.NewEncoder(response).Encode(teachers)

}

func SaveTeacher(response http.ResponseWriter, request *http.Request) {

	var teacher model.Teacher

	err := json.NewDecoder(request.Body).Decode(&teacher)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not add new Teacher by error: %v.", err)
		return
	}

	dbtools.SaveTeacher(teacher)

}

func DeleteTeacherId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Id teacher not found.")
	}

	// convert string to int
	idTeacher, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	teacher := dbtools.DeleteTeacherId(idTeacher)

	json.NewEncoder(response).Encode(teacher)

}

func DeleteAllTeachers(response http.ResponseWriter, request *http.Request) {

	teacher := dbtools.DeleteAllTeachers()

	json.NewEncoder(response).Encode(teacher)

}

func UpdateTeacher(response http.ResponseWriter, request *http.Request) {

	var teacher model.Teacher

	err := json.NewDecoder(request.Body).Decode(&teacher)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update teacher by error: %v.", err)
		return
	}

	dbtools.UpdateTeacher(teacher)

}

func SelectTeacherBasedId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Teacher not found.")
	}

	teacher, err := dbtools.SelectTeacherBasedId(id)

	if err != nil {
		json.NewEncoder(response).Encode("Teacher not found.")
	} else {
		json.NewEncoder(response).Encode(teacher)
	}
}

func SelectTeacherBasedName(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	name, ok := vars["name"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Teacher not found.")
	}

	teacher, err := dbtools.SelectTeacherBasedName(name)

	if err != nil {
		json.NewEncoder(response).Encode("Teacher not found.")
	} else {
		json.NewEncoder(response).Encode(teacher)
	}

}

func SelectTeacherBasedReg(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	reg, ok := vars["reg"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Teacher not found.")
	}

	teacher, err := dbtools.SelectTeacherBasedReg(reg)

	if err != nil {
		json.NewEncoder(response).Encode("Teacher not found.")
	} else {
		json.NewEncoder(response).Encode(teacher)
	}

}

func SaveMultipleTeachers(response http.ResponseWriter, request *http.Request) {

	var teachers []model.Teacher

	err := json.NewDecoder(request.Body).Decode(&teachers)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update teacher by error: %v.", err)
		return
	}

	dbtools.SaveMultipleTeachers(teachers)

}

// ---
// Rest Layer API Courses
// ---

func SaveCourse(response http.ResponseWriter, request *http.Request) {

	var course model.Course

	err := json.NewDecoder(request.Body).Decode(&course)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not add new Course by error: %v.", err)
		return
	}

	dbtools.SaveCourse(course)

}

func SelectAllCourses(response http.ResponseWriter, request *http.Request) {

	courses := dbtools.SelectAllCourses()

	json.NewEncoder(response).Encode(courses)

}

func SaveMultipleCourses(response http.ResponseWriter, request *http.Request) {

	var courses []model.Course

	err := json.NewDecoder(request.Body).Decode(&courses)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update course by error: %v.", err)
		return
	}

	dbtools.SaveMultipleCourses(courses)

}

func UpdateCourse(response http.ResponseWriter, request *http.Request) {

	var course model.Course

	err := json.NewDecoder(request.Body).Decode(&course)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update course by error: %v.", err)
		return
	}

	dbtools.UpdateCourse(course)

}

func DeleteAllCourses(response http.ResponseWriter, request *http.Request) {

	course := dbtools.DeleteAllCourses()

	json.NewEncoder(response).Encode(course)

}

func DeleteCourseId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Id course not found.")
	}

	// convert string to int
	idCourse, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	course := dbtools.DeleteCourseId(idCourse)

	json.NewEncoder(response).Encode(course)

}

// Teacher Course Link Functions

func TeacherCourseLink(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idTeacher, ok := vars["teacherid"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "ID teacher not found.")
	}

	idCourse, ok := vars["courseid"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "ID course not found.")
	}

	idCoursei, err := strconv.Atoi(idCourse)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	idTeacheri, err := strconv.Atoi(idTeacher)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	dbtools.ConnectionCoursesTeachers(idTeacheri, idCoursei)

}

func CheckTeacherCourseLink(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idTeacher, ok := vars["teacherid"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "ID teacher not found.")
	}

	idCourse, ok := vars["courseid"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "ID course not found.")
	}

	idCoursei, err := strconv.Atoi(idCourse)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	idTeacheri, err := strconv.Atoi(idTeacher)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	checkConnection := dbtools.CheckConnectionCoursesTeachers(idTeacheri, idCoursei)

	if checkConnection {
		json.NewEncoder(response).Encode("connected")
	} else {
		json.NewEncoder(response).Encode("notconnected")
	}

}

// Student Course Link Functions

func StudentCourseLink(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStudent, ok := vars["studentid"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "ID student not found.")
	}

	idCourse, ok := vars["courseid"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "ID course not found.")
	}

	idCoursei, err := strconv.Atoi(idCourse)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	idStudenti, err := strconv.Atoi(idStudent)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	dbtools.ConnectionCoursesStudents(idStudenti, idCoursei)

}

func CheckStudentCourseLink(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStudent, ok := vars["studentid"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "ID student not found.")
	}

	idCourse, ok := vars["courseid"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "ID course not found.")
	}

	idCoursei, err := strconv.Atoi(idCourse)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	idStudenti, err := strconv.Atoi(idStudent)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	checkConnection := dbtools.CheckConnectionCoursesStudents(idStudenti, idCoursei)

	if checkConnection {
		json.NewEncoder(response).Encode("connected")
	} else {
		json.NewEncoder(response).Encode("notconnected")
	}

}

// ---
// HTML Files Template Handling
// ---

// --- These functions allow to show the data in HTML format on the same domain (CORS policy avoid) ---

func ShowHtmlFile(response http.ResponseWriter, request *http.Request) {

	http.ServeFile(response, request, request.URL.Path[1:])

}

// ---
// HTML Template Handling
// ---

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

// ---
// Student HTML Template Handling
// ---

func createStudentTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/student/create_student_template.html",
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

func deleteAllStudentsTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/student/delete_all_student_template.html",
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

func deleteIdStudentTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/student/delete_id_student_template.html",
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

func editIdStudentTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/student/edit_student_template.html",
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

func showAllStudentsTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/student/show_all_student_template.html",
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

func showNameStudentTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/student/show_name_student_template.html",
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

func showIdStudentTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/student/show_id_student_template.html",
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

func showRegStudentTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/student/show_reg_student_template.html",
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

// ---
// Teacher HTML Template Handling
// ---

func showAllTeachersTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/teacher/show_all_teacher_template.html",
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

func createTeacherTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/teacher/create_teacher_template.html",
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

func deleteIdTeacherTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/teacher/delete_id_teacher_template.html",
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

func deleteAllTeachersTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/teacher/delete_all_teacher_template.html",
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

func editIdTeacherTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/teacher/edit_teacher_template.html",
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

func showIdTeacherTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/teacher/show_id_teacher_template.html",
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

func showNameTeacherTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/teacher/show_name_teacher_template.html",
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

func showRegTeacherTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/teacher/show_reg_teacher_template.html",
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

// ---
// Course HTML Template Handling
// ---

func createCourseTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/course/create_course_template.html",
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

func showAllCoursesTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/course/show_all_course_template.html",
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

func editIdCourseTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/course/edit_course_template.html",
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

func deleteAllCoursesTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/course/delete_all_course_template.html",
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

func deleteIdCourseTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/course/delete_id_course_template.html",
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

func linkCourseTeacherTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/course/link_course_teacher_template.html",
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

func linkCourseStudentTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/course/link_course_student_template.html",
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
