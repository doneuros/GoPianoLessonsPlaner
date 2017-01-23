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
        records := readFile("test.csv")
        for _, h := range records {
           for _, cell := range h {
                fmt.Print(cell, " ")
            }
            fmt.Println()
        }
        //writeData(records, "result.csv")
        createMonthFile(0, 24*4)

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


                for j :=start+1; j<end+1; j++ {
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
        writeData(data, currentMonth+":"+currentYear+".csv")




}

func readFile(filePath string) [][]string {
        csvfile1, err := os.Open(filePath)
        if err != nil {
                log.Fatal(err)
        }
        defer csvfile1.Close()

        r := csv.NewReader(csvfile1)
        r.Comma = ';'
        records,err := r.ReadAll()
        if err !=nil {
                log.Fatal(err)
        }
        return records
}

func writeData(data [][]string, fileName string) {
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

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
