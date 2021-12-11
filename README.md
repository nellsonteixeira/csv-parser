## About this project
This project parse CSV files coming from different employers to create a list of eligible employees. <br />
Although files need to have key pieces of data, do not control the overall structure of the files. For example, column names and order can be different from file to file. <br />
### Adding new structure files
To add new file structures you just need to extend the function Read() for the new file type, because the functions OpenFile(), Validate(), WriteFile() SaveData(), SaveWrongData() was developed for reused by parse.go 
## Prerequisites

* Golang installed

## Usage
1. Run app <br />

```sh
   go run main.go
```
2. Run tests <br />

```sh
   cd test
   go test
```
3. Build <br />
  go build
```sh
