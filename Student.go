package main

import(
	"os"
	"encoding/csv"
	"fmt"
	"errors"
)

type Student struct {
	id int
	name string

}
func getStudentFileName() string {
	return "students.csv"
}
func createStudentTemplate()  {
	data := [4][2]string{{"1","Test1"},{"2","Test2"},{"3","Test3"},{"4","Test4"}}
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

func getStudentId(name string) int {
	stuArray := getStudents(getStudentFileName())
	for _, stu := range stuArray {
		if stu.name == name {
			return stu.id
		}
	}
	return -1
}

func getStudents(filePath string) []Student {
	csvfile1, _ := os.Open(filePath)
	defer csvfile1.Close()
	r := csv.NewReader(csvfile1)
	r.Comma = ';'
	records,_ := r.ReadAll()
	studArray := make([]Student,len(records))
	for i, stu := range records {
		if(len(stu)!=2){
			fmt.Println("Not Correct users")
		} else {
			id, _ := strconv.Atoi(stu[0])
			studArray[i] = Student{id, stu[1]}
		}
	}

	return studArray
}