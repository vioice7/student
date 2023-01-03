package dbtools

import (
	"database/sql"
	"fmt"
	"log"
	"rest/database/model"

	_ "github.com/go-sql-driver/mysql"
)

//
// Conection and initialase mysql
//

var driverName string
var dataSourceName string

func DBInitilize(dn, dsn string) {

	driverName = dn
	dataSourceName = dsn
}

func connect() (db *sql.DB) {

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

// -------------------------
// Students table connection
// -------------------------

func SelectAllStudents() []model.Student {

	db := connect()

	rows, err := db.Query("select * from students")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	students := []model.Student{}

	for rows.Next() {

		student := model.Student{}

		err = rows.Scan(&student.ID, &student.Name, &student.Age, &student.Reg)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		students = append(students, student)
	}

	return students
}

func SelectStudentBasedName(name string) (model.Student, error) {

	db := connect()

	rows := db.QueryRow("select * from students where name = ?", name)

	defer db.Close()

	student := model.Student{}

	err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Reg)

	return student, err

}

func SelectStudentBasedId(id string) (model.Student, error) {

	db := connect()

	rows := db.QueryRow("select * from students where id = ?", id)

	defer db.Close()

	student := model.Student{}

	err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Reg)

	return student, err

}

func SelectStudentBasedReg(reg string) (model.Student, error) {

	db := connect()

	defer db.Close()

	rows := db.QueryRow("select * from students where reg = ?", reg)

	student := model.Student{}

	err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Reg)

	return student, err

}

func SaveStudent(student model.Student) int64 {

	db := connect()

	defer db.Close()

	rows := db.QueryRow("select reg from students where reg = ?", student.Reg)

	errDupl := rows.Scan(&student.Reg)

	// if there is no record don't add a record and return

	if student.Reg == "" {
		return 0
	}

	// if there is a duplicate entry don't add a record and return

	if errDupl == nil {
		fmt.Println("Duplicate entry in database! No records added!")
		return 0
	}

	save, err := db.Prepare("insert into students(id,name,age,reg) values(?,?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := save.Exec(student.ID, student.Name, student.Age, student.Reg)

	if err != nil {
		log.Fatal(err.Error())
	}

	studentID, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}

	return studentID
}

