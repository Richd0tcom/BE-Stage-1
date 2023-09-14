# BE-Stage-1

## Overview
This api is capable of CRUD operations on a "user" resource. It makes use of Golang while interfacing with a postgres database

## Pre-requisites

- Golang: Make sure you have Golang installed on your system.
- Docker: Make sure you have Docker installed on your system.
- golang-migrate: Install the [golang-migrate]('https://github.com/golang-migrate/migrate/tree/master/cmd/migrate') CLI tool (globally) to aid migrations.

## Setup 
A make file has been provided to aid easy setup of the project.You can edit the arguments in the makefile to customize the setup process.

Create a file named `app.env` and paste the environment variables below

```env
DB_DRIVER=postgres
DB_URI=postgresql://root:Madara123@localhost:5432/hngx?sslmode=disable

```

### steps
-  With docker running on your PC, run the `make postgres` command in your terminal. To pull a Postgres-15 docker image and start up the container.
-  Next, run the `make createdb` command in your terminal to create the database. 
-  With golang-migrate installed on your PC, Run `make migrateup` to run the migrations.
-  Finally run `go run main.go` to start the server.

