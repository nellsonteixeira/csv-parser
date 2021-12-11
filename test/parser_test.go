package parser

import (
	"csv/models"
	"csv/parser"
	"os"
	"testing"
)

func TestValidate(t *testing.T) {
	emp := models.Employee{
		Id:     "1",
		Email:  "",
		Salary: "100",
		Name:   "Joe",
	}
	err := parser.Validate(emp)

	expect := "Email is Required."
	if err != expect {
		t.Errorf("Validate was incorrect, got: %s, want: %s.", err, expect)
	}
}

func TestOpenFile(t *testing.T) {
	fileName := "roster1.csv"
	basePath, _ := os.Getwd()
	path := basePath + "\\input\\" + fileName

	data := parser.OpenFile(path, fileName)

	if data != nil && len(data) == 0 {
		t.Errorf("OpenFile was incorrect, got: %d, want: %s.", 0, "data lenght > 0")
	}
}