func UpdateStudent(student model.Student) int64 {

	db := connect()

	defer db.Close()

	// check if there are records to compare

	rowsCheck := db.QueryRow("select id from students where id = ?", student.ID)

	studentCheck := model.Student{}

	errCheck := rowsCheck.Scan(&studentCheck.ID)

	if errCheck != nil {
		fmt.Println("Error, no records with that id available!")
		return 0
	}

	// check for unique reg

	rowsDupl := db.QueryRow("select reg from students where reg = ?", student.Reg)

	errDupl := rowsDupl.Scan(&student.Reg)

	// if there is a duplicate entry don't add a record and return

	if errDupl == nil {
		fmt.Println("Duplicate entry in database! No records added!")
		return 0
	}

	// create a new model to compare

	studentData := model.Student{}

	rows := db.QueryRow("select * from students where id = ?", student.ID)

	err := rows.Scan(&studentData.ID, &studentData.Name, &studentData.Age, &studentData.Reg)

	if err != nil {
		log.Fatal(err.Error())
	}

	// compare the data between models in order to remove blank data

	if student.Name == "" {
		student.Name = studentData.Name
	}

	if student.Age <= 0 {
		student.Age = studentData.Age
	}

	if student.Reg == "" {
		student.Reg = studentData.Reg
	}

	update, err := db.Prepare("update students set name=?, age=?, reg=? where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := update.Exec(student.Name, student.Age, student.Reg, student.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsEffected

}

func DeleteStudentId(id int) int64 {

	db := connect()

	defer db.Close()

	delete, err := db.Prepare("delete from students where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := delete.Exec(id)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsAffected
}

func DeleteAllStudents() int64 {

	db := connect()

	defer db.Close()

	// count records before truncate table

	var count int64

	err := db.QueryRow("select count(*) from students").Scan(&count)

	if err != nil {
		log.Fatal(err.Error())
	}

	// truncate table to delete all records

	deleteAll, err := db.Prepare("truncate table students")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = deleteAll.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	return count
}

func SaveMultipleStudents(students []model.Student) int64 {

	db := connect()

	defer db.Close()

	var nrAddedStudents int64

	for _, student := range students {

		// check for unique reg

		rowsDupl := db.QueryRow("select reg from students where reg = ?", student.Reg)

		errDupl := rowsDupl.Scan(&student.Reg)

		// if there is a duplicate entry don't add a record and return

		if errDupl == nil {
			fmt.Println("Duplicate entry in database! No records added!")
			continue
		}

		save, err := db.Prepare("insert into students(id,name,age,reg) values(?,?,?,?)")

		if err != nil {
			log.Fatal(err.Error())
		}

		result, err := save.Exec(student.ID, student.Name, student.Age, student.Reg)

		if err != nil {
			log.Fatal(err.Error())
		}

		studentID, err := result.LastInsertId()

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Student ID ", studentID, " inserted!")

	}

	return nrAddedStudents

}

// -------------------------
// Teachers table connection
// -------------------------

func SelectAllTeachers() []model.Teacher {

	db := connect()

	rows, err := db.Query("select * from teachers")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	teachers := []model.Teacher{}

	for rows.Next() {

		teacher := model.Teacher{}

		err = rows.Scan(&teacher.ID, &teacher.Name, &teacher.Age, &teacher.Reg)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		teachers = append(teachers, teacher)
	}

	return teachers
}

func SaveTeacher(teacher model.Teacher) int64 {

	db := connect()

	defer db.Close()

	rows := db.QueryRow("select reg from teachers where reg = ?", teacher.Reg)

	errDupl := rows.Scan(&teacher.Reg)

	// if there is no record don't add a record and return

	if teacher.Reg == "" {
		return 0
	}

	// if there is a duplicate entry don't add a record and return

	if errDupl == nil {
		fmt.Println("Duplicate entry in database! No records added!")
		return 0
	}

	save, err := db.Prepare("insert into teachers(id,name,age,reg) values(?,?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := save.Exec(teacher.ID, teacher.Name, teacher.Age, teacher.Reg)

	if err != nil {
		log.Fatal(err.Error())
	}

	teacherID, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}

	return teacherID
}

func DeleteTeacherId(id int) int64 {

	db := connect()

	defer db.Close()

	delete, err := db.Prepare("delete from teachers where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := delete.Exec(id)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsAffected
}

func DeleteAllTeachers() int64 {

	db := connect()

	defer db.Close()

	// count records before truncate table

	var count int64

	err := db.QueryRow("select count(*) from teachers").Scan(&count)

	if err != nil {
		log.Fatal(err.Error())
	}

	// truncate table to delete all records

	deleteAll, err := db.Prepare("truncate table teachers")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = deleteAll.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	return count
}

func UpdateTeacher(teacher model.Teacher) int64 {

	db := connect()

	defer db.Close()

	// check if there are records to compare

	rowsCheck := db.QueryRow("select id from teachers where id = ?", teacher.ID)

	teacherCheck := model.Teacher{}

	errCheck := rowsCheck.Scan(&teacherCheck.ID)

	if errCheck != nil {
		fmt.Println("Error, no records with that id available!")
		return 0
	}

	// check for unique reg

	rowsDupl := db.QueryRow("select reg from teachers where reg = ?", teacher.Reg)

	errDupl := rowsDupl.Scan(&teacher.Reg)

	// if there is a duplicate entry don't add a record and return

	if errDupl == nil {
		fmt.Println("Duplicate entry in database! No records added!")
		return 0
	}

	// create a new model to compare

	teacherData := model.Teacher{}

	rows := db.QueryRow("select * from teachers where id = ?", teacher.ID)

	err := rows.Scan(&teacherData.ID, &teacherData.Name, &teacherData.Age, &teacherData.Reg)

	if err != nil {
		log.Fatal(err.Error())
	}

	// compare the data between models in order to remove blank data

	if teacher.Name == "" {
		teacher.Name = teacherData.Name
	}

	if teacher.Age <= 0 {
		teacher.Age = teacherData.Age
	}

	if teacher.Reg == "" {
		teacher.Reg = teacherData.Reg
	}

	update, err := db.Prepare("update teachers set name=?, age=?, reg=? where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := update.Exec(teacher.Name, teacher.Age, teacher.Reg, teacher.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsEffected
}

func SelectTeacherBasedId(id string) (model.Teacher, error) {

	db := connect()

	rows := db.QueryRow("select * from teachers where id = ?", id)

	defer db.Close()

	teacher := model.Teacher{}

	err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Age, &teacher.Reg)

	return teacher, err

}

func SelectTeacherBasedName(name string) (model.Teacher, error) {

	db := connect()

	rows := db.QueryRow("select * from teachers where name = ?", name)

	defer db.Close()

	teacher := model.Teacher{}

	err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Age, &teacher.Reg)

	return teacher, err

}

func SelectTeacherBasedReg(reg string) (model.Teacher, error) {

	db := connect()

	defer db.Close()

	rows := db.QueryRow("select * from teachers where reg = ?", reg)

	teacher := model.Teacher{}

	err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Age, &teacher.Reg)

	return teacher, err

}

func SaveMultipleTeachers(teachers []model.Teacher) int64 {

	db := connect()

	defer db.Close()

	var nrAddedTeachers int64

	for _, teacher := range teachers {

		// check for unique reg

		rowsDupl := db.QueryRow("select reg from teachers where reg = ?", teacher.Reg)

		errDupl := rowsDupl.Scan(&teacher.Reg)

		// if there is a duplicate entry don't add a record and return

		if errDupl == nil {
			fmt.Println("Duplicate entry in database! No records added!")
			continue
		}

		save, err := db.Prepare("insert into teachers(id,name,age,reg) values(?,?,?,?)")

		if err != nil {
			log.Fatal(err.Error())
		}

		result, err := save.Exec(teacher.ID, teacher.Name, teacher.Age, teacher.Reg)

		if err != nil {
			log.Fatal(err.Error())
		}

		teacherID, err := result.LastInsertId()

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Teacher ID ", teacherID, " inserted!")

	}

	return nrAddedTeachers

}

// -----------------------
// Course table connection
// -----------------------

func SaveCourse(course model.Course) int64 {

	db := connect()

	defer db.Close()

	rows := db.QueryRow("select reg from course where reg = ?", course.Reg)

	errDupl := rows.Scan(&course.Reg)

	// if there is no record don't add a record and return

	if course.Reg == "" {
		return 0
	}

	// if there is a duplicate entry don't add a record and return

	if errDupl == nil {
		fmt.Println("Duplicate entry in database! No records added!")
		return 0
	}

	save, err := db.Prepare("insert into courses(id,name,description,teacherreg,reg) values(?,?,?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := save.Exec(course.ID, course.Name, course.Description, course.TeacherReg, course.Reg)

	if err != nil {
		log.Fatal(err.Error())
	}

	courseID, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}

	return courseID
}

func SelectAllCourses() []model.Course {

	db := connect()

	rows, err := db.Query("select * from courses")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	courses := []model.Course{}

	for rows.Next() {

		course := model.Course{}

		err = rows.Scan(&course.ID, &course.Name, &course.Description, &course.TeacherReg, &course.Reg)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		courses = append(courses, course)
	}

	return courses

}

func SaveMultipleCourses(courses []model.Course) int64 {

	db := connect()

	defer db.Close()

	var nrAddedCourses int64

	for _, course := range courses {

		// check for unique reg

		rowsDupl := db.QueryRow("select reg from courses where reg = ?", course.Reg)

		errDupl := rowsDupl.Scan(&course.Reg)

		// if there is a duplicate entry don't add a record and return

		if errDupl == nil {
			fmt.Println("Duplicate entry in database! No records added!")
			continue
		}

		save, err := db.Prepare("insert into courses(id,name,description,teacherreg,reg) values(?,?,?,?,?)")

		if err != nil {
			log.Fatal(err.Error())
		}

		result, err := save.Exec(course.ID, course.Name, course.Description, course.TeacherReg, course.Reg)

		if err != nil {
			log.Fatal(err.Error())
		}

		courseID, err := result.LastInsertId()

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Course ID ", courseID, " inserted!")

	}

	return nrAddedCourses

}

func DeleteAllCourses() int64 {

	db := connect()

	defer db.Close()

	// count records before truncate table

	var count int64

	err := db.QueryRow("select count(*) from courses").Scan(&count)

	if err != nil {
		log.Fatal(err.Error())
	}

	// truncate table to delete all records

	deleteAll, err := db.Prepare("truncate table courses")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = deleteAll.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	return count
}

func DeleteCourseId(id int) int64 {

	db := connect()

	defer db.Close()

	delete, err := db.Prepare("delete from courses where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := delete.Exec(id)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsAffected
}

func UpdateCourse(course model.Course) int64 {

	db := connect()

	defer db.Close()

	// check if there are records to compare

	rowsCheck := db.QueryRow("select id from courses where id = ?", course.ID)

	courseCheck := model.Course{}

	errCheck := rowsCheck.Scan(&courseCheck.ID)

	if errCheck != nil {
		fmt.Println("Error, no records with that id available!")
		return 0
	}

	// check for unique reg

	rowsDupl := db.QueryRow("select reg from courses where reg = ?", course.Reg)

	errDupl := rowsDupl.Scan(&course.Reg)

	// if there is a duplicate entry don't add a record and return

	if errDupl == nil {
		fmt.Println("Duplicate entry in database! No records added!")
		return 0
	}

	// create a new model to compare

	courseData := model.Course{}

	rows := db.QueryRow("select * from courses where id = ?", course.ID)

	err := rows.Scan(&courseData.ID, &courseData.Name, &courseData.Description, &courseData.TeacherReg, &courseData.Reg)

	if err != nil {
		log.Fatal(err.Error())
	}

	// compare the data between models in order to remove blank data

	if course.Name == "" {
		course.Name = courseData.Name
	}

	if course.Description == "" {
		course.Description = courseData.Description
	}

	if course.TeacherReg == "" {
		course.TeacherReg = courseData.TeacherReg
	}

	if course.Reg == "" {
		course.Reg = courseData.Reg
	}

	update, err := db.Prepare("update courses set name=?, description=?, teacherreg=?, reg=? where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := update.Exec(course.Name, course.Description, course.TeacherReg, course.Reg, course.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsEffected

}
