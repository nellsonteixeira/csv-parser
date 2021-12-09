package parser

import "csv/models"

type File interface {
	Read() []*models.Employee
}
