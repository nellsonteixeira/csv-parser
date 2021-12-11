package parser

import (
	"csv/models"
	"fmt"
	"os"
)

const (
	name     = iota // 0
	email           // 1
	salary          // 2
	id              // 3
	FileName = "roster1.csv"
)

type FileType1 struct {
}

func (r FileType1) Read() {
	basePath, _ := os.Getwd()
	path := basePath + FilesDir + FileName

	csvLines := OpenFile(path, FileName)

	if len(csvLines) > 0 {
		fmt.Println("Validating data...")
		Employers = []models.Employee{}

		for row := 1; row < len(csvLines); row++ {
			emp := models.Employee{
				Name:   csvLines[row][name],
				Email:  csvLines[row][email],
				Salary: csvLines[row][salary],
				Id:     csvLines[row][id],
			}

			errorMessage := Validate(emp)

			if len(errorMessage) == 0 {
				Employers = append(Employers, emp)
			} else {
				emp.Erros = errorMessage
				BadData = append(BadData, emp)
			}
		}
	} else {
		fmt.Println("The file" + FileName + "is empty")
	}
}
