package parser

import (
	"csv/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	FilesDir          = "\\input\\"
	outputDir         = "\\output\\"
	incorrectDataFile = "errors.json"
	outputFile        = "output.json"
)

type WrongData struct {
	data   models.Employee
	errors []string
}

var Employers []models.Employee
var BadData = []WrongData{}

func Validate(emp models.Employee) WrongData {
	var erro = WrongData{data: emp}

	if len(emp.Id) == 0 {
		erro.errors = append(erro.errors, "Id is Required")
	}

	if len(emp.Email) == 0 {
		erro.errors = append(erro.errors, "Email is Required")
	}

	if len(emp.Name) == 0 {
		erro.errors = append(erro.errors, "Name is Required")
	}

	if len(emp.Salary) == 0 {
		erro.errors = append(erro.errors, "Salary is Required")
	}

	if len(erro.errors) > 0 {
		BadData = append(BadData, erro)
	}

	return erro
}

func SaveData(emp []models.Employee) {

	if len(emp) > 0 {
		fmt.Println("Saving data...")
		basePath, _ := os.Getwd()
		path := basePath + outputDir

		if _, Err := os.Stat(path); os.IsNotExist(Err) {
			os.Mkdir(path, os.ModePerm)
		}

		data, _ := json.Marshal(emp)
		filePath := path + outputFile

		writeFail := ioutil.WriteFile(filePath, data, 0644)

		if writeFail != nil {
			log.Fatal(writeFail)
		}
		fmt.Println("Successfully saved Corret data file")
	}
}

func SaveWrongData(erros []WrongData) {
	if len(erros) > 0 {
		fmt.Println("Saving incorrect data file...")
		basePath, _ := os.Getwd()
		path := basePath + outputDir

		if _, Err := os.Stat(path); os.IsNotExist(Err) {
			os.Mkdir(path, os.ModePerm)
		}

		data, _ := json.Marshal(erros)
		filePath := path + incorrectDataFile

		writeFail := ioutil.WriteFile(filePath, data, 0644)

		if writeFail != nil {
			log.Fatal(writeFail)
		}
		fmt.Println("Successfully saved incorrect data file")
	}
}
