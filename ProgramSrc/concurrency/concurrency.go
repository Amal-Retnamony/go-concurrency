package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

type excelType map[string]interface{}


func main(){
	value, err := xlsx.OpenFile("../original.xlsx")
	if err!=nil{
		fmt.Println("eror occured", err)
	}
	for _, sheet := range value.Sheets {
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
                text := cell.String()
                fmt.Printf("%s\n", text)
            }
        }
    }
}