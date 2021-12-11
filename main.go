package main

import (
	"csv/parser"
)

func main() {
	//Call Read operation for for FileType1
	var fileType1 parser.FileType1
	fileType1.Read()

	//Read operation for others file types struct can be developed and called

	parser.SaveData()
	parser.SaveWrongData()
}
