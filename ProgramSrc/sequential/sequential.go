package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"database/sql"
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
	value, err := xlsx.OpenFile("../original.xlsx")
	if(err != nil){
		fmt.Println("error while opening excel file",err )
		return
	}
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
	if err != nil {
	  fmt.Println("error while creating db table",err)
	  return
	}
	insertStatement := `INSERT INTO demo (
		columnA,
		columnB,
		columnC
	  ) VALUES ($1, $2, $3);`
	for _, sheet := range value.Sheets {
        for _, row := range sheet.Rows {
			var args []interface{}
            for _,cell:= range row.Cells {
                args=append(args,cell.String())
			}
			_, err =db.Exec(insertStatement,args...)
			if(err != nil){
				fmt.Println("error while inserting", err)
				return
			}
        }
    }
}