package main

import(
	"os"
	"encoding/csv"
	"fmt"
)

type Student struct {
	id int
	name string

}
func getStudentFileName() string {
	return "students.csv"
}
func createStudentTemplate()  {
	writeData([4][2]string{{"1","Test1"},{"2","Test2"},{"3","Test3"},{"4","Test4"}}, getStudentFileName())
}

func getStudent(id long) Student {
	stuArray := getStudents(getStudentFileName())
	for _, stu := range stuArray {
		if stu.id == id {
			return stu
		}
	}
	return _
}

func getStudentId(name string) int {
	stuArray := getStudents(getStudentFileName())
	for _, stu := range stuArray {
		if stu.name == name {
			return stu
		}
	}
	return _
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
			studArray[i] = Student{stu[0], stu[1]}
		}
	}

	return studArray
}