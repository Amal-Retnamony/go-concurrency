package main

import (
	"fmt"
	// "time"
	"database/sql"
	
	"github.com/tealeg/xlsx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qburst"
	dbname   = "go_demo"
  )


func main(){
	ch := make(chan []string, 3)
	check := make(chan bool)
	go readFile(ch)
	go writeTable(ch, check)
	<- check
}

func readFile(ch chan []string) {
	value,err := xlsx.OpenFile("../original.xlsx")
	if err!=nil{
		fmt.Println("error occured while opening file", err)
		return 
	}
  	for _, sheet := range value.Sheets {
        for _, row := range sheet.Rows {
						var args []string
            for _,cell:= range row.Cells {
                args=append(args,cell.String())
			}
			  ch <- args 
			}
		close(ch)
}
}


func writeTable(ch chan []string,check chan bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	err = db.Ping()
	if err != nil{
		fmt.Println("ping error", err)
		return
	}
	defer db.Close()
	for value := range ch {
				insertStatement := `INSERT INTO demo (
					 columnA,
					 columnB,
					 columnC
					) VALUES ($1, $2, $3);`
				_, err =db.Exec(insertStatement,value[0], value[1], value[2])
				if(err != nil){
					fmt.Println("error while inserting", err)
					return
				}
	}
	check <- true
}
