package main

import(
	"os"
	"encoding/csv"
	"errors"
	"strconv"
)

type Student struct {
	id int
	name string
	appointmentHour int
	appointmentMinutes int
	appointmentDay int

}
func getStudentFileName() string {
	return "students.csv"
}
func createStudentTemplate()  {
	data := [][]string{{"1","Test1","14","30","1"},{"2","Test2","15","00","5"},{"3","Test3","16","30","4"},{"4","Test4","17","45","1"}}
	writeData(data, getStudentFileName())
}

func getStudent(id int) (Student, error) {
	stuArray := getStudents(getStudentFileName())
	for _, stu := range stuArray {
		if stu.id == id {
			return stu, nil
		}
	}
	return Student{}, errors.New("Error")
}

func writeStudents(students []Student){
	data := make([][]string,len(students))
	for i,stu := range students {
		id := strconv.Itoa(stu.id)
		appointmentHour := strconv.Itoa(stu.appointmentHour)
		appointmentMinutes := strconv.Itoa(stu.appointmentMinutes)
		appointmentDay := strconv.Itoa(stu.appointmentDay)
		data[i] = []string{id,stu.name,appointmentHour,appointmentMinutes,appointmentDay}

	}
	writeData(data, getStudentFileName())
}

func getStudentId(name string) int {
	stuArray := getStudents(getStudentFileName())
	for _, stu := range stuArray {
		if stu.name == name {
			return stu.id
		}
	}
	return -1
}

func getStudentsArray() []Student {
	return getStudents(getStudentFileName())
}

func getStudents(filePath string) []Student {
	csvfile1, _ := os.Open(filePath)
	defer csvfile1.Close()
	r := csv.NewReader(csvfile1)
	r.Comma = ','
	records,_ := r.ReadAll()
	studArray := make([]Student,len(records))
	for i, stu := range records {
		if(len(stu)!=5){
			//fmt.Println("Not Correct users")
		} else {
			id, _ := strconv.Atoi(stu[0])
			appointmentHour, _ := strconv.Atoi(stu[2])
			appointmentMinutes, _ := strconv.Atoi(stu[3])
			appointmentDay, _ := strconv.Atoi(stu[4])
			//fmt.Println(id)
			studArray[i] = Student{id, stu[1], appointmentHour, appointmentMinutes, appointmentDay}

		}
	}

	return studArray
}
