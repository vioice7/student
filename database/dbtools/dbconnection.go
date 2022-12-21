package dbtools

import (
	"database/sql"
	"fmt"
	"log"
	"rest/database/model"

	_ "github.com/go-sql-driver/mysql"
)

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

func Save(student model.Student) int64 {

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

func Update(student model.Student) int64 {

	db := connect()

	defer db.Close()

	// get the old data from the database

	rows := db.QueryRow("select * from students where id = ?", student.ID)

	// create a new model to compare

	studentData := model.Student{}

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

	// check for unique reg

	rowsDupl := db.QueryRow("select reg from students where reg = ?", student.Reg)

	errDupl := rowsDupl.Scan(&student.Reg)

	// if there is a duplicate entry don't add a record and return

	if errDupl == nil {
		fmt.Println("Duplicate entry in database! No records added!")
		return 0
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

func DeleteId(id int) int64 {

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

func DeleteAll() int64 {

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
