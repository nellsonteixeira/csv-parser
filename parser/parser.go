package parser

import (
	"csv/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

const (
	FilesDir   = "\\input\\"
	outputDir  = "\\output\\"
	errosFile  = "errors.json"
	outputFile = "output.json"
)

var Employers []models.Employee
var BadData []models.Employee

func Validate(emp models.Employee) string {
	var erro string

	ind := sort.Search(len(Employers), func(ind int) bool {
		return Employers[ind].Email <= emp.Email || Employers[ind].Id <= emp.Id
	})
	if len(emp.Email) > 0 && len(emp.Email) > 0 &&
		ind < len(Employers) &&
		(Employers[ind].Email == emp.Email || Employers[ind].Id == emp.Id) {
		erro += "Id or Email already entered."
	}

	if len(emp.Id) == 0 {
		erro += "Id is Required."
	}

	if len(emp.Email) == 0 {
		erro += "Email is Required."
	}

	if len(emp.Name) == 0 {
		erro += "Name is Required."
	}

	if len(emp.Salary) == 0 {
		erro += "Salary is Required."
	}

	return erro
}

func OpenFile(path string, fileName string) [][]string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened" + fileName + "file")

	defer file.Close()
	fmt.Println("Reading CSV file...")
	csvLines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	return csvLines
}

func SaveData() {
	if len(Employers) > 0 {
		fmt.Println("Saving data...")

		basePath, _ := os.Getwd()
		path := basePath + outputDir

		WriteFile(path, outputFile, Employers)
		fmt.Println("Successfully saved Correct data file in " + outputDir)
	}
}
func SaveWrongData() {
	if len(BadData) > 0 {
		fmt.Println("Saving incorrect data file...")

		basePath, _ := os.Getwd()
		path := basePath + outputDir

		WriteFile(path, errosFile, BadData)
		fmt.Println("Successfully saved Incorrect data file in " + outputDir)
	}
}

func WriteFile(path string, fileName string, employers []models.Employee) {
	if _, Err := os.Stat(path); os.IsNotExist(Err) {
		os.Mkdir(path, os.ModePerm)
	}

	data, _ := json.Marshal(employers)
	filePath := path + fileName

	writeFail := ioutil.WriteFile(filePath, data, 0644)

	if writeFail != nil {
		log.Fatal(writeFail)
	}
}
