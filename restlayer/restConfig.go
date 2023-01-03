package restlayer

import (
	"net/http"

	"github.com/gorilla/mux"
)

func restConfig(router *mux.Router) {

	restRouter := router.PathPrefix("/rest/api").Subrouter()

	httpRouter := router.PathPrefix("").Subrouter()

	// ------------
	// Students API
	// ------------

	// localhost:8080/rest/api/students
	restRouter.Methods("GET").Path("/students").HandlerFunc(SelectAllStudents)

	//localhost:8080/rest/api/student/{name}
	restRouter.Methods("GET").Path("/student/{name}").HandlerFunc(SelectStudentBasedName)

	//localhost:8080/rest/api/student/id/{id}
	restRouter.Methods("GET").Path("/student/id/{id}").HandlerFunc(SelectStudentBasedId)

	//localhost:8080/rest/api/student/reg/{reg}
	restRouter.Methods("GET").Path("/student/reg/{reg}").HandlerFunc(SelectStudentBasedReg)

	//localhost:8080/rest/api/student/add
	restRouter.Methods("PUT").Path("/student/add").HandlerFunc(SaveStudent)

	//localhost:8080/rest/api/student/edit
	restRouter.Methods("POST").Path("/student/edit").HandlerFunc(UpdateStudent)

	//localhost:8080/rest/api/student/deleteid/{id}
	restRouter.Methods("DELETE").Path("/student/deleteid/{id}").HandlerFunc(DeleteStudentId)

	//localhost:8080/rest/api/student/deleteall
	restRouter.Methods("DELETE").Path("/student/deleteall").HandlerFunc(DeleteAllStudents)

	//localhost:8080/rest/api/student/addmultiple
	restRouter.Methods("POST").Path("/student/addmultiple").HandlerFunc(SaveMultipleStudents)

	// ---
	// Teachers API
	// ---

	// localhost:8080/rest/api/teachers
	restRouter.Methods("GET").Path("/teachers").HandlerFunc(SelectAllTeachers)

	//localhost:8080/rest/api/teacher/add
	restRouter.Methods("PUT").Path("/teacher/add").HandlerFunc(SaveTeacher)

	//localhost:8080/rest/api/teacher/deleteid/{id}
	restRouter.Methods("DELETE").Path("/teacher/deleteid/{id}").HandlerFunc(DeleteTeacherId)

	//localhost:8080/rest/api/teacher/deleteall
	restRouter.Methods("DELETE").Path("/teacher/deleteall").HandlerFunc(DeleteAllTeachers)

	//localhost:8080/rest/api/teacher/edit
	restRouter.Methods("POST").Path("/teacher/edit").HandlerFunc(UpdateTeacher)

	//localhost:8080/rest/api/teacher/id/{id}
	restRouter.Methods("GET").Path("/teacher/id/{id}").HandlerFunc(SelectTeacherBasedId)

	//localhost:8080/rest/api/teacher/{name}
	restRouter.Methods("GET").Path("/teacher/{name}").HandlerFunc(SelectTeacherBasedName)

	//localhost:8080/rest/api/teacher/reg/{reg}
	restRouter.Methods("GET").Path("/teacher/reg/{reg}").HandlerFunc(SelectTeacherBasedReg)

	//localhost:8080/rest/api/teacher/addmultiple
	restRouter.Methods("POST").Path("/teacher/addmultiple").HandlerFunc(SaveMultipleTeachers)

	// ---
	// Courses API
	// ---

	//localhost:8080/rest/api/course/add
	restRouter.Methods("PUT").Path("/course/add").HandlerFunc(SaveCourse)

	// localhost:8080/rest/api/courses
	restRouter.Methods("GET").Path("/courses").HandlerFunc(SelectAllCourses)

	//localhost:8080/rest/api/course/addmultiple
	restRouter.Methods("POST").Path("/course/addmultiple").HandlerFunc(SaveMultipleCourses)

	//localhost:8080/rest/api/course/deleteall
	restRouter.Methods("DELETE").Path("/course/deleteall").HandlerFunc(DeleteAllCourses)

	//localhost:8080/rest/api/course/deleteid/{id}
	restRouter.Methods("DELETE").Path("/course/deleteid/{id}").HandlerFunc(DeleteCourseId)

	//localhost:8080/rest/api/course/edit
	restRouter.Methods("POST").Path("/course/edit").HandlerFunc(UpdateCourse)

	// ---
	// File Show System
	// ---

	//localhost:8080/configHtmlServer.json
	httpRouter.Methods("GET").Path("/configHtmlServer.json").HandlerFunc(ShowHtmlFile)

	// ---
	// HTML Files Show
	// ---

	//localhost:8080/index.html
	httpRouter.HandleFunc("/", indexTemplateHandling)

	// ---
	// HTML Files Show Students
	// ---

	//localhost:8080/create_student.html
	httpRouter.HandleFunc("/create_student.html", createStudentTemplateHandling)

	//localhost:8080/delete_all_student.html
	httpRouter.HandleFunc("/delete_all_student.html", deleteAllStudentsTemplateHandling)

	//localhost:8080/delete_id_student.html
	httpRouter.HandleFunc("/delete_id_student.html", deleteIdStudentTemplateHandling)

	//localhost:8080/edit_id_student.html
	httpRouter.HandleFunc("/edit_id_student.html", editIdStudentTemplateHandling)

	//localhost:8080/show_all_student.html
	httpRouter.HandleFunc("/show_all_student.html", showAllStudentsTemplateHandling)

	//localhost:8080/show_name_student.html
	httpRouter.HandleFunc("/show_name_student.html", showNameStudentTemplateHandling)

	//localhost:8080/show_id_student.html
	httpRouter.HandleFunc("/show_id_student.html", showIdStudentTemplateHandling)

	//localhost:8080/show_reg_student.html
	httpRouter.HandleFunc("/show_reg_student.html", showRegStudentTemplateHandling)

	// ---
	// HTML Files Show Teachers
	// ---

	//localhost:8080/create_teacher.html
	httpRouter.HandleFunc("/create_teacher.html", createTeacherTemplateHandling)

	//localhost:8080/delete_id_teacher.html
	httpRouter.HandleFunc("/delete_id_teacher.html", deleteIdTeacherTemplateHandling)

	//localhost:8080/show_all_teacher.html
	httpRouter.HandleFunc("/show_all_teacher.html", showAllTeachersTemplateHandling)

	//localhost:8080/delete_all_teacher.html
	httpRouter.HandleFunc("/delete_all_teacher.html", deleteAllTeachersTemplateHandling)

	//localhost:8080/edit_id_teacher.html
	httpRouter.HandleFunc("/edit_id_teacher.html", editIdTeacherTemplateHandling)

	//localhost:8080/show_id_teacher.html
	httpRouter.HandleFunc("/show_id_teacher.html", showIdTeacherTemplateHandling)

	//localhost:8080/show_name_teacher.html
	httpRouter.HandleFunc("/show_name_teacher.html", showNameTeacherTemplateHandling)

	//localhost:8080/show_reg_teacher.html
	httpRouter.HandleFunc("/show_reg_teacher.html", showRegTeacherTemplateHandling)

	// ---
	// HTML Files Show Courses
	// ---

	//localhost:8080/create_course.html
	httpRouter.HandleFunc("/create_course.html", createCourseTemplateHandling)

	//localhost:8080/show_all_course.html
	httpRouter.HandleFunc("/show_all_course.html", showAllCoursesTemplateHandling)

	//localhost:8080/edit_id_course.html
	httpRouter.HandleFunc("/edit_id_course.html", editIdCourseTemplateHandling)

	//localhost:8080/delete_all_course.html
	httpRouter.HandleFunc("/delete_all_course.html", deleteAllCoursesTemplateHandling)

	//localhost:8080/delete_id_course.html
	httpRouter.HandleFunc("/delete_id_course.html", deleteIdCourseTemplateHandling)

}

func RestStart(endpoint string) error {

	router := mux.NewRouter()

	restConfig(router)

	return http.ListenAndServe(endpoint, router)
}
