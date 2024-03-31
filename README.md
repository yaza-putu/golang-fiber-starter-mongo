# Golang Starter API With Mongo
![golang](https://repository-images.githubusercontent.com/346033147/dfa04d00-974e-11eb-8e48-94f80160841c)

Project layout design influenced by [standard go project layout](https://docs.gofiber.io/img/logo-dark.svg)
### Note : Be sure to rename the name of module before beginning the project.
## How to start


- install depedency
  ```bash
  make tidy
  # or
  go mod tidy
  ```
- copy environment
  ```bash
  make config
  #or
  cp .env.example .env
  ```


- generate key
  ```bash
    make key
  ```

- run dev mode
  ```bash
    make serve
  ```
- build
  ```bash
  make build
  ```

- run test
  ```bash
   make gotest
  ```

- make migration
  ```bash
   make migration table="name_of_table"
  ```
  
- run migration
  ```bash
   make migrate-up
  ```


## Validation
- unique
```go
type v struct {
	Name string `validate:"unique=table_name:column_name"`
}
// ecample
type v struct {
Name string `validate:"unique=users:name"`
}
```
- unique with ignore
```go
type v struct {
Name string `validate:"unique=table_name:column_name:ignore_with_field_name"`
ID   string `validate:"required"`
}
// example
type v struct {
Name string `validate:"unique=users:name:ID"`
ID   string `validate:"required" json:"id"`
}
```

- file validation
```go
type FileHandler struct {
  File multipart.File `validate:"required,filetype=image/png image/jpeg image/jpg"`
}

fs := FileHandler{}

f, err := ctx.FormFile("file")
if err == nil {
// send file into FileHandler struct to validate
  fs.File, err = f.Open()
  if err != nil {
    return err
  }
}
// validate with custom validation from go-playground/validator 
val, err := request.Validation(&fs)
