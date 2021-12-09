package parser

import (
	"csv/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const (
	name   = iota // 0
	email         // 1
	salary        // 2
	id            // 3

	fileName = "roster1.csv"
)

type FileType1 struct {
}

func (r FileType1) Read() []models.Employee {
	basePath, _ := os.Getwd()
	path := basePath + FilesDir + fileName

	csvFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Opened" + fileName + "file")

		defer csvFile.Close()

		fmt.Println("Reading CSV file...")
		csvLines, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Validating data...")
		for row := 1; row < len(csvLines); row++ {
			emp := models.Employee{
				Name:   csvLines[row][name],
				Email:  csvLines[row][email],
				Salary: csvLines[row][salary],
				Id:     csvLines[row][id],
			}

			error := Validate(emp)

			if len(error.errors) == 0 {
				Employers = append(Employers, emp)
			} else {
				BadData = append(BadData, error)
			}
		}
	}
	return Employers
}
