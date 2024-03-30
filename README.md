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
- copy environment dev, test and set app_status=test for .env.test
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
    make run
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
