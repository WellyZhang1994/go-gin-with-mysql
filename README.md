# Basic folder structure with gin framework and mysql database

---
#### Description
- A basic folder structure includes the following components: Server, Controller, Model, DB, Middleware, Forms, and Config.
- This structure is referenced from https://github.com/vsouza/go-gin-boilerplate.
- I use MySQL as the database to quickly build a backend system.
#### Usage
- restore package
    ```shell
    go mod tidy
    ```
- copy the config/.template.yaml to development.yaml and setup the correct setting
- start
   ```shell
    go run main.go
    ```