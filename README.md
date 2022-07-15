## About this project
This project parse CSV files coming from different employers to create a list of eligible employees. <br />
<br />
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
3. Build app<br />
```sh
  go build
```
