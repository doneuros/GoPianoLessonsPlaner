package main

import(
        "encoding/csv"
        "fmt"
        "os"
        "log"
        "time"
        "strconv"
)


func main(){
        createStudentTemplate()
        lessonOccur()
        createMonthFile(0, 24*4)

}

func lessonOccur() [][] string {
        month := "January"
        students := getStudentsArray()
        for _, stu := range students {
                fmt.Println("Hat die Stunde von %s stattgefunden, (Ja/Y/Yes oder Nein/N/No)", stu.name)
                var input string
                fmt.Scanln(&input)
                if(input == "Ja" || input == "Y" || input == "Yes"){
                        fmt.Print("Write to file lesson taken")
                        data := readFile(month+"_2017.csv")
                        fmt.Println("Hat die Stunde am %d. %s um %d:%d Uhr stattgefunden? (Ja/Y/Yes oder Nein/N/No)", stu.appointmentDay, month, stu.appointmentHour, stu.appointmentMinutes)
                        fmt.Scanln(&input)
                        if(input == "Nein" || input == "N" || input == "No"){
                                fmt.Println("Ist die Stunde dauerhaft verschoben? (Ja/Y/Yes oder Nein/N/No)")
                                if(input == "Ja" || input == "Y" || input == "Yes"){
                                        fmt.Println("Stunde eintrage: ")
                                        fmt.Scanln(&input)
                                        stu.appointmentHour = strconv.Atoi(input)
                                        fmt.Println("Minute eintrage: ")
                                        fmt.Scanln(&input)
                                        stu.appointmentMinutes = strconv.Atoi(input)
                                        fmt.Println("Tag eintrage: ")
                                        fmt.Scanln(&input)
                                        stu.appointmentDay = strconv.Atoi(input)
                                        //data[stu.appointmentDay][stu.appointmentHour*4+stu.appointmentMinutes/15]
                                }
                        } else {
                                data[stu.appointmentDay][stu.appointmentHour*4+stu.appointmentMinutes/15]
                        }
                }

        }
        writeStudents(students)

}

func createMonthFile(start int, end int) {
        now := time.Now()
        currentYear, currentMonth, _ := now.Date()
        currentLocation := now.Location()

        firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
        lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

        fmt.Println(firstOfMonth.Day())
        fmt.Println(lastOfMonth.Day())

        fmt.Println(currentMonth)
        fmt.Println(currentYear)
        //24*4 = Number of quarters in on Day

        days := lastOfMonth.Day()+1
        data := make([][]string,days)
        for i := 0; i< days; i++ {
                data[i] = make([]string, 24*4)
                if i==0 {
                        data[i][start] = "Tag"
                } else {
                        data[i][start] = strconv.Itoa(i)
                }


                for j :=start+1; j<end; j++ {
                        if i==start {
                                data[i][j] = strconv.Itoa((j)/4)+":"+strconv.Itoa(((j)%4)*15)
                        } else {
                                data[i][j] = strconv.Itoa(0)
                        }

                }
        }
        //For Debug
        for _, value := range data {
                fmt.Print(value)
        }


        writeData(data, currentMonth.String()+"_"+strconv.Itoa(currentYear)+".csv")



}



func readFile(filePath string) [][]string {
        csvfile1, err := os.Open(filePath)
        if err != nil {
                log.Fatal(err)
        }
        defer csvfile1.Close()

        r := csv.NewReader(csvfile1)
        //r.Comma = ','
        records,err := r.ReadAll()
        if err !=nil {
                log.Fatal(err)
        }
        return records
}

func writeData(data [][]string, fileName string) {
        if(fileExists(fileName)){
                return 
        }
        file, err := os.Create(fileName)
        checkError("Cannot create file", err)
        defer file.Close()

        writer := csv.NewWriter(file)

        for _, value := range data {
                err := writer.Write(value)
                checkError("Cannot write to file", err)
        }

        defer writer.Flush()
}

func fileExists(name string) bool {
        _, err := os.Stat(name)
        return !os.IsNotExist(err)
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
