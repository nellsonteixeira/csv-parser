package parser

import (
	"csv/parser"
	"testing"
)

func TestRead(t *testing.T) {
	var fileType1 parser.FileType1
	fileType1.Read()

	if len(parser.Employers) == 0 {
		t.Errorf("Read( was incorrect, got: %d, want: %s.", 0, "Employers lenght > 0")
	}

}
