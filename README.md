# Implementing CRUD in GoLang REST API with Mux & GORM & Postgres DB

![CRUD in GoLang REST API with Mux & GORM & Postgres DB](https://user-images.githubusercontent.com/12379287/165107926-91a34de1-19b1-4ef6-95da-0642d0fa28ae.png)

## Features

- Implementing CRUD in Golang Rest API
  - Create
  - Get By ID
  - Get All
  - Update
  - Delete
- Testing CRUD Operations

# Development Setup

## Tools

- Go version 1.8.1
- Postgres for Database
- gorilla/mux for routing
- gorm for ORM
- viper for enviroment configuration
- Fake Data Generator for Struct 

## Run Program

Run the following steps:
`git clone github.com/CRUD-Rest-API-Golang/`
`go install`

1. Create a database named crud_rest in Postgres
2. `go run database/seeder/seeder.go` Seed the database.
3. `go run main.go` Run the app in Dev mode.
