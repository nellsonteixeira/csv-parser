package main

import (
	"csv/parser"
)

func main() {
	var fileType1 parser.FileType1

	data := fileType1.Read()
	parser.SaveData(data)

	parser.SaveWrongData(parser.BadData)
}
