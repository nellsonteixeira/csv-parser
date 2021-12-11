package main

import (
	"csv/parser"
)

func main() {
	var fileType1 parser.FileType1
	fileType1.Read()

	parser.SaveData()
	parser.SaveWrongData()
}
