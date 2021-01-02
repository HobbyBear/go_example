package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	streamWtiter ,_ := f.NewStreamWriter("Sheet2")

	rowData := []interface{}{
		"无敌",
		"寂寞",
	}
	cell, _ := excelize.CoordinatesToCellName(1, 1)
	if err := streamWtiter.SetRow(cell, rowData); err != nil {
		log.Println("export data fail", err)
	}

	// Save xlsx file by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
