## About this project
This project parse CSV files coming from different employers to create a list of eligible employees. <br />
Although files need to have key pieces of data, do not control the overall structure of the files. For example, column names and order can be different from file to file. <br />
### Built With
To add new file structures you just need to extend the function Read for that new file types because the functions OpenFile(), Validate(), WriteFile() SaveData(), SaveWrong() was developed for reused by parse.go 
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
